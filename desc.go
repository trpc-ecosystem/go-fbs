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
	"strings"

	"trpc.group/trpc-go/fbs/internal/ast"
)

// Desc provides an interface for all descriptors in flatbuffers.
type Desc interface{ FbsDesc() }

var _ Desc = (*SchemaDesc)(nil)
var _ Desc = (*TableDesc)(nil)
var _ Desc = (*StructDesc)(nil)
var _ Desc = (*EnumDesc)(nil)
var _ Desc = (*EnumValDesc)(nil)
var _ Desc = (*UnionDesc)(nil)
var _ Desc = (*UnionValDesc)(nil)
var _ Desc = (*RPCDesc)(nil)
var _ Desc = (*MethodDesc)(nil)

var _ Desc = (*FieldDesc)(nil)
var _ Desc = (*MetadataDesc)(nil)

// NamespaceDesc provides an interface for descriptors that store their own namespaces.
type NamespaceDesc interface{ GetNamespace() string }

var _ NamespaceDesc = (*TableDesc)(nil)
var _ NamespaceDesc = (*StructDesc)(nil)
var _ NamespaceDesc = (*EnumDesc)(nil)
var _ NamespaceDesc = (*UnionDesc)(nil)
var _ NamespaceDesc = (*RPCDesc)(nil)

// TableStructDesc provides an interface for table and struct descriptors.
type TableStructDesc interface {
	Desc
	NamespaceDesc
	GetName() string
	GetFields() []*FieldDesc
}

var _ TableStructDesc = (*TableDesc)(nil)
var _ TableStructDesc = (*StructDesc)(nil)

// SchemaDesc describes a complete .fbs file.
type SchemaDesc struct {
	// Schema stores the root node of the AST.
	Schema *ast.SchemaNode
	// Name stores the .fbs file name.
	Name string
	// Namespace of schema will change in the course of processing each decl, these will be stored
	// in this slice. Decls of table/struct/enum/union will set their own namespaces to be the last
	// namespace in this slice.
	// See parseResult.createSchemaDescriptor.
	Namespaces []string
	// Root stores root_type declaration in flatbuffers file.
	Root string
	// FileExt stores file_extension declaration in flatbuffers file.
	FileExt string
	// FileIdent stores file_identifier declaration in flatbuffers file.
	FileIdent string
	// Attrs stores attributes definitions in flatbuffers file.
	Attrs []string
	// Includes stores all included file names.
	Includes []string
	// Dependencies stores all descriptors corresponding to included files.
	Dependencies []*SchemaDesc
	// Tables stores all table descriptors.
	Tables []*TableDesc
	// Structs stores all struct descriptors.
	Structs []*StructDesc
	// Enums stores all enum descriptors.
	Enums []*EnumDesc
	// Unions stores all union descriptors.
	Unions []*UnionDesc
	// RPCs stores all rpc service descriptors.
	RPCs []*RPCDesc
}

// FbsDesc implements Desc interface.
func (SchemaDesc) FbsDesc() {}

// TableDesc describes the structure of table in flatbuffers.
type TableDesc struct {
	Schema    *SchemaDesc  // Schema stores the descriptor that contains this table.
	Namespace string       // Namespace will be set as the current namespace of the schema node.
	Name      string       // Name is the name of table.
	Fields    []*FieldDesc // Fields list the fields of table.
}

// FbsDesc implements Desc interface.
func (TableDesc) FbsDesc() {}

// GetNamespace implements NamespaceDesc interface.
func (t *TableDesc) GetNamespace() string {
	return t.Namespace
}

// GetName implements TableStructDesc interface.
func (t *TableDesc) GetName() string {
	return t.Name
}

// GetFields implements TableStructDesc interface.
func (t *TableDesc) GetFields() []*FieldDesc {
	return t.Fields
}

// StructDesc describes the structure of struct in flatbuffers.
type StructDesc struct {
	Namespace string       // Namespace will be set as schema's namespace.
	Name      string       // Name is the name of struct.
	Fields    []*FieldDesc // Fields lists the fields of the struct.
}

// FbsDesc implements Desc.
func (StructDesc) FbsDesc() {}

// GetNamespace implements NamespaceDesc interface.
func (s *StructDesc) GetNamespace() string {
	return s.Namespace
}

// GetName implements TableStructDesc interface.
func (s *StructDesc) GetName() string {
	return s.Name
}

// GetFields implements TableStructDesc interface.
func (s *StructDesc) GetFields() []*FieldDesc {
	return s.Fields
}

