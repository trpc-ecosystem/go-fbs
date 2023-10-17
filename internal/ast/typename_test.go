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

package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestTypeName(t *testing.T) {
	openBracket := ast.NewRuneNode('[', ast.Token{})
	closeBracket := ast.NewRuneNode(']', ast.Token{})
	nameNode := ast.NewIdentNode("TypeName", ast.Token{})
	typeName := ast.NewTypeNameNode(openBracket, nameNode, closeBracket)
	assert.Equal(t, openBracket, typeName.OpenBracket)
	assert.Equal(t, closeBracket, typeName.CloseBracket)
	assert.Equal(t, nameNode, typeName.TypeName)
}
