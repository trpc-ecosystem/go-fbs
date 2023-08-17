package fbs

import (
	"math"

	"trpc.group/trpc-go/fbs/internal/ast"
)

// String literals for streaming option used in metadata fields of rpc methods.
const (
	Streaming       = "streaming"
	ClientStreaming = "client"
	ServerStreaming = "server"
	BidiStreaming   = "bidi"
)

// parseResults stores all parsed results index by their original filename.
type parseResults struct {
	// resultsByFilename maps filename to *parseResult, Example:
	//
	//  "monster_test.fbs" => parseres1
	//  "dir1/file1.fbs" => parseres2
	resultsByFilename map[string]*parseResult
	// filenames stores keys in resultsByFilename, in parsing order. Example:
	//
	//  ["monster_test.fbs", "dir1/file1.fbs"]
	filenames []string
	// recursive decides whether to parse included files recursively.
	recursive bool
	// createDescriptorFbs decides whether to create fbs descriptor. Default is true.
	createDescriptorFbs bool
}

func (p *parseResults) has(filename string) bool {
	_, ok := p.resultsByFilename[filename]
	return ok
}

func (p *parseResults) add(filename string, result *parseResult) {
	p.resultsByFilename[filename] = result
	p.filenames = append(p.filenames, filename)
}

// parseResult stores a single parse result of a flatbuffers file.
type parseResult struct {
	// root stores SchemaNode in the abstract syntax tree.
	root *ast.SchemaNode
	// fd stores schema descriptor.
	fd *SchemaDesc
	// handler stores error handler.
	handler *errorHandler
	// descToNode map descriptor to node in AST.
	// this is typically used to extract position
	// information when validating the descriptor.
	descToNode map[Desc]ast.Node
}

// newParseResult creates a new parse result out of schema node.
func newParseResult(filename string, schema *ast.SchemaNode, handler *errorHandler, createFbs bool) *parseResult {
	res := &parseResult{
		root:       schema,
		handler:    handler,
		descToNode: map[Desc]ast.Node{},
	}
	if createFbs {
		res.createSchemaDescriptor(filename, schema)
	}
	return res
}

// createSchemaDescriptor traverses the schema node to construct schema descriptor.
func (p *parseResult) createSchemaDescriptor(filename string, schema *ast.SchemaNode) {
	fd := &SchemaDesc{
		Schema:     schema,
		Name:       filename,
		Namespaces: []string{""},
	}
	p.fd = fd
	p.putSchemaNode(fd, schema)
	for _, incl := range schema.Includes {
		fd.Includes = append(fd.Includes, incl.Name.Val)
	}
	for _, decl := range schema.Decls {
		switch decl := decl.(type) {
		case *ast.TableDeclNode:
			fd.Tables = append(fd.Tables, p.asTableDesc(decl))
		case *ast.StructDeclNode:
			fd.Structs = append(fd.Structs, p.asStructDesc(decl))
		case *ast.EnumDeclNode:
			fd.Enums = append(fd.Enums, p.asEnumDesc(decl))
		case *ast.UnionDeclNode:
			fd.Unions = append(fd.Unions, p.asUnionDesc(decl))
		case *ast.RPCDeclNode:
			fd.RPCs = append(fd.RPCs, p.asRPCDesc(decl))
		case *ast.NamespaceDeclNode: // record namespaces as we run through the decl list.
			fd.Namespaces = append(fd.Namespaces, string(decl.Name.Identifier()))
		case *ast.RootDeclNode:
			fd.Root = string(decl.Name.Identifier())
		case *ast.FileExtDeclNode:
			fd.FileExt = decl.Name.Val
		case *ast.FileIdentDeclNode:
			fd.FileIdent = decl.Name.Val
		case *ast.AttrDeclNode:
			fd.Attrs = append(fd.Attrs, decl.Name.Val)
		default:
			// no action
		}
	}
}

func (p *parseResult) asTableDesc(n *ast.TableDeclNode) *TableDesc {
	d := &TableDesc{
		Schema:    p.fd,
		Namespace: p.fd.Namespaces[len(p.fd.Namespaces)-1],
		Name:      n.Name.Val,
	}
	p.putTableNode(d, n)
	p.addTableFields(d, n.Fields)
	return d
}

func (p *parseResult) asStructDesc(n *ast.StructDeclNode) *StructDesc {
	d := &StructDesc{
		Namespace: p.fd.Namespaces[len(p.fd.Namespaces)-1],
		Name:      n.Name.Val,
	}
	p.putStructNode(d, n)
	p.addStructFields(d, n.Fields)
	return d
}

func (p *parseResult) asEnumDesc(n *ast.EnumDeclNode) *EnumDesc {
	d := &EnumDesc{
		Namespace: p.fd.Namespaces[len(p.fd.Namespaces)-1],
		Name:      n.Name.Val,
	}
	p.putEnumNode(d, n)
	enumNum := int32(0)
	// Append backwards, since the constructed AST is in reverse order.
	for i := len(n.Decls) - 1; i >= 0; i-- {
		decl := n.Decls[i]
		d.Values = append(d.Values, p.asEnumVal(decl, &enumNum))
	}
	return d
}

