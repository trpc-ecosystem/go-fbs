//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 THL A29 Limited, a Tencent company.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

package fbs

import (
	"fmt"
	"sort"
	"strings"

	"trpc.group/trpc-go/fbs/internal/ast"
)

// sentinelMissingSymbol is used when a symbol cannot be resolved as valid descriptor, but is
// part of the namespaces.
var sentinelMissingSymbol = &TableDesc{}

// linker links multiple files. It checks duplicate type definition and resolves type references.
type linker struct {
	// files map filenames to *parseResult.
	// Note: If parseResults.recursive (see newLinker) is set, then `files` will contain
	// all the files being included by the user provided files.
	files map[string]*parseResult
	// filenames contain the keys of files.
	filenames []string
	handler   *errorHandler
	// descPool maps schema descriptor to a map(pool) that maps fully qualified name to the corresponding
	// descriptor.
	// Examples:
	// schemaDesc =>
	//   "rpc.app.server.MyTable" => tabledesc
	//   "rpc.app.server.MyTable.field1" => field1desc
	//   "rpc.app.server.MyStruct" => structdesc
	//   "rpc.app.server.MyStruct.field2" => field2desc
	//   "rpc.app.server.MyEnum" => enumdesc
	//   "rpc.app.server.MyEnumVal" => enumvaldesc
	//   "rpc.app.server.MyUnion" => uniondesc
	//   "rpc.app.server.MyUnionVal" => unionvaldesc
	//   "rpc.app.server.MyRPCService" => rpcdesc
	//   "rpc.app.server.MyRPCService.MyMethod" => methoddesc
	descPool map[*SchemaDesc]map[string]Desc
	// packageNamespaces maps schema descriptor to a map that contains a set of namespace strings.
	// This is used to check whether a string can be resolved as a part of namespace.
	packageNamespaces map[*SchemaDesc]map[string]struct{}
	// usedIncludes maps schema descriptor to a map that contains a set of included file names.
	// This is used to check unused includes.
	usedIncludes map[*SchemaDesc]map[string]struct{}
}

// newLinker creates a linker.
func newLinker(results *parseResults, handler *errorHandler) *linker {
	return &linker{
		files:     results.resultsByFilename,
		filenames: results.filenames,
		handler:   handler,
	}
}

// linkFiles links the parsed results for each independent file.
// Two steps are used to produce the final results:
// 1. Put all symbols(type definitions) into a pool and check duplicate definitions.
// 2. Resolve type references using the pool created in the prior step.
func (l *linker) linkFiles() (map[string]*SchemaDesc, error) {
	// Step1: Put all symbols into a pool. Ensure no duplicates.
	if err := l.createDescPool(); err != nil {
		return nil, err
	}
	// Step2: Try to resolve all type references. They will be re-written
	// to be fully-qualified references (with leading dot '.' ).
	if err := l.resolveReferences(); err != nil {
		return nil, err
	}
	// Result of Step2: field type name of tables/structs and input/output type name of rpc methods
	// will all be solved as fully-qualified name. Example: .rpc.app.server.InputTypeName
	linked := map[string]*SchemaDesc{}
	for name, result := range l.files {
		linked[name] = result.fd
	}
	return linked, nil
}

// createDescPool fills descPool defined in the linker. Returns error if any
// duplicates are found across the files.
func (l *linker) createDescPool() error {
	l.descPool = map[*SchemaDesc]map[string]Desc{}
	l.packageNamespaces = map[*SchemaDesc]map[string]struct{}{}
	if err := l.addDescToPool(); err != nil {
		return err
	}
	if err := l.symbolDuplicated(); err != nil {
		return err
	}
	return nil
}

// addDescToPool will fill each pool with fully qualified name to
// descriptor map. Duplicates will be checked at each file level.
func (l *linker) addDescToPool() error {
	for _, filename := range l.filenames {
		r := l.files[filename]
		pool := map[string]Desc{}
		l.descPool[r.fd] = pool
		// fd.Namespaces example: ["namespace1", "namespace2", "rpc.app.server"]
		l.packageNamespaces[r.fd] = getAllNamespaces(r.fd.Namespaces)
		if err := l.addTableStructToPool(r); err != nil {
			return err
		}
		if err := l.addEnumUnionToPool(r); err != nil {
			return err
		}
		if err := l.addRPCServiceToPool(r); err != nil {
			return err
		}
	}
	return nil
}

