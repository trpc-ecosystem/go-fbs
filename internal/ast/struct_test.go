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

func TestStruct(t *testing.T) {
	keyword := ast.NewIdentNode("keyword", ast.Token{}).ToKeyword()
	name := ast.NewIdentNode("Name", ast.Token{})
	openParen := ast.NewRuneNode('(', ast.Token{})
	closeParen := ast.NewRuneNode(')', ast.Token{})
	entryKey := ast.NewIdentNode("mykey", ast.Token{})
	entryValue := ast.NewIdentNode("myvalue", ast.Token{})
	colon := ast.NewRuneNode(':', ast.Token{})
	metadataEntry := ast.NewMetadataEntryNode(entryKey, colon, entryValue)
	entries := []*ast.MetadataEntryNode{metadataEntry}
	metadata := ast.NewMetadataNode(openParen, entries, closeParen)
	fieldName := ast.NewIdentNode("fieldName", ast.Token{})
	openBracket := ast.NewRuneNode('[', ast.Token{})
	closeBracket := ast.NewRuneNode(']', ast.Token{})
	nameNode := ast.NewIdentNode("TypeName", ast.Token{})
	typeName := ast.NewTypeNameNode(openBracket, nameNode, closeBracket)
	equal := ast.NewRuneNode('=', ast.Token{})
	scalar := ast.NewUintLiteralNode(10, ast.Token{})
	semicolon := ast.NewRuneNode(';', ast.Token{})
	var fieldOpts []ast.FieldOption
	fieldOpts = append(fieldOpts, ast.WithFieldName(fieldName))
	fieldOpts = append(fieldOpts, ast.WithFieldColon(colon))
	fieldOpts = append(fieldOpts, ast.WithFieldTypeName(typeName))
	fieldOpts = append(fieldOpts, ast.WithFieldEqual(equal))
	fieldOpts = append(fieldOpts, ast.WithFieldScalar(scalar))
	fieldOpts = append(fieldOpts, ast.WithFieldMetadata(metadata))
	fieldOpts = append(fieldOpts, ast.WithFieldSemicolon(semicolon))
	field := ast.NewFieldNode(fieldOpts...)
	fields := []*ast.FieldNode{field}
	openBrace := ast.NewRuneNode('{', ast.Token{})
	closeBrace := ast.NewRuneNode('}', ast.Token{})
	var opts []ast.StructDeclOption
	opts = append(opts, ast.WithStructKeyword(keyword))
	opts = append(opts, ast.WithStructName(name))
	opts = append(opts, ast.WithStructMetadata(metadata))
	opts = append(opts, ast.WithStructOpenBrace(openBrace))
	opts = append(opts, ast.WithStructFields(fields))
	opts = append(opts, ast.WithStructCloseBrace(closeBrace))
	structDecl := ast.NewStructDeclNode(opts...)
	structDecl.AsDeclElement()
	assert.Equal(t, keyword, structDecl.Keyword)
	assert.Equal(t, name, structDecl.Name)
	assert.Equal(t, metadata, structDecl.Metadata)
	assert.Equal(t, openBrace, structDecl.OpenBrace)
	assert.Equal(t, fields, structDecl.Fields)
	assert.Equal(t, closeBrace, structDecl.CloseBrace)
}
