package ast

import (
	"fmt"
	"strings"
)

// Identifier distinguishes itself from string literals.
type Identifier string

// IdentLiteralElement provides an interface for identifier literals.
type IdentLiteralElement interface {
	ValueNode
	Identifier() Identifier
}

var _ IdentLiteralElement = (*IdentNode)(nil)
var _ IdentLiteralElement = (*CompoundIdentNode)(nil)

// IdentNode represents all kinds of identifiers. Examples:
//
// anidentifier hahaha
type IdentNode struct {
	terminalNode
	Val string
}

// NewIdentNode creates terminal node identifier.
// This will be used by the lexer.
func NewIdentNode(val string, info Token) *IdentNode {
	return &IdentNode{
		terminalNode: info.asTerminalNode(),
		Val:          val,
	}
}

// Value implements ValueNode interface.
func (i *IdentNode) Value() interface{} {
	return Identifier(i.Val)
}

// Identifier implements IdentLiteralElement interface.
func (i *IdentNode) Identifier() Identifier {
	return Identifier(i.Val)
}

// ToKeyword converts the identifier node to keyword node.
func (i *IdentNode) ToKeyword() *KeywordNode {
	return (*KeywordNode)(i)
}

// CompoundIdentNode represents identifier containing dots. Example:
//
// app.server.service
type CompoundIdentNode struct {
	compositeNode
	LeadingDot *RuneNode
	Idents     []*IdentNode
	Dots       []*RuneNode
	Val        string
}

// NewCompoundIdentNode creates a compound identifier node.
func NewCompoundIdentNode(leadingDot *RuneNode, idents []*IdentNode, dots []*RuneNode) *CompoundIdentNode {
	checkDotNum(idents, dots)
	var children []Node
	var b strings.Builder
	if leadingDot != nil {
		children = append(children, leadingDot)
		b.WriteRune(leadingDot.Rune)
	}
	for i, id := range idents {
		if i > 0 {
			dot := dots[i-1]
			children = append(children, dot)
			b.WriteRune(dot.Rune)
		}
		children = append(children, id)
		b.WriteString(id.Val)
	}
	return &CompoundIdentNode{
		compositeNode: compositeNode{children: children},
		LeadingDot:    leadingDot,
		Idents:        idents,
		Dots:          dots,
		Val:           b.String(),
	}
}

func checkDotNum(idents []*IdentNode, dots []*RuneNode) {
	dl, il := len(dots), len(idents)
	if il == 0 {
		panic("should have at least one subpart")
	}
	if dl != il-1 {
		panic(fmt.Sprintf("%d idents needs %d dots, not %d", il, il-1, dl))
	}
}

// Value implements ValueNode interface.
func (c *CompoundIdentNode) Value() interface{} {
	return Identifier(c.Val)
}

// Identifier implements IdentLiteralElement.
func (c *CompoundIdentNode) Identifier() Identifier {
	return Identifier(c.Val)
}

// IdentList is a linked list of IdentNode, represents identifiers
// of the form 'mytestapp.mytestserver.mytestservice'
type IdentList struct {
	Ident *IdentNode
	Dot   *RuneNode
	Next  *IdentList
}

// ToIdentValueNode gathers together the linked list in IdentList to
// form a unified identifier.
func (i *IdentList) ToIdentValueNode(leadingDot *RuneNode) IdentLiteralElement {
	if i.Next == nil && leadingDot == nil {
		return i.Ident
	}
	length := lengthOfIdentList(i)
	idents := make([]*IdentNode, length)
	dots := make([]*RuneNode, length-1)
	for cur, idx := i, 0; cur != nil; cur, idx = cur.Next, idx+1 {
		idents[idx] = cur.Ident
		if cur.Dot != nil {
			dots[idx] = cur.Dot
		}
	}
	return NewCompoundIdentNode(leadingDot, idents, dots)
}

func lengthOfIdentList(i *IdentList) int {
	var length int
	for cur := i; cur != nil; cur = cur.Next {
		length++
	}
	return length
}

// KeywordNode represents a set of special identifiers
// which are reserved for particular usages.
type KeywordNode IdentNode
