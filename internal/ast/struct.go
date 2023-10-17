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

package ast

// StructDeclNode represents struct definition in flatbuffers using 'struct'. Example:
//
// struct Test { a:short; b:byte; }
type StructDeclNode struct {
	compositeNode
	Keyword    *KeywordNode
	Name       *IdentNode
	Metadata   *MetadataNode
	OpenBrace  *RuneNode
	Fields     []*FieldNode
	CloseBrace *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*StructDeclNode) AsDeclElement() {}

// StructDeclOption provides functional options pattern.
type StructDeclOption func(*StructDeclNode)

// WithStructKeyword create option to set struct keyword.
func WithStructKeyword(keyword *KeywordNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.Keyword = keyword
		node.compositeNode.children = append(node.compositeNode.children, keyword)
	}
}

// WithStructName creates option to set struct name.
func WithStructName(name *IdentNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithStructMetadata creates option to set struct metadata.
func WithStructMetadata(metadata *MetadataNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithStructOpenBrace creates option to set struct open brace.
func WithStructOpenBrace(openBrace *RuneNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.OpenBrace = openBrace
		node.compositeNode.children = append(node.compositeNode.children, openBrace)
	}
}

// WithStructFields creates option to set struct fields.
func WithStructFields(fields []*FieldNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.Fields = fields
		for _, decl := range fields {
			node.compositeNode.children = append(node.compositeNode.children, decl)
		}
	}
}

// WithStructCloseBrace creates option to set struct close brace.
func WithStructCloseBrace(closeBrace *RuneNode) StructDeclOption {
	return func(node *StructDeclNode) {
		node.CloseBrace = closeBrace
		node.compositeNode.children = append(node.compositeNode.children, closeBrace)
	}
}

// NewStructDeclNode creates struct declaration node. Note: metadata could be nil.
func NewStructDeclNode(opts ...StructDeclOption) *StructDeclNode {
	res := &StructDeclNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}
