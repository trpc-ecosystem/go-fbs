package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestRPC(t *testing.T) {
	keyword := ast.NewIdentNode("keyword", ast.Token{}).ToKeyword()
	name := ast.NewIdentNode("Name", ast.Token{})
	colon := ast.NewRuneNode(':', ast.Token{})
	openParen := ast.NewRuneNode('(', ast.Token{})
	closeParen := ast.NewRuneNode(')', ast.Token{})
	entryKey := ast.NewIdentNode("mykey", ast.Token{})
	entryValue := ast.NewIdentNode("myvalue", ast.Token{})
	metadataEntry := ast.NewMetadataEntryNode(entryKey, colon, entryValue)
	entries := []*ast.MetadataEntryNode{metadataEntry}
	metadata := ast.NewMetadataNode(openParen, entries, closeParen)
	openBrace := ast.NewRuneNode('{', ast.Token{})
	closeBrace := ast.NewRuneNode('}', ast.Token{})
	semicolon := ast.NewRuneNode(';', ast.Token{})
	var opts []ast.MethodOption
	opts = append(opts, ast.WithMethodName(name))
	opts = append(opts, ast.WithMethodOpenParen(openParen))
	opts = append(opts, ast.WithMethodReqName(name))
	opts = append(opts, ast.WithMethodCloseParen(closeParen))
	opts = append(opts, ast.WithMethodColon(colon))
	opts = append(opts, ast.WithMethodRspName(name))
	opts = append(opts, ast.WithMethodMetadata(metadata))
	opts = append(opts, ast.WithMethodSemicolon(semicolon))
	rpcMethod := ast.NewRPCMethodNode(opts...)
	methods := []*ast.RPCMethodNode{rpcMethod}
	rpcDecl := ast.NewRPCDeclNode(keyword, name, openBrace, methods, closeBrace)
	rpcDecl.AsDeclElement()
	assert.Equal(t, name, rpcDecl.Name)
	assert.Equal(t, openBrace, rpcDecl.OpenBrace)
	assert.Equal(t, methods, rpcDecl.Methods)
	assert.Equal(t, closeBrace, rpcDecl.CloseBrace)
}
