package ast

import (
	"math"
)

// ValueNode should be implemented by all AST nodes with literal types.
type ValueNode interface {
	Node
	Value() interface{}
}

var _ ValueNode = (*IdentNode)(nil)
var _ ValueNode = (*CompoundIdentNode)(nil)
var _ ValueNode = (*StringLiteralNode)(nil)
var _ ValueNode = (*BoolLiteralNode)(nil)

var _ ValueNode = (*UintLiteralNode)(nil)
var _ ValueNode = (*PositiveUintLiteralNode)(nil)
var _ ValueNode = (*NegativeIntLiteralNode)(nil)

var _ ValueNode = (*FloatLiteralNode)(nil)
var _ ValueNode = (*SpecialFloatLiteralNode)(nil)
var _ ValueNode = (*SignedFloatLiteralNode)(nil)

// StringLiteralNode represents a string literal. Example:
//
// "a string literal"
type StringLiteralNode struct {
	terminalNode
	Val string
}

// NewStringLiteralNode creates terminal node for string literal.
// This is used by lexer.
func NewStringLiteralNode(val string, info Token) *StringLiteralNode {
	return &StringLiteralNode{
		terminalNode: info.asTerminalNode(),
		Val:          val,
	}
}

// Value implements ValueNode interface.
func (s *StringLiteralNode) Value() interface{} {
	return s.Val
}

// BoolLiteralNode represents boolean identifier. Examples:
//
// true false
type BoolLiteralNode struct {
	*KeywordNode
	Val bool
}

// NewBoolLiteralNode creates boolean node.
func NewBoolLiteralNode(name *KeywordNode) *BoolLiteralNode {
	return &BoolLiteralNode{
		KeywordNode: name,
		Val:         name.Val == "true",
	}
}

// Value implements ValueNode interface.
func (b *BoolLiteralNode) Value() interface{} {
	return b.Val
}

// IntValueNode represents either positive or negative integer literals.
type IntValueNode interface {
	ValueNode
	AsInt64() (int64, bool)
	AsUint64() (uint64, bool)
}

var _ IntValueNode = (*UintLiteralNode)(nil)
var _ IntValueNode = (*PositiveUintLiteralNode)(nil)
var _ IntValueNode = (*NegativeIntLiteralNode)(nil)

// AsInt32 returns int32 type value for IntValueNode.
func AsInt32(n IntValueNode, mi, mx int32) (int32, bool) {
	i, ok := n.AsInt64()
	if !ok {
		return 0, false
	}
	if i < int64(mi) || i > int64(mx) {
		return 0, false
	}
	return int32(i), true
}

// UintLiteralNode represents raw integer without sign.
type UintLiteralNode struct {
	terminalNode
	Val uint64
}

// NewUintLiteralNode creates a terminal uint literal node.
// This will be used by the lexer.
func NewUintLiteralNode(val uint64, info Token) *UintLiteralNode {
	return &UintLiteralNode{
		terminalNode: info.asTerminalNode(),
		Val:          val,
	}
}

// Value implements ValueNode interface.
func (u *UintLiteralNode) Value() interface{} {
	return u.Val
}

// AsInt64 implements IntValueNode interface.
func (u *UintLiteralNode) AsInt64() (int64, bool) {
	if u.Val > math.MaxInt64 {
		return 0, false
	}
	return int64(u.Val), true
}

// AsUint64 implements IntValueNode interface.
func (u *UintLiteralNode) AsUint64() (uint64, bool) {
	return u.Val, true
}

// PositiveUintLiteralNode represents integer with positive sign. Example:
//
// +78
type PositiveUintLiteralNode struct {
	compositeNode
	Plus *RuneNode
	Uint *UintLiteralNode
	Val  uint64
}

// NewPositiveUintLiteralNode creates a positive uint literal node
// out of sign and uint literal node.
func NewPositiveUintLiteralNode(sign *RuneNode, i *UintLiteralNode) *PositiveUintLiteralNode {
	children := []Node{sign, i}
	return &PositiveUintLiteralNode{
		compositeNode: compositeNode{children: children},
		Plus:          sign,
		Uint:          i,
		Val:           i.Val,
	}
}

// Value implements ValueNode interface.
func (p *PositiveUintLiteralNode) Value() interface{} {
	return p.Val
}

// AsInt64 implements IntValueNode interface.
func (p *PositiveUintLiteralNode) AsInt64() (int64, bool) {
	if p.Val > math.MaxInt64 {
		return 0, false
	}
	return int64(p.Val), true
}