func (p *parseResult) asEnumVal(n *ast.EnumValueNode, enumNum *int32) *EnumValDesc {
	if n.IntVal != nil {
		var ok bool
		*enumNum, ok = ast.AsInt32(n.IntVal, math.MinInt32, math.MaxInt32)
		if !ok {
			_ = p.handler.handleErrorWithPos(n.IntVal.Start(), "value %d is out of range: [%d,%d]",
				n.IntVal.Value(), math.MinInt32, math.MaxInt32)
		}
	}
	d := &EnumValDesc{
		Name:   n.Name.Val,
		Number: *enumNum,
	}
	*enumNum++
	p.putEnumValNode(d, n)
	return d
}

func (p *parseResult) asUnionDesc(n *ast.UnionDeclNode) *UnionDesc {
	d := &UnionDesc{
		Namespace: p.fd.Namespaces[len(p.fd.Namespaces)-1],
		Name:      n.Name.Val,
	}
	p.putUnionNode(d, n)
	for i := len(n.Decls) - 1; i >= 0; i-- {
		decl := n.Decls[i]
		d.Values = append(d.Values, p.asUnionVal(decl))
	}
	return d
}

func (p *parseResult) asUnionVal(n *ast.UnionValueNode) *UnionValDesc {
	name := ""
	if n.Name != nil {
		name = n.Name.Val
	}
	d := &UnionValDesc{Name: name, TypeName: string(n.TypeName.TypeName.Identifier())}
	p.putUnionValNode(d, n)
	return d
}

func (p *parseResult) asRPCDesc(n *ast.RPCDeclNode) *RPCDesc {
	d := &RPCDesc{
		Namespace: p.fd.Namespaces[len(p.fd.Namespaces)-1],
		Name:      n.Name.Val,
	}
	p.putRPCNode(d, n)
	for _, decl := range n.Methods {
		d.Methods = append(d.Methods, p.asMethodDesc(decl))
	}
	return d
}

func setStreaming(d *MethodDesc, v interface{}) {
	if v == ClientStreaming {
		d.ClientStreaming = true
	} else if v == ServerStreaming {
		d.ServerStreaming = true
	} else if v == BidiStreaming {
		d.ClientStreaming = true
		d.ServerStreaming = true
	}
}

func (p *parseResult) asMethodDesc(n *ast.RPCMethodNode) *MethodDesc {
	d := &MethodDesc{
		Name:       n.Name.Val,
		InputType:  string(n.ReqName.Identifier()),
		OutputType: string(n.RspName.Identifier()),
		Metadata:   p.asMetadataDesc(n.Metadata),
	}
	if d.Metadata != nil {
		v, ok := d.Metadata.KV[Streaming]
		if ok {
			setStreaming(d, v)
		}
	}
	p.putMethodNode(d, n)
	return d
}

func (p *parseResult) asMetadataDesc(n *ast.MetadataNode) *MetadataDesc {
	if n == nil {
		return nil
	}
	d := &MetadataDesc{KV: map[string]interface{}{}}
	for _, entry := range n.Entries {
		if entry.Value == nil {
			d.KV[entry.Key.Val] = nil
		} else {
			d.KV[entry.Key.Val] = entry.Value.Value()
		}
	}
	return d
}

func (p *parseResult) asFieldDesc(n *ast.FieldNode) *FieldDesc {
	d := &FieldDesc{
		Name:     n.Name.Val,
		TypeName: string(n.TypeName.TypeName.Identifier()),
		IsVector: n.TypeName.OpenBracket != nil && n.TypeName.CloseBracket != nil,
	}
	p.putFieldNode(d, n)
	return d
}

func (p *parseResult) addTableFields(d *TableDesc, fields []*ast.FieldNode) {
	for _, field := range fields {
		d.Fields = append(d.Fields, p.asFieldDesc(field))
	}
}

func (p *parseResult) addStructFields(d *StructDesc, fields []*ast.FieldNode) {
	for _, field := range fields {
		d.Fields = append(d.Fields, p.asFieldDesc(field))
	}
}

func (p *parseResult) putSchemaNode(d *SchemaDesc, n *ast.SchemaNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putTableNode(d *TableDesc, n *ast.TableDeclNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putStructNode(d *StructDesc, n *ast.StructDeclNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putEnumNode(d *EnumDesc, n *ast.EnumDeclNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putEnumValNode(d *EnumValDesc, n *ast.EnumValueNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putUnionNode(d *UnionDesc, n *ast.UnionDeclNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putUnionValNode(d *UnionValDesc, n *ast.UnionValueNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putRPCNode(d *RPCDesc, n *ast.RPCDeclNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putMethodNode(d *MethodDesc, n *ast.RPCMethodNode) {
	p.descToNode[d] = n
}

func (p *parseResult) putFieldNode(d *FieldDesc, n *ast.FieldNode) {
	p.descToNode[d] = n
}

func (p *parseResult) getSchemaNode(d *SchemaDesc) *ast.SchemaNode {
	if p.descToNode == nil {
		return nil
	}
	return p.descToNode[d].(*ast.SchemaNode)
}

func (p *parseResult) getFieldNode(d *FieldDesc) *ast.FieldNode {
	if p.descToNode == nil {
		return nil
	}
	return p.descToNode[d].(*ast.FieldNode)
}

func (p *parseResult) getMethodNode(d *MethodDesc) *ast.RPCMethodNode {
	if p.descToNode == nil {
		return nil
	}
	return p.descToNode[d].(*ast.RPCMethodNode)
}
