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

// RPCDeclNode represents a rpc service statement. Example:
//
//	rpc_service MonsterStorage {
//	  Store(Monster):Stat (streaming: "none");
//	  Retrieve(Stat):Monster (streaming: "server", idempotent);
//	  GetMaxHitPoint(Monster):Stat (streaming: "client");
//	  GetMinMaxHitPoints(Monster):Stat (streaming: "bidi");
//	}
type RPCDeclNode struct {
	compositeNode
	Keyword    *KeywordNode
	Name       *IdentNode
	OpenBrace  *RuneNode
	Methods    []*RPCMethodNode
	CloseBrace *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*RPCDeclNode) AsDeclElement() {}

// NewRPCDeclNode creates a RPC service declaration node.
func NewRPCDeclNode(keyword *KeywordNode, name *IdentNode, openBrace *RuneNode,
	methods []*RPCMethodNode, closeBrace *RuneNode) *RPCDeclNode {
	var children []Node
	children = append(children, keyword, name, openBrace)
	for _, method := range methods {
		children = append(children, method)
	}
	return &RPCDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		OpenBrace:     openBrace,
		Methods:       methods,
		CloseBrace:    closeBrace,
	}
}

// RPCMethodNode represents a method inside an rpc service. Examples:
//
//	Store(Monster):Stat (streaming: "none");
//	Retrieve(Stat):Monster (streaming: "server", idempotent);
//	GetMaxHitPoint(Monster):Stat (streaming: "client");
//	GetMinMaxHitPoints(Monster):Stat (streaming: "bidi");
//	SayHello(HelloRequest):HelloReply;
//	SayManyHellos(HelloRequest):HelloReply (streaming: "server");
type RPCMethodNode struct {
	compositeNode
	Name       *IdentNode
	OpenParen  *RuneNode
	ReqName    IdentLiteralElement
	CloseParen *RuneNode
	Colon      *RuneNode
	RspName    IdentLiteralElement
	Metadata   *MetadataNode
	Semicolon  *RuneNode
}

// MethodOption provides functional options pattern.
type MethodOption func(*RPCMethodNode)

// WithMethodName creates option to set method name.
func WithMethodName(name *IdentNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithMethodOpenParen creates option to set method open parenthesis.
func WithMethodOpenParen(openParen *RuneNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.OpenParen = openParen
		node.compositeNode.children = append(node.compositeNode.children, openParen)
	}
}

// WithMethodReqName creates option to set method request name.
func WithMethodReqName(name IdentLiteralElement) MethodOption {
	return func(node *RPCMethodNode) {
		node.ReqName = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithMethodCloseParen creates option to set method close parenthesis.
func WithMethodCloseParen(closeParen *RuneNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.CloseParen = closeParen
		node.compositeNode.children = append(node.compositeNode.children, closeParen)
	}
}

// WithMethodColon creates option to set method colon rune.
func WithMethodColon(colon *RuneNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.Colon = colon
		node.compositeNode.children = append(node.compositeNode.children, colon)
	}
}

// WithMethodRspName creates option to set method response name.
func WithMethodRspName(name IdentLiteralElement) MethodOption {
	return func(node *RPCMethodNode) {
		node.RspName = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithMethodMetadata creates option to set method metadata.
func WithMethodMetadata(metadata *MetadataNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithMethodSemicolon creates option to set method semicolon rune.
func WithMethodSemicolon(semicolon *RuneNode) MethodOption {
	return func(node *RPCMethodNode) {
		node.Semicolon = semicolon
		node.compositeNode.children = append(node.compositeNode.children, semicolon)
	}
}

// NewRPCMethodNode creates a RPC method node.
func NewRPCMethodNode(opts ...MethodOption) *RPCMethodNode {
	res := &RPCMethodNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}
