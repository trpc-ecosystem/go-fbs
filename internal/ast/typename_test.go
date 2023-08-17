package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestTypeName(t *testing.T) {
	openBracket := ast.NewRuneNode('[', ast.Token{})
	closeBracket := ast.NewRuneNode(']', ast.Token{})
	nameNode := ast.NewIdentNode("TypeName", ast.Token{})
	typeName := ast.NewTypeNameNode(openBracket, nameNode, closeBracket)
	assert.Equal(t, openBracket, typeName.OpenBracket)
	assert.Equal(t, closeBracket, typeName.CloseBracket)
	assert.Equal(t, nameNode, typeName.TypeName)
}
