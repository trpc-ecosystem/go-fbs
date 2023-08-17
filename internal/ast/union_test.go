package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestUnion(t *testing.T) {
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
	unionVal := ast.NewUnionValueNode(name, colon, typeName)
	unions := []*ast.UnionValueNode{unionVal}
	var opts []ast.UnionDeclOption
	opts = append(opts, ast.WithUnionKeyword(keyword))
	opts = append(opts, ast.WithUnionName(name))
	opts = append(opts, ast.WithUnionMetadata(metadata))
	opts = append(opts, ast.WithUnionOpenBrace(openBrace))
	opts = append(opts, ast.WithUnionDecls(unions))
	opts = append(opts, ast.WithUnionCloseBrace(closeBrace))
	// Construct a union decl node.
	unionDecl := ast.NewUnionDeclNode(opts...)
	unionDecl.AsDeclElement()
	assert.Equal(t, keyword, unionDecl.Keyword)
	assert.Equal(t, name, unionDecl.Name)
	assert.Equal(t, metadata, unionDecl.Metadata)
	assert.Equal(t, openBrace, unionDecl.OpenBrace)
	assert.Equal(t, unions, unionDecl.Decls)
	assert.Equal(t, closeBrace, unionDecl.CloseBrace)
}
