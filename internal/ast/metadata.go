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

// MetadataNode represents metadata information of type or identifier.
// Usually when you see parentheses, you encounter metadata. Example:
//
//	table Referrable {
//	  id:ulong(key, hash:"fnv1a_64");
//	          ^^^^^^^^^^^^^^^^^^^^^^
//	  //        these are metadata
//	}
type MetadataNode struct {
	compositeNode
	OpenParen  *RuneNode
	Entries    []*MetadataEntryNode
	CloseParen *RuneNode
}

// NewMetadataNode creates metadata node.
func NewMetadataNode(openParen *RuneNode, entries []*MetadataEntryNode, closeParen *RuneNode) *MetadataNode {
	var children []Node
	children = append(children, openParen)
	for _, entry := range entries {
		children = append(children, entry)
	}
	return &MetadataNode{
		compositeNode: compositeNode{children: children},
		OpenParen:     openParen,
		Entries:       entries,
		CloseParen:    closeParen,
	}
}

// MetadataEntryNode represents an entry of metadata. It can be a key-value pair
// or just key. Example:
//
//	table Referable {
//	  id:ulong(key, hash:"fnv1a_64");
//	           ^^^  ^^^^^^^^^^^^^^^
//	  //     entry1     entry2
//	}
type MetadataEntryNode struct {
	compositeNode
	Key   *IdentNode
	Colon *RuneNode
	Value ValueNode
}

// NewMetadataEntryNode creates an entry for metadata node.
func NewMetadataEntryNode(key *IdentNode, colon *RuneNode, value ValueNode) *MetadataEntryNode {
	var children []Node
	children = append(children, key)
	if colon != nil && value != nil {
		children = append(children, colon, value)
	}
	return &MetadataEntryNode{
		compositeNode: compositeNode{children: children},
		Key:           key,
		Colon:         colon,
		Value:         value,
	}
}
