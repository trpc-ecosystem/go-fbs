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

func TestMetadata(t *testing.T) {
	colon := ast.NewRuneNode(':', ast.Token{})
	openParen := ast.NewRuneNode('(', ast.Token{})
	closeParen := ast.NewRuneNode(')', ast.Token{})
	entryKey := ast.NewIdentNode("mykey", ast.Token{})
	entryValue := ast.NewIdentNode("myvalue", ast.Token{})
	metadataEntry := ast.NewMetadataEntryNode(entryKey, colon, entryValue)
	entries := []*ast.MetadataEntryNode{metadataEntry}
	metadata := ast.NewMetadataNode(openParen, entries, closeParen)
	assert.Equal(t, openParen, metadata.OpenParen)
	assert.Equal(t, closeParen, metadata.CloseParen)
	assert.Equal(t, entries, metadata.Entries)
}