// addTableStructToPool iterates through all tables and structs to add their type definition
// into pool. The same function addTableStruct is used to reduce redundant code.
func (l *linker) addTableStructToPool(r *parseResult) error {
	for _, d := range r.fd.Tables {
		if err := l.addTableStruct(r, d); err != nil {
			return err
		}
	}
	for _, d := range r.fd.Structs {
		if err := l.addTableStruct(r, d); err != nil {
			return err
		}
	}
	return nil
}

// addTableStruct can add either table type or struct type into the pool.
func (l *linker) addTableStruct(r *parseResult, d TableStructDesc) error {
	prefix := getPrefix(d) // use its own namespace.
	// fqn: fully qualified name
	fqn := prefix + d.GetName() // example: "rpc.app.server.MyTable"
	if err := l.addToPool(r, fqn, d); err != nil {
		return err
	}
	prefix = fqn + "." // example: "rpc.app.server.MyTable."
	for _, dd := range d.GetFields() {
		if err := l.addFieldToPool(r, prefix, dd); err != nil {
			return err
		}
	}
	return nil
}

// addEnumUnionToPool iterates through all enumerations and unions to the pool.
func (l *linker) addEnumUnionToPool(r *parseResult) error {
	for _, d := range r.fd.Enums {
		if err := l.addEnumToPool(r, d); err != nil {
			return err
		}
	}
	for _, d := range r.fd.Unions {
		if err := l.addUnionToPool(r, d); err != nil {
			return err
		}
	}
	return nil
}

// addRPCServiceToPool add all rpc service definitions into the pool.
func (l *linker) addRPCServiceToPool(r *parseResult) error {
	for _, d := range r.fd.RPCs {
		if err := l.addRPCToPool(r, d); err != nil {
			return err
		}
	}
	return nil
}

// addEnumToPool add enumeration definitions into the pool.
func (l *linker) addEnumToPool(r *parseResult, d *EnumDesc) error {
	prefix := getPrefix(d) // use its own namespace.
	fqn := prefix + d.Name // example: "rpc.app.server.MyEnum"
	if err := l.addToPool(r, fqn, d); err != nil {
		return err
	}
	for _, dd := range d.Values {
		// enum value type name is in the same scope as the enum name.
		vfqn := prefix + dd.Name // example: "rpc.app.server.MyEnumValueName"
		if err := l.addToPool(r, vfqn, dd); err != nil {
			return err
		}
	}
	return nil
}

// addUnionToPool add union definitions into the pool.
func (l *linker) addUnionToPool(r *parseResult, d *UnionDesc) error {
	prefix := getPrefix(d) // use its own namespace.
	fqn := prefix + d.Name // example: "rpc.app.server.MyUnion"
	if err := l.addToPool(r, fqn, d); err != nil {
		return err
	}
	return nil
}

// addRPCToPool add rpc service name into the pool. It further iterates through all its
// methods to add method name into the pool.
func (l *linker) addRPCToPool(r *parseResult, d *RPCDesc) error {
	prefix := getPrefix(d) // use its own namespace.
	fqn := prefix + d.Name // example: "rpc.app.server.MyRPCService"
	if err := l.addToPool(r, fqn, d); err != nil {
		return err
	}
	for _, dd := range d.Methods {
		mfqn := fqn + "." + dd.Name // example: "rpc.app.server.MyRPCService.MyMethod"
		if err := l.addToPool(r, mfqn, dd); err != nil {
			return err
		}
	}
	return nil
}

// addFieldToPool add field name into the pool.
func (l *linker) addFieldToPool(r *parseResult, prefix string, d *FieldDesc) error {
	fqn := prefix + d.Name
	return l.addToPool(r, fqn, d)
}

// addToPool manipulates linker.descPool to map fully qualified name to descriptor.
func (l *linker) addToPool(r *parseResult, fqn string, dsc Desc) error {
	if d, ok := l.descPool[r.fd][fqn]; ok { // ok means duplicate!
		node := r.descToNode[dsc]
		if err := l.handler.handleErrorWithPos(node.Start(),
			"duplicate symbol %s: already defined as %s", fqn, descType(d)); err != nil {
			return err
		}
	}
	l.descPool[r.fd][fqn] = dsc
	return nil
}

// entry combines filename and its descriptor. This helps detect duplicated symbols.
type entry struct {
	file string
	dsc  Desc
}

