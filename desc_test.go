package fbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDesc(t *testing.T) {
	schemaDesc := &SchemaDesc{}
	schemaDesc.FbsDesc()
	tableDesc := &TableDesc{}
	tableDesc.FbsDesc()
	structDesc := &StructDesc{}
	structDesc.FbsDesc()
	enumDesc := &EnumDesc{}
	enumDesc.FbsDesc()
	enumValDesc := &EnumValDesc{}
	enumValDesc.FbsDesc()
	unionDesc := &UnionDesc{}
	unionDesc.FbsDesc()
	unionValDesc := &UnionValDesc{}
	unionValDesc.FbsDesc()
	rpcDesc := &RPCDesc{}
	rpcDesc.FbsDesc()
	methodDesc := &RPCDesc{}
	methodDesc.FbsDesc()
	fieldDesc := &FieldDesc{}
	fieldDesc.FbsDesc()
	metadataDesc := &MethodDesc{}
	metadataDesc.FbsDesc()
	name := "mynamespace"
	tableDesc = &TableDesc{Namespace: name}
	assert.Equal(t, name, tableDesc.GetNamespace())
	structDesc = &StructDesc{Namespace: name}
	assert.Equal(t, name, structDesc.GetNamespace())
	enumDesc = &EnumDesc{Namespace: name}
	assert.Equal(t, name, enumDesc.GetNamespace())
	unionDesc = &UnionDesc{Namespace: name}
	assert.Equal(t, name, unionDesc.GetNamespace())
	rpcDesc = &RPCDesc{Namespace: name}
	assert.Equal(t, name, rpcDesc.GetNamespace())
}
