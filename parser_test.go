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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyParse(t *testing.T) {
	filenames := []string{
		"this_file_does_not_exist.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.NotNil(t, err)
	assert.Nil(t, fds)
}

func TestSimpleParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/simple_test1.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err, "expected nil, got: %v", err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	assert.Equal(t, "./fbsfiles/simple_test1.fbs", fd.Name)
	assert.Equal(t, []string{"", "mynamespace"}, fd.Namespaces)
	assert.Equal(t, []string{"myattr"}, fd.Attrs)
	// Check enums.
	assert.Equal(t, 1, len(fd.Enums))
	enum := fd.Enums[0]
	assert.Equal(t, "mynamespace", enum.Namespace)
	assert.Equal(t, "MyEnum", enum.Name)
	assert.Equal(t, 3, len(enum.Values))
	// union value list are parsed backwards.
	red, green, blue := enum.Values[0], enum.Values[1], enum.Values[2]
	assert.Equal(t, "Red", red.Name)
	assert.Equal(t, "Green", green.Name)
	assert.Equal(t, "Blue", blue.Name)
	assert.Equal(t, int32(0), red.Number)
	assert.Equal(t, int32(1), green.Number)
	assert.Equal(t, int32(3), blue.Number)
	// Check unions.
	assert.Equal(t, 1, len(fd.Unions))
	union := fd.Unions[0]
	assert.Equal(t, "mynamespace", union.Namespace)
	assert.Equal(t, "Any", union.Name)
	assert.Equal(t, "MyTable1", union.Values[0].TypeName)
	assert.Equal(t, "MyTable2", union.Values[1].TypeName)
	// Check tables.
	assert.Equal(t, 2, len(fd.Tables))
	table1, table2 := fd.Tables[0], fd.Tables[1]
	assert.Equal(t, "mynamespace", table1.Namespace)
	assert.Equal(t, "MyTable1", table1.Name)
	assert.Equal(t, 2, len(table1.Fields))
	assert.Equal(t, "myfield1", table1.Fields[0].Name)
	assert.Equal(t, "int", table1.Fields[0].TypeName)
	assert.Equal(t, "myfield2", table1.Fields[1].Name)
	assert.Equal(t, "int", table1.Fields[1].TypeName)
	assert.Equal(t, "mynamespace", table2.Namespace)
	assert.Equal(t, "MyTable2", table2.Name)
	assert.Equal(t, 2, len(table2.Fields))
	assert.Equal(t, "myfield1", table2.Fields[0].Name)
	assert.Equal(t, "short", table2.Fields[0].TypeName)
	assert.Equal(t, "myfield2", table2.Fields[1].Name)
	assert.Equal(t, "short", table2.Fields[1].TypeName)
	// Check structs.
	assert.Equal(t, 1, len(fd.Structs))
	s := fd.Structs[0]
	assert.Equal(t, "mynamespace", s.Namespace)
	assert.Equal(t, "MyEmptyStruct", s.Name)
	assert.Nil(t, s.Fields)
	// Check rpc services.
	assert.Equal(t, 1, len(fd.RPCs))
	rpc := fd.RPCs[0]
	assert.Equal(t, "mynamespace", rpc.Namespace)
	assert.Equal(t, "MyService1", rpc.Name)
	assert.Equal(t, 5, len(rpc.Methods))
	// Check service methods.
	method := rpc.Methods[0]
	assert.Equal(t, "MyMethod1", method.Name)
	assert.Equal(t, ".mynamespace.MyTable1", method.InputType)
	assert.Equal(t, ".mynamespace.MyTable2", method.OutputType)
	assert.False(t, method.ClientStreaming)
	assert.False(t, method.ServerStreaming)
	assert.Nil(t, method.Metadata)
	method = rpc.Methods[1]
	assert.Equal(t, "MyMethod2", method.Name)
	assert.False(t, method.ClientStreaming)
	assert.False(t, method.ServerStreaming)
	method = rpc.Methods[2]
	assert.Equal(t, "MyMethod3", method.Name)
	assert.True(t, method.ClientStreaming)
	assert.False(t, method.ServerStreaming)
	method = rpc.Methods[3]
	assert.Equal(t, "MyMethod4", method.Name)
	assert.False(t, method.ClientStreaming)
	assert.True(t, method.ServerStreaming)
	method = rpc.Methods[4]
	assert.Equal(t, "MyMethod5", method.Name)
	assert.True(t, method.ClientStreaming)
	assert.True(t, method.ServerStreaming)
	// Check file identifier, file extension and root_type.
	assert.Equal(t, "FIDT", fd.FileIdent)
	assert.Equal(t, "fid", fd.FileExt)
	assert.Equal(t, "MyTable1", fd.Root)
}

func TestSimpleNamespaceParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/simple_test2.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check namespaces.
	assert.Equal(t, 4, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "rpc.ns1", fd.Namespaces[1])
	assert.Equal(t, "rpc.ns2", fd.Namespaces[2])
	assert.Equal(t, "rpc.ns3", fd.Namespaces[3])
	// Check structs in different namespaces.
	assert.Equal(t, 2, len(fd.Structs))
	s := fd.Structs[0]
	assert.Equal(t, "rpc.ns1", s.Namespace)
	assert.Equal(t, "MyStruct1", s.Name)
	assert.Nil(t, s.Fields)
	s = fd.Structs[1]
	assert.Equal(t, "rpc.ns2", s.Namespace)
	assert.Equal(t, "MyStruct1", s.Name)
	assert.Nil(t, s.Fields)
	// Check tables in different namespaces.
	assert.Equal(t, 2, len(fd.Tables))
	table := fd.Tables[0]
	assert.Equal(t, "rpc.ns2", table.Namespace)
	assert.Equal(t, "MyTable1", table.Name)
	assert.Equal(t, 2, len(table.Fields))
	field := table.Fields[0]
	assert.Equal(t, ".rpc.ns1.MyEnum1", field.TypeName)
	field = table.Fields[1]
	assert.Equal(t, ".rpc.ns2.MyStruct1", field.TypeName)
	table = fd.Tables[1]
	assert.Equal(t, "rpc.ns3", table.Namespace)
	assert.Equal(t, "MyTable2", table.Name)
	assert.Equal(t, 1, len(table.Fields))
	field = table.Fields[0]
	assert.Equal(t, ".rpc.ns1.MyStruct1", field.TypeName)
	// Check rpc server method typename resolution.
	assert.Equal(t, 1, len(fd.RPCs))
	rpc := fd.RPCs[0]
	assert.Equal(t, 1, len(rpc.Methods))
	method := rpc.Methods[0]
	assert.Equal(t, ".rpc.ns2.MyTable1", method.InputType)
	assert.Equal(t, ".rpc.ns3.MyTable2", method.OutputType)
}

func TestNamespaceParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/namespace_test/namespace_test2.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	assert.Equal(t, 3, len(fd.Tables))
	// Check TableInFirstNS field typename resolution.
	table := fd.Tables[0]
	assert.Equal(t, 4, len(table.Fields))
	assert.Equal(t, ".NamespaceA.NamespaceB.TableInNestedNS", table.Fields[0].TypeName)
	assert.Equal(t, ".NamespaceA.NamespaceB.EnumInNestedNS", table.Fields[1].TypeName)
	assert.Equal(t, ".NamespaceA.NamespaceB.UnionInNestedNS", table.Fields[2].TypeName)
	assert.Equal(t, ".NamespaceA.NamespaceB.StructInNestedNS", table.Fields[3].TypeName)
	// Check TableInC field typename resolution.
	table = fd.Tables[1]
	assert.Equal(t, 2, len(table.Fields))
	assert.Equal(t, ".NamespaceA.TableInFirstNS", table.Fields[0].TypeName)
	assert.Equal(t, ".NamespaceA.SecondTableInA", table.Fields[1].TypeName)
	// Check SecondTableInA field typename resolution.
	table = fd.Tables[2]
	assert.Equal(t, 1, len(table.Fields))
	assert.Equal(t, ".NamespaceC.TableInC", table.Fields[0].TypeName)
}

func TestAnnotatedParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/annotated_test.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	assert.Equal(t, 1, len(fd.Includes))
	assert.Equal(t, 2, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "rpc.app.server", fd.Namespaces[1])
	assert.Equal(t, 1, len(fd.Attrs))
	assert.Equal(t, "myattr", fd.Attrs[0])
	assert.Equal(t, 1, len(fd.Enums))
	assert.Equal(t, 2, len(fd.Unions))
	assert.Equal(t, 2, len(fd.Tables))
	assert.Equal(t, 1, len(fd.Structs))
	assert.Equal(t, 1, len(fd.RPCs))
	assert.Equal(t, "FIDT", fd.FileIdent)
	assert.Equal(t, "fid", fd.FileExt)
}

func TestMonsterParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/monster_test.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check includes.
	assert.Equal(t, 1, len(fd.Includes))
	include := fd.Includes[0]
	assert.Equal(t, "include_test1.fbs", include)
	// Check namespaces.
	assert.Equal(t, 4, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "MyGame", fd.Namespaces[1])
	assert.Equal(t, "MyGame.Example2", fd.Namespaces[2])
	assert.Equal(t, "MyGame.Example", fd.Namespaces[3])
	// Check tables.
	assert.Equal(t, 7, len(fd.Tables))
	assert.Equal(t, "InParentNamespace", fd.Tables[0].Name)
	assert.Equal(t, "Monster", fd.Tables[1].Name)
	assert.Equal(t, "TestSimpleTableWithEnum", fd.Tables[2].Name)
	assert.Equal(t, "Stat", fd.Tables[3].Name)
	assert.Equal(t, "Referrable", fd.Tables[4].Name)
	assert.Equal(t, "Monster", fd.Tables[5].Name)
	assert.Equal(t, "TypeAliases", fd.Tables[6].Name)
	// Check attributes.
	assert.Equal(t, 1, len(fd.Attrs))
	assert.Equal(t, "priority", fd.Attrs[0])
	// Check enums.
	assert.Equal(t, 2, len(fd.Enums))
	assert.Equal(t, "Color", fd.Enums[0].Name)
	assert.Equal(t, "Race", fd.Enums[1].Name)
	// Check unions.
	assert.Equal(t, 3, len(fd.Unions))
	assert.Equal(t, "Any", fd.Unions[0].Name)
	assert.Equal(t, "AnyUniqueAliases", fd.Unions[1].Name)
	assert.Equal(t, "AnyAmbiguousAliases", fd.Unions[2].Name)
	// Check rpc_services.
	assert.Equal(t, 1, len(fd.RPCs))
	assert.Equal(t, "MonsterStorage", fd.RPCs[0].Name)
	// Check file identifier.
	assert.Equal(t, "MONS", fd.FileIdent)
	assert.Equal(t, "mon", fd.FileExt)
}

func TestMonsterExtraParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/monster_extra.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check namespaces.
	assert.Equal(t, 2, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "MyGame", fd.Namespaces[1])
	// Check tables.
	assert.Equal(t, 1, len(fd.Tables))
	table := fd.Tables[0]
	assert.Equal(t, "MonsterExtra", table.Name)
	// Check fields in table.
	assert.Equal(t, 11, len(table.Fields))
	fs := table.Fields
	assert.Equal(t, "d0", fs[0].Name)
	assert.Equal(t, "d1", fs[1].Name)
	assert.Equal(t, "d2", fs[2].Name)
	assert.Equal(t, "d3", fs[3].Name)
	assert.Equal(t, "f0", fs[4].Name)
	assert.Equal(t, "f1", fs[5].Name)
	assert.Equal(t, "f2", fs[6].Name)
	assert.Equal(t, "f3", fs[7].Name)
	assert.Equal(t, "dvec", fs[8].Name)
	assert.True(t, fs[8].IsVector)
	assert.Equal(t, "fvec", fs[9].Name)
	assert.True(t, fs[9].IsVector)
	assert.Equal(t, "deprec", fs[10].Name)
	// Check file identifier.
	assert.Equal(t, "MONE", fd.FileIdent)
	// Check file extension.
	assert.Equal(t, "mon", fd.FileExt)
}

func TestReflectionParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/reflection.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check namespaces.
	assert.Equal(t, 2, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "reflection", fd.Namespaces[1])
	// Check enums.
	assert.Equal(t, 2, len(fd.Enums))
	assert.Equal(t, "BaseType", fd.Enums[0].Name)
	assert.Equal(t, 19, len(fd.Enums[0].Values))
	assert.Equal(t, "AdvancedFeatures", fd.Enums[1].Name)
	assert.Equal(t, 4, len(fd.Enums[1].Values))
	// Check tables.
	assert.Equal(t, 9, len(fd.Tables))
	assert.Equal(t, "Type", fd.Tables[0].Name)
	assert.Equal(t, "KeyValue", fd.Tables[1].Name)
	assert.Equal(t, "EnumVal", fd.Tables[2].Name)
	assert.Equal(t, "Enum", fd.Tables[3].Name)
	assert.Equal(t, "Field", fd.Tables[4].Name)
	assert.Equal(t, "Object", fd.Tables[5].Name)
	assert.Equal(t, "RPCCall", fd.Tables[6].Name)
	assert.Equal(t, "Service", fd.Tables[7].Name)
	assert.Equal(t, "Schema", fd.Tables[8].Name)
	// Check file identifier.
	assert.Equal(t, "BFBS", fd.FileIdent)
	assert.Equal(t, "bfbs", fd.FileExt)
}

func TestOptionalScalarsParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/optional_scalars.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check namespaces.
	assert.Equal(t, 2, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	assert.Equal(t, "optional_scalars", fd.Namespaces[1])
	// Check enums.
	assert.Equal(t, 1, len(fd.Enums))
	assert.Equal(t, "OptionalByte", fd.Enums[0].Name)
	assert.Equal(t, 3, len(fd.Enums[0].Values))
	// Check tables.
	assert.Equal(t, 1, len(fd.Tables))
	assert.Equal(t, "ScalarStuff", fd.Tables[0].Name)
	assert.Equal(t, 36, len(fd.Tables[0].Fields))
	// Check file identifier.
	assert.Equal(t, "NULL", fd.FileIdent)
	// Check file extension.
	assert.Equal(t, "mon", fd.FileExt)
}

func TestUnionVectorParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/union_vector.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
	fd := fds[0]
	// Check namespaces.
	assert.Equal(t, 1, len(fd.Namespaces))
	assert.Equal(t, "", fd.Namespaces[0])
	// Check tables.
	assert.Equal(t, 2, len(fd.Tables))
	assert.Equal(t, "Attacker", fd.Tables[0].Name)
	assert.Equal(t, "Movie", fd.Tables[1].Name)

	// Check structs
	assert.Equal(t, 2, len(fd.Structs))
	assert.Equal(t, "Rapunzel", fd.Structs[0].Name)
	assert.Equal(t, "BookReader", fd.Structs[1].Name)
	// Check unions.
	assert.Equal(t, 1, len(fd.Unions))
	assert.Equal(t, "Character", fd.Unions[0].Name)
	// Check file identifier.
	assert.Equal(t, "MOVI", fd.FileIdent)
}

func TestDuplicateFilesParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/custom_test.fbs",
		"./fbsfiles/custom_test.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(fds))
	assert.Equal(t, fds[0], fds[1])
}

func TestErrorParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/error_test/parse_test1.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "syntax error")
	assert.Nil(t, fds)
}

func TestEmptyFileParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/error_test/empty_test.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
}

func TestIncludeErrorFileParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/error_test/parse_test2.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "syntax error")
	assert.Nil(t, fds)
}

func TestNonrecursiveParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/error_test/parse_test2.fbs",
	}
	p := NewParser()
	p.SetRecursive(false)
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
}

func TestCustomParse(t *testing.T) {
	filenames := []string{
		"./fbsfiles/custom_test.fbs",
	}
	p := NewParser()
	fds, err := p.ParseFiles(filenames...)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(fds))
}