// symbolDuplicated will check duplicate symbols according to all files rather
// than per-file level.
func (l *linker) symbolDuplicated() error {
	// Put everything into a single pool, ensure no symbol is declared more than once.
	pool := map[string]entry{}
	for _, filename := range l.filenames {
		fd := l.files[filename].fd
		p := l.descPool[fd]
		keys := make([]string, 0, len(p))
		keys = appendDescPoolKeys(keys, p)
		sort.Strings(keys) // Sort the keys to generate error deterministically.
		for _, k := range keys {
			v := p[k]
			if e, ok := pool[k]; ok {
				return l.errDuplicateSymbol(k, &e, &entry{fd.Name, v})
			}
			pool[k] = entry{file: fd.Name, dsc: v}
		}
	}
	return nil
}

// appendDescPoolKeys appends the keys of the descriptor pool.
func appendDescPoolKeys(keys []string, p map[string]Desc) []string {
	for k := range p {
		keys = append(keys, k)
	}
	return keys
}

// errDuplicateSymbol forms a deterministic error to indicate that a symbol has already been
// defined before. If it is defined in two files, error will report that it is defined first in
// the file that has an alphabetically smaller filename.
func (l *linker) errDuplicateSymbol(s string, e1, e2 *entry) error {
	if e2.file < e1.file {
		e1, e2 = e2, e1
	}
	node := l.files[e2.file].descToNode[e2.dsc]
	return l.handler.handleErrorWithPos(node.Start(),
		"duplicate symbol %s: already defined as %s in %q", s, descType(e1.dsc), e1.file)
}

// resolveReferences resolves type references using type definitions stored in the pool.
func (l *linker) resolveReferences() error {
	l.usedIncludes = map[*SchemaDesc]map[string]struct{}{}
	for _, filename := range l.filenames {
		r := l.files[filename]
		fd := r.fd
		scopes := []scope{schemaScope(fd, l)}
		if err := l.resolveTypeReferences(r, scopes); err != nil {
			return err
		}
		for _, d := range fd.RPCs {
			if err := l.resolveRPCs(r, d, scopes); err != nil {
				return err
			}
		}
	}
	return nil
}

// resolveTypeReferences resolves types used in the fields of tables/structs. Example:
//
//	table Monster { pos:Vec3 (id: 0); }
//	                    ^^^^ This is going to be resolved.
func (l *linker) resolveTypeReferences(r *parseResult, scopes []scope) error {
	for _, d := range r.fd.Tables {
		if err := l.resolveTableStruct(r, d, scopes); err != nil {
			return err
		}
	}
	for _, d := range r.fd.Structs {
		if err := l.resolveTableStruct(r, d, scopes); err != nil {
			return err
		}
	}
	return nil
}

// resolveTableStruct resolves either field of table or struct.
func (l *linker) resolveTableStruct(r *parseResult, d TableStructDesc, scopes []scope) error {
	prefix := getPrefix(d)
	fqn := prefix + d.GetName() // example: "rpc.app.server.MyTable"
	prefix = fqn + "."          // example: "rpc.app.server.MyTable."
	for _, dd := range d.GetFields() {
		if err := l.resolveFields(r, prefix, dd, scopes); err != nil {
			return err
		}
	}
	return nil
}

// ReqRspType provides an interface for Request(input) and Response(output) types.
type ReqRspType interface {
	MethodName() string
	TypeName() string
	SetTypeName(string)
	SetTypeDesc(*TableDesc)
	StartPosition() *ast.Position
}

var _ ReqRspType = (*ReqType)(nil)
var _ ReqRspType = (*RspType)(nil)

// ReqType implements ReqRspType, representing request type.
type ReqType struct {
	r  *parseResult
	dd *MethodDesc
}

// MethodName implements interface ReqRspType.
func (r *ReqType) MethodName() string {
	return r.dd.Name
}

// TypeName implements interface ReqRspType.
func (r *ReqType) TypeName() string {
	return r.dd.InputType
}

// SetTypeName implements interface ReqRspType.
func (r *ReqType) SetTypeName(s string) {
	r.dd.InputType = s
}

// SetTypeDesc implements interface ReqRspType.
func (r *ReqType) SetTypeDesc(d *TableDesc) {
	r.dd.InputTypeDesc = d
}

// StartPosition implements interface ReqRspType.
func (r *ReqType) StartPosition() *ast.Position {
	return r.r.getMethodNode(r.dd).ReqName.Start()
}

// RspType implements ReqRspType, representing response type.
type RspType struct {
	r  *parseResult
	dd *MethodDesc
}

// MethodName implements interface ReqRspType.
func (r *RspType) MethodName() string {
	return r.dd.Name
}

