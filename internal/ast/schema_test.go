package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestSchema(t *testing.T) {
	keyword := ast.NewIdentNode("keyword", ast.Token{}).ToKeyword()
	name := ast.NewIdentNode("Name", ast.Token{})
	semicolon := ast.NewRuneNode(';', ast.Token{})
	str := ast.NewStringLiteralNode("a string literal node", ast.Token{})
	includeDecl := ast.NewIncludeNode(keyword, str, semicolon)
	namespaceDecl := ast.NewNamespaceDeclNode(keyword, name, semicolon)
	rootDecl := ast.NewRootDeclNode(keyword, name, semicolon)
	fileExtDecl := ast.NewFileExtDeclNode(keyword, str, semicolon)
	fileIdentDecl := ast.NewFileIdentDeclNode(keyword, str, semicolon)
	attrDecl := ast.NewAttrDeclNode(keyword, str, semicolon)
	namespaceDecl.AsDeclElement()
	rootDecl.AsDeclElement()
	fileExtDecl.AsDeclElement()
	fileIdentDecl.AsDeclElement()
	attrDecl.AsDeclElement()
	includes := []*ast.IncludeNode{includeDecl}
	decls := []ast.DeclElement{namespaceDecl, rootDecl, fileExtDecl, fileIdentDecl, attrDecl}
	schema := ast.NewSchemaNode(includes, decls)
	assert.Equal(t, includes, schema.Includes)
	assert.Equal(t, decls, schema.Decls)
}