// FieldDesc describes the structure of field in flatbuffers.
type FieldDesc struct {
	Name     string
	TypeName string
	IsVector bool // [typename] is a vector of typename.
}

// FbsDesc implements Desc interface.
func (FieldDesc) FbsDesc() {}

// EnumDesc describes the structure of enum in flatbuffers.
type EnumDesc struct {
	Namespace string // Namespace will be set as the current namespace of schema.
	Name      string // Name is the name of this enum.
	Values    []*EnumValDesc
}

// FbsDesc implements Desc interface.
func (EnumDesc) FbsDesc() {}

// GetNamespace implements NamespaceDesc interface.
func (e *EnumDesc) GetNamespace() string {
	return e.Namespace
}

// EnumValDesc describes the structure of enum value in flatbuffers.
type EnumValDesc struct {
	Name   string
	Number int32
}

// FbsDesc implements Desc interface.
func (EnumValDesc) FbsDesc() {}

// UnionDesc describes the structure of union in flatbuffers.
type UnionDesc struct {
	Namespace string // Namespace will be set as the current namespace of schema.
	Name      string
	Values    []*UnionValDesc
}

// FbsDesc implements Desc interface.
func (UnionDesc) FbsDesc() {}

// GetNamespace implements NamespaceDesc interface.
func (u *UnionDesc) GetNamespace() string {
	return u.Namespace
}

// UnionValDesc describes the structure of union value in flatbuffers.
type UnionValDesc struct {
	Name     string
	TypeName string
}

// FbsDesc implements Desc interface.
func (UnionValDesc) FbsDesc() {}

// RPCDesc describes the structure of rpc_service in flatbuffers.
type RPCDesc struct {
	Namespace string // Namespace will be set as schema's namespace.
	Name      string
	Methods   []*MethodDesc
}

// FbsDesc implements Desc interface.
func (RPCDesc) FbsDesc() {}

// GetNamespace implements NamespaceDesc interface.
func (r *RPCDesc) GetNamespace() string {
	return r.Namespace
}

// MethodDesc describes the structure of method in flatbuffers.
type MethodDesc struct {
	Name            string
	InputType       string
	InputTypeDesc   *TableDesc
	OutputType      string
	OutputTypeDesc  *TableDesc
	ClientStreaming bool
	ServerStreaming bool
	Metadata        *MetadataDesc
}

// FbsDesc implements Desc interface.
func (MethodDesc) FbsDesc() {}

// MetadataDesc describes the structure of metadata in flatbuffers.
type MetadataDesc struct {
	KV map[string]interface{}
}

// FbsDesc implements Desc interface.
func (MetadataDesc) FbsDesc() {}

// isType returns whether the given descriptor is a valid type.
func isType(d Desc) bool {
	switch d.(type) {
	case *TableDesc, *StructDesc, *EnumDesc, *EnumValDesc, *UnionDesc:
		return true
	}
	return false
}

// descType returns the corresponding string representation of given descriptor.
func descType(d Desc) string {
	switch d := d.(type) {
	case *TableDesc:
		return "table"
	case *StructDesc:
		return "struct"
	case *EnumDesc:
		return "enum"
	case *UnionDesc:
		return "union"
	case *FieldDesc:
		return "field"
	case *EnumValDesc:
		return "enum value"
	case *UnionValDesc:
		return "union value"
	case *RPCDesc:
		return "rpc"
	case *MethodDesc:
		return "method"
	case *SchemaDesc:
		return "schema"
	case *MetadataDesc:
		return "metadata"
	default:
		return fmt.Sprintf("%T", d)
	}
}

// getAllNamespaces will generate a list of prefixes out of a list
// of namespaces.
//
// Input: ["", "namespace1", "rpc.app.server"]
// Output:
//
//	map[string]struct{} {
//	   "": struct{},
//	   "namespace1": struct{},
//		"rpc": struct{},
//		"rpc.app": struct{},
//	   "rpc.app.server": struct{},
//	}
func getAllNamespaces(namespaces []string) map[string]struct{} {
	if namespaces == nil {
		return nil
	}
	allnss := map[string]struct{}{}
	for _, namespace := range namespaces {
		var offs int
		for {
			pos := strings.IndexByte(namespace[offs:], '.')
			if pos == -1 {
				break
			}
			allnss[namespace[:offs+pos]] = struct{}{}
			offs += pos + 1
		}
		allnss[namespace] = struct{}{}
	}
	return allnss
}

// getPrefix creates a prefix for a given namespace descriptor.
func getPrefix(d NamespaceDesc) string {
	if d.GetNamespace() == "" {
		return ""
	}
	return d.GetNamespace() + "."
}