// TypeName implements interface ReqRspType.
func (r *RspType) TypeName() string {
	return r.dd.OutputType
}

// SetTypeName implements interface ReqRspType.
func (r *RspType) SetTypeName(s string) {
	r.dd.OutputType = s
}

// SetTypeDesc implements interface ReqRspType.
func (r *RspType) SetTypeDesc(d *TableDesc) {
	r.dd.OutputTypeDesc = d
}

// StartPosition implements interface ReqRspType.
func (r *RspType) StartPosition() *ast.Position {
	return r.r.getMethodNode(r.dd).RspName.Start()
}

// resolveRPCs resolves type references used in methods' input/output. Examples:
//
//	rpc_service MonsterStorage { Store(Monster):Stat (streaming: "none"); }
//	                                   ^^^^^^^  ^^^^
//	                                 InputType    OutputType
//	                                      ^^^^^^^^^^^^^^^  This is going to be resolved.
func (l *linker) resolveRPCs(r *parseResult, d *RPCDesc, scopes []scope) error {
	prefix := getPrefix(d)
	rpcServiceName := prefix + d.Name
	for _, dd := range d.Methods {
		// resolve request type
		if err := l.resolveReqRsp(&ReqType{r, dd}, r, scopes, rpcServiceName); err != nil {
			return err
		}
		// resolve response type
		if err := l.resolveReqRsp(&RspType{r, dd}, r, scopes, rpcServiceName); err != nil {
			return err
		}
	}
	return nil
}

// resolveReqRsp resolves either input(request) type or output(response) type. This function use interface to
// avoid duplicate code.
func (l *linker) resolveReqRsp(rt ReqRspType, r *parseResult, scopes []scope, rpcServiceName string) error {
	scope := fmt.Sprintf("method %s.%s", rpcServiceName, rt.MethodName())
	fqn, dsc := l.resolve(r.fd, rt.TypeName(), scopes)
	if dsc == nil {
		return l.handler.handleErrorWithPos(rt.StartPosition(), "%s: unknown response type %s",
			scope, rt.TypeName())
	}
	if dsc == sentinelMissingSymbol {
		return l.handler.handleErrorWithPos(rt.StartPosition(),
			"%s: unknown response type %s; resolved to %s which is not defined",
			scope, rt.TypeName(), fqn)
	}
	d, ok := dsc.(*TableDesc)
	if !ok {
		otherType := descType(dsc)
		return l.handler.handleErrorWithPos(rt.StartPosition(),
			"%s: invalid response type: %s is a %s, not a table", scope, fqn, otherType)
	}
	rt.SetTypeName("." + fqn)
	rt.SetTypeDesc(d)
	return nil
}

// resolveFields resolves TypeName used in tables/structs to be fully qualified name. Examples:
//
//	table Monster { pos : namespace2.Vec3; }
//	                      ^^^^^^^^^^^^^^^
//	                       TypeName
//	                       ^^^^^^^^ needed to be resolved
//	               goal: resolve it to be .namespace2.Vec3
//
//	// MyStruct is defined inside namespace rpc.app.server
//	// server2.MyFieldTypeName is defined in namespace rpc.app.server2.
//	struct MyStruct { myfieldname : server2.MyFieldTypeName; }
//	                                ^^^^^^^^^^^^^^^^^^^^^^^
//	               goal: resolve it to be .rpc.app.server2.MyFieldTypeName
//
//	// MyFieldTypeName is defined inside namespace rpc.app.server2
//	struct MyFieldTypeName { .. }
//	// Prior to the resolve step, symbol `rpc.app.server2.MyFieldTypeName` has been put into
//	// `linker.descPool`, therefore it can be successfully resolved.
func (l *linker) resolveFields(r *parseResult, prefix string, d *FieldDesc, scopes []scope) error {
	if _, ok := keywords[d.TypeName]; ok {
		return nil
	}
	thisName := prefix + d.Name // example: "rpc.app.server.MyTable.MyFieldName"
	scope := fmt.Sprintf("field %s", thisName)
	node := r.getFieldNode(d)
	// d.TypeName example: "namespace2.MyFieldTypeName"
	fqn, dsc := l.resolve(r.fd, d.TypeName, scopes)
	if dsc == nil {
		return l.handler.handleErrorWithPos(node.Start(), "%s: unknown type %s", scope, d.TypeName)
	}
	if dsc == sentinelMissingSymbol {
		return l.handler.handleErrorWithPos(node.Start(),
			"%s: unknown type %s; resolved to %s which is not defined", scope, d.TypeName, fqn)
	}
	switch dsc := dsc.(type) {
	case *TableDesc, *StructDesc, *EnumDesc, *UnionDesc:
		d.TypeName = "." + fqn // Transform d.TypeName to be fully qualified.
	default:
		otherType := descType(dsc)
		return l.handler.handleErrorWithPos(node.Start(), "%s: invalid type: %s is a %s",
			scope, fqn, otherType)
	}
	return nil
}

