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

func TestIdent(t *testing.T) {
	// Test identifier node.
	identNode := ast.NewIdentNode("identifierx", ast.Token{})
	assert.Equal(t, ast.Identifier("identifierx"), identNode.Value())
	assert.Equal(t, ast.Identifier("identifierx"), identNode.Identifier())
	runeNode := ast.NewRuneNode(';', ast.Token{})
	assert.Equal(t, ';', runeNode.Rune)
	// Test compound identifier node.
	leadingDot := ast.NewRuneNode('.', ast.Token{})
	id1 := ast.NewIdentNode("id1", ast.Token{})
	id2 := ast.NewIdentNode("id2", ast.Token{})
	idents := []*ast.IdentNode{id1, id2}
	dot := ast.NewRuneNode('.', ast.Token{})
	dots := []*ast.RuneNode{dot}
	compoundIdentNode := ast.NewCompoundIdentNode(leadingDot, idents, dots)
	assert.Equal(t, ast.Identifier(".id1.id2"), compoundIdentNode.Value())
	assert.Equal(t, ast.Identifier(".id1.id2"), compoundIdentNode.Identifier())
	// Test identifier list.
	l2 := &ast.IdentList{id2, nil, nil}
	l1 := &ast.IdentList{id1, dot, l2}
	ident := l1.ToIdentValueNode(leadingDot)
	assert.Equal(t, ast.Identifier(".id1.id2"), ident.Value())
	ident = l2.ToIdentValueNode(nil)
	assert.Equal(t, ast.Identifier("id2"), ident.Value())
}