// AsUint64 implements IntValueNode interface.
func (p *PositiveUintLiteralNode) AsUint64() (uint64, bool) {
	return p.Val, true
}

// NegativeIntLiteralNode represents integer with negative sign. Example:
//
// -78
type NegativeIntLiteralNode struct {
	compositeNode
	Minus *RuneNode
	Uint  *UintLiteralNode
	Val   int64
}

// NewNegativeIntLiteralNode creates negative int literal node
// out of sign and uint literal node.
func NewNegativeIntLiteralNode(sign *RuneNode, i *UintLiteralNode) *NegativeIntLiteralNode {
	children := []Node{sign, i}
	return &NegativeIntLiteralNode{
		compositeNode: compositeNode{children: children},
		Minus:         sign,
		Uint:          i,
		Val:           -int64(i.Val),
	}
}

// Value implements ValueNode interface.
func (n *NegativeIntLiteralNode) Value() interface{} {
	return n.Val
}

// AsInt64 implements IntValueNode interface.
func (n *NegativeIntLiteralNode) AsInt64() (int64, bool) {
	return n.Val, true
}

// AsUint64 implements IntValueNode interface.
func (n *NegativeIntLiteralNode) AsUint64() (uint64, bool) {
	if n.Val < 0 {
		return 0, false
	}
	return uint64(n.Val), true
}

// FloatValueNode interface should be implemented by all
// floating point literals.
type FloatValueNode interface {
	ValueNode
	AsFloat() float64
}

var _ FloatValueNode = (*FloatLiteralNode)(nil)
var _ FloatValueNode = (*SpecialFloatLiteralNode)(nil)
var _ FloatValueNode = (*SignedFloatLiteralNode)(nil)

// FloatLiteralNode represents a raw floating point number without sign.
// Examples:
//
// 6.828 6.824
type FloatLiteralNode struct {
	terminalNode
	Val float64
}

// NewFloatLiteralNode creates terminal float literal node.
// This will be used by the lexer.
func NewFloatLiteralNode(val float64, info Token) *FloatLiteralNode {
	return &FloatLiteralNode{
		terminalNode: info.asTerminalNode(),
		Val:          val,
	}
}

// Value implements ValueNode interface.
func (f *FloatLiteralNode) Value() interface{} {
	return f.Val
}

// AsFloat implements FloatValueNode interface.
func (f *FloatLiteralNode) AsFloat() float64 {
	return f.Val
}

// SpecialFloatLiteralNode represents either "inf", "infinity" or "nan".
type SpecialFloatLiteralNode struct {
	*KeywordNode
	Val float64
}

// NewSpecialFloatLiteralNode creates node for special floating point such as
// "inf", "infinity" or "nan"
func NewSpecialFloatLiteralNode(name *KeywordNode) *SpecialFloatLiteralNode {
	var val float64
	if name.Val == "inf" || name.Val == "infinity" {
		val = math.Inf(1)
	} else {
		val = math.NaN()
	}
	return &SpecialFloatLiteralNode{
		KeywordNode: name,
		Val:         val,
	}
}

// Value implements ValueNode interface.
func (s *SpecialFloatLiteralNode) Value() interface{} {
	return s.Val
}

// AsFloat implements FloatValueNode interface.
func (s *SpecialFloatLiteralNode) AsFloat() float64 {
	return s.Val
}

// SignedFloatLiteralNode represents floating point values with signs. Examples:
//
// +3.1415 -0.618 +inf -nan
type SignedFloatLiteralNode struct {
	compositeNode
	Sign  *RuneNode
	Float FloatValueNode
	Val   float64
}

// NewSignedFloatLiteralNode creates signed floating point using given sign
// and value.
func NewSignedFloatLiteralNode(sign *RuneNode, f FloatValueNode) *SignedFloatLiteralNode {
	children := []Node{sign, f}
	val := f.AsFloat()
	if sign.Rune == '-' {
		val = -val
	}
	return &SignedFloatLiteralNode{
		compositeNode: compositeNode{children: children},
		Sign:          sign,
		Float:         f,
		Val:           val,
	}
}

// Value implements ValueNode interface.
func (s *SignedFloatLiteralNode) Value() interface{} {
	return s.Val
}

// AsFloat implements FloatValueNode interface.
func (s *SignedFloatLiteralNode) AsFloat() float64 {
	return s.Val
}
