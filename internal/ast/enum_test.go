package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestEnum(t *testing.T) {
	keyword := ast.NewIdentNode("keyword", ast.Token{}).ToKeyword()
	name := ast.NewIdentNode("Name", ast.Token{})
	colon := ast.NewRuneNode(':', ast.Token{})
	nameNode := ast.NewIdentNode("TypeName", ast.Token{})
	openBracket := ast.NewRuneNode('[', ast.Token{})
	closeBracket := ast.NewRuneNode(']', ast.Token{})
	typeName := ast.NewTypeNameNode(openBracket, nameNode, closeBracket)
	openParen := ast.NewRuneNode('(', ast.Token{})
	closeParen := ast.NewRuneNode(')', ast.Token{})
	entryKey := ast.NewIdentNode("mykey", ast.Token{})
	entryValue := ast.NewIdentNode("myvalue", ast.Token{})
	metadataEntry := ast.NewMetadataEntryNode(entryKey, colon, entryValue)
	entries := []*ast.MetadataEntryNode{metadataEntry}
	metadata := ast.NewMetadataNode(openParen, entries, closeParen)
	openBrace := ast.NewRuneNode('{', ast.Token{})
	closeBrace := ast.NewRuneNode('}', ast.Token{})
	equal := ast.NewRuneNode('=', ast.Token{})
	intVal := ast.NewUintLiteralNode(10, ast.Token{})
	enumVal := ast.NewEnumValueNode(name, equal, intVal)
	enums := []*ast.EnumValueNode{enumVal}
	var opts []ast.EnumDeclOption
	opts = append(opts, ast.WithEnumKeyword(keyword))
	opts = append(opts, ast.WithEnumName(name))
	opts = append(opts, ast.WithEnumColon(colon))
	opts = append(opts, ast.WithEnumTypeName(typeName))
	opts = append(opts, ast.WithEnumMetadata(metadata))
	opts = append(opts, ast.WithEnumOpenBrace(openBrace))
	opts = append(opts, ast.WithEnumDecls(enums))
	opts = append(opts, ast.WithEnumCloseBrace(closeBrace))
	// Construct a enum decl node.
	enumDecl := ast.NewEnumDeclNode(opts...)
	enumDecl.AsDeclElement()
	assert.Equal(t, keyword, enumDecl.Keyword)
	assert.Equal(t, name, enumDecl.Name)
	assert.Equal(t, colon, enumDecl.Colon)
	assert.Equal(t, typeName, enumDecl.TypeName)
	assert.Equal(t, metadata, enumDecl.Metadata)
	assert.Equal(t, openBrace, enumDecl.OpenBrace)
	assert.Equal(t, enums, enumDecl.Decls)
	assert.Equal(t, closeBrace, enumDecl.CloseBrace)
}
