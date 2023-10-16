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

// Node is the interface that will be implemented by all nodes in the AST.
type Node interface {
	Start() *Position
	End() *Position
}

// TerminalNode should be implemented by all terminal node types.
type TerminalNode interface {
	Node
	RawText() string
}

var _ TerminalNode = (*StringLiteralNode)(nil)
var _ TerminalNode = (*UintLiteralNode)(nil)
var _ TerminalNode = (*FloatLiteralNode)(nil)
var _ TerminalNode = (*IdentNode)(nil)
var _ TerminalNode = (*RuneNode)(nil)

// terminalNode represents all terminal tokens. It records
// position information and implements the TerminalNode
// interface.
type terminalNode struct {
	posRange PosRange
	raw      string
}

// Start implements Node interface, providing the start of the span.
func (t *terminalNode) Start() *Position {
	return &t.posRange.Start
}

// End implements Node interface, providing the end of the span.
func (t *terminalNode) End() *Position {
	return &t.posRange.End
}

// RawText returns the underlying test of the terminal node.
func (t *terminalNode) RawText() string {
	return t.raw
}

// compositeNode represents all nodes that are not terminals.
// Typically it has children.
type compositeNode struct {
	children []Node
}

// Start implements Node interface.
func (c *compositeNode) Start() *Position {
	return c.children[0].Start()
}

// End implements Node interface.
func (c *compositeNode) End() *Position {
	return c.children[len(c.children)-1].End()
}

// Children returns children of a compositeNode.
func (c *compositeNode) Children() []Node {
	return c.children
}

// Token stores information of a token from lexer.
type Token struct {
	PosRange        // Location of the token in the source file.
	RawText  string // Raw text of the token.
}

// asTerminalNode create a terminalNode out of a Token.
func (t *Token) asTerminalNode() terminalNode {
	return terminalNode{
		posRange: t.PosRange,
		raw:      t.RawText,
	}
}

// RuneNode represents a single rune type value in Go. Examples:
//
// '=' ';' ':' '{' '}' '\\' '/' '?' '.'
type RuneNode struct {
	terminalNode
	Rune rune
}

// NewRuneNode creates a node representing a rune, which is a terminal node.
func NewRuneNode(r rune, info Token) *RuneNode {
	return &RuneNode{
		terminalNode: info.asTerminalNode(),
		Rune:         r,
	}
}