// resolve resolves `name` to be a predefined descriptor (either defined in this file scope or in
// the included files' scope). Typically, name is the type name used in fields in table/struct or methods
// in rpc declarations. Examples:
//
//	table Monster { pos : namespace2.Vec3; }
//	                ^^^   ^^^^^^^^^^^^^^^
//	               Name    TypeName
//	                       ^^^^^^^^ needed to be resolved
//
//	rpc_service MonsterStorage { Store(Monster):Stat (streaming: "none"); }
//	                                   ^^^^^^^  ^^^^
//	                                  InputType  OutputType
//	                                      ^^^^^^^^^^^^^   needed to be resolved
func (l *linker) resolve(fd *SchemaDesc, name string, scopes []scope) (string, Desc) {
	if strings.HasPrefix(name, ".") {
		// already fully-qualified
		d := l.findSymbol(fd, name[1:])
		if d != nil {
			return name[1:], d
		}
		return "", nil
	}
	// unqualified, look in the enclosing (last) scope first and move towards
	// outermost (first) scope, trying to resolve the symbol
	pos := strings.IndexByte(name, '.')
	firstName := name
	if pos > 0 {
		firstName = name[:pos] // example: "namespace2"
	}
	return l.resolveUnqualified(firstName, name, scopes)
}

// resolveUnqualified uses scopes to resolve the unqualified name.
func (l *linker) resolveUnqualified(firstName, name string, scopes []scope) (string, Desc) {
	var bestGuess Desc
	var bestGuessFqn string
	for i := len(scopes) - 1; i >= 0; i-- {
		// example: firstName: "namespace2"
		// example: name: "namespace2.MyFieldTypeName"
		fqn, d := scopes[i](firstName, name)
		if d != nil {
			if isType(d) {
				return fqn, d
			}
			if bestGuess == nil {
				bestGuess = d
				bestGuessFqn = fqn
			}
		}
	}
	return bestGuessFqn, bestGuess
}

// findSymbol will find `name` in fd and its included file's descPool.
func (l *linker) findSymbol(fd *SchemaDesc, name string) Desc {
	return l.findSymbolRecursive(fd, fd, name, map[*SchemaDesc]struct{}{})
}

// findSymbolRecursive will find `name` in fd's descPool recursively.
func (l *linker) findSymbolRecursive(entryPoint, fd *SchemaDesc, name string,
	checked map[*SchemaDesc]struct{}) Desc {
	if _, ok := checked[fd]; ok {
		return nil
	}
	checked[fd] = struct{}{}
	d := l.findSymbolInFile(name, fd)
	if d != nil {
		return d
	}
	d = l.findSymbolFromIncludes(entryPoint, fd, name, checked)
	if d != nil {
		return d
	}
	return nil
}

// findSymbolFromIncludes iterates all included files to find symbol.
func (l *linker) findSymbolFromIncludes(entryPoint, fd *SchemaDesc, name string,
	checked map[*SchemaDesc]struct{}) Desc {
	for _, incl := range fd.Includes {
		res := l.files[incl]
		if res == nil {
			continue
		}
		if d := l.findSymbolRecursive(entryPoint, res.fd, name, checked); d != nil {
			l.markUsed(entryPoint, res.fd)
			return d
		}
	}
	return nil
}

// findSymbolInFile will find `name` in fd's descPool.
func (l *linker) findSymbolInFile(name string, fd *SchemaDesc) Desc {
	d, ok := l.descPool[fd][name]
	if ok {
		return d
	}
	_, ok = l.packageNamespaces[fd][name]
	if ok {
		// this name is a valid namespace but not a descriptor
		return sentinelMissingSymbol
	}
	return nil
}

// markUsed marks the include file as used.
func (l *linker) markUsed(entryPoint, used *SchemaDesc) {
	includesForFile := l.usedIncludes[entryPoint]
	if includesForFile == nil {
		includesForFile = map[string]struct{}{}
		l.usedIncludes[entryPoint] = includesForFile
	}
	includesForFile[used.Name] = struct{}{}
}
