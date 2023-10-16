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

func TestNode(t *testing.T) {
	start := ast.Position{
		Filename: "filename",
		Line:     12,
		Col:      13,
		Offset:   5,
	}
	end := ast.Position{
		Filename: "filename",
		Line:     12,
		Col:      17,
		Offset:   5,
	}
	tokenInfo := ast.Token{
		PosRange: ast.PosRange{Start: start, End: end},
		RawText:  "ident",
	}
	name := ast.NewIdentNode("ident", tokenInfo)
	assert.Equal(t, "filename:12:13", name.Start().String())
	assert.Equal(t, "filename:12:17", name.End().String())
	assert.Equal(t, "ident", name.RawText())
	colon := ast.NewRuneNode(':', ast.Token{})
	metadata := ast.NewMetadataEntryNode(name, colon, name)
	assert.Equal(t, "filename:12:13", metadata.Start().String())
	assert.Equal(t, "filename:12:17", metadata.End().String())
	children := metadata.Children()
	assert.Equal(t, 3, len(children))
}
