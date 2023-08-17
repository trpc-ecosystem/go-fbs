package ast_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestLiterals(t *testing.T) {
	// Test string literal.
	stringLiteralNode := ast.NewStringLiteralNode("a string literal node", ast.Token{})
	assert.Equal(t, "a string literal node", stringLiteralNode.Value())
	// Test integer values.
	uintLiteralNode := ast.NewUintLiteralNode(10, ast.Token{})
	assert.Equal(t, uint64(10), uintLiteralNode.Value())
	int64Literal, ok := uintLiteralNode.AsInt64()
	assert.True(t, ok)
	assert.Equal(t, int64(10), int64Literal)
	largeIntLiteralNode := ast.NewUintLiteralNode(uint64(math.MaxInt64)+1, ast.Token{})
	int64Literal, ok = largeIntLiteralNode.AsInt64()
	assert.False(t, ok)
	assert.Equal(t, int64(0), int64Literal)
	uint64Literal, ok := largeIntLiteralNode.AsUint64()
	assert.Equal(t, uint64(math.MaxInt64)+1, uint64Literal)
	assert.True(t, ok)
	// Test positive integer values.
	plusRuneNode := ast.NewRuneNode('+', ast.Token{})
	positiveUintLiteralNode := ast.NewPositiveUintLiteralNode(plusRuneNode, uintLiteralNode)
	assert.Equal(t, uint64(10), positiveUintLiteralNode.Value())
	int64PositiveUintLiteralNode, ok := positiveUintLiteralNode.AsInt64()
	assert.Equal(t, int64(10), int64PositiveUintLiteralNode)
	assert.True(t, ok)
	largePositiveUintLiteralNode := ast.NewPositiveUintLiteralNode(plusRuneNode, largeIntLiteralNode)
	int64Literal, ok = largePositiveUintLiteralNode.AsInt64()
	assert.Equal(t, int64(0), int64Literal)
	assert.False(t, ok)
	uint64Literal, ok = largePositiveUintLiteralNode.AsUint64()
	assert.Equal(t, uint64(math.MaxInt64)+1, uint64Literal)
	assert.True(t, ok)
	int32Literal, ok := ast.AsInt32(largeIntLiteralNode, math.MinInt32, math.MaxInt32)
	assert.Equal(t, int32(0), int32Literal)
	assert.False(t, ok)
	int32Literal, ok = ast.AsInt32(uintLiteralNode, 0, 0)
	assert.Equal(t, int32(0), int32Literal)
	assert.False(t, ok)
	int32Literal, ok = ast.AsInt32(uintLiteralNode, 0, 20)
	assert.Equal(t, int32(10), int32Literal)
	assert.True(t, ok)
	// Test negative integer values.
	minusRuneNode := ast.NewRuneNode('-', ast.Token{})
	negativeIntLiteralNode := ast.NewNegativeIntLiteralNode(minusRuneNode, uintLiteralNode)
	assert.Equal(t, -int64(10), negativeIntLiteralNode.Value())
	int64Negative, ok := negativeIntLiteralNode.AsInt64()
	assert.Equal(t, int64Negative, -int64(10))
	assert.True(t, ok)
	uint64Negative, ok := negativeIntLiteralNode.AsUint64()
	assert.Equal(t, uint64(0), uint64Negative)
	assert.False(t, ok)
	zeroUintLiteralNode := ast.NewUintLiteralNode(0, ast.Token{})
	zeroNegativeNode := ast.NewNegativeIntLiteralNode(minusRuneNode, zeroUintLiteralNode)
	uint64Negative, ok = zeroNegativeNode.AsUint64()
	assert.Equal(t, uint64(0), uint64Negative)
	assert.True(t, ok)
	// Test special float numbers.
	infNode := ast.NewIdentNode("inf", ast.Token{})
	specialNode := ast.NewSpecialFloatLiteralNode(infNode.ToKeyword())
	assert.Equal(t, math.Inf(1), specialNode.Value())
	assert.Equal(t, math.Inf(1), specialNode.AsFloat())
	nanNode := ast.NewIdentNode("nan", ast.Token{})
	specialNode = ast.NewSpecialFloatLiteralNode(nanNode.ToKeyword())
	assert.True(t, math.IsNaN(specialNode.Value().(float64)))
	assert.True(t, math.IsNaN(specialNode.AsFloat()))
	// Test signed float numbers.
	floatLiteralNode := ast.NewFloatLiteralNode(10.6, ast.Token{})
	assert.Equal(t, 10.6, floatLiteralNode.Value())
	assert.Equal(t, 10.6, floatLiteralNode.AsFloat())
	positiveFloat := ast.NewSignedFloatLiteralNode(plusRuneNode, floatLiteralNode)
	assert.Equal(t, 10.6, positiveFloat.Value())
	assert.Equal(t, 10.6, positiveFloat.AsFloat())
	negativeFloat := ast.NewSignedFloatLiteralNode(minusRuneNode, floatLiteralNode)
	assert.Equal(t, -10.6, negativeFloat.Value())
	assert.Equal(t, -10.6, negativeFloat.AsFloat())
	// Test boolean literal.
	trueNode := ast.NewIdentNode("true", ast.Token{})
	boolLiteralNode := ast.NewBoolLiteralNode(trueNode.ToKeyword())
	assert.Equal(t, true, boolLiteralNode.Value())
}
