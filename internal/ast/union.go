package ast

// UnionDeclNode represents a union in flatbuffers.
// Keyword is 'union'. Example:
//
// union Any { Monster, TestSimpleTableWithEnum, MyGame.Example2.Monster }
// union AnyUniqueAliases { M: Monster, TS: TestSimpleTableWithEnum, M2: MyGame.Example2.Monster }
type UnionDeclNode struct {
	compositeNode
	Keyword    *KeywordNode
	Name       *IdentNode
	Metadata   *MetadataNode
	OpenBrace  *RuneNode
	Decls      []*UnionValueNode
	CloseBrace *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*UnionDeclNode) AsDeclElement() {}

// UnionDeclOption provides functional options pattern.
type UnionDeclOption func(*UnionDeclNode)

// WithUnionKeyword create option to set enum keyword.
func WithUnionKeyword(keyword *KeywordNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.Keyword = keyword
		node.compositeNode.children = append(node.compositeNode.children, keyword)
	}
}

// WithUnionName creates option to set enum name.
func WithUnionName(name *IdentNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithUnionMetadata creates option to set enum metadata.
func WithUnionMetadata(metadata *MetadataNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithUnionOpenBrace creates option to set enum open brace.
func WithUnionOpenBrace(openBrace *RuneNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.OpenBrace = openBrace
		node.compositeNode.children = append(node.compositeNode.children, openBrace)
	}
}

// WithUnionDecls creates option to set enum decls.
func WithUnionDecls(decls []*UnionValueNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.Decls = decls
		for _, decl := range decls {
			node.compositeNode.children = append(node.compositeNode.children, decl)
		}
	}
}

// WithUnionCloseBrace creates option to set enum close brace.
func WithUnionCloseBrace(closeBrace *RuneNode) UnionDeclOption {
	return func(node *UnionDeclNode) {
		node.CloseBrace = closeBrace
		node.compositeNode.children = append(node.compositeNode.children, closeBrace)
	}
}

// NewUnionDeclNode creates union declaration node. Note: metadata could be nil.
func NewUnionDeclNode(opts ...UnionDeclOption) *UnionDeclNode {
	res := &UnionDeclNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

// UnionValueNode represents a field inside enum or union. Example:
//
// M: Monster
// TS: TestSimpleTableWithEnum
// M2: MyGame.Example2.Monster
type UnionValueNode struct {
	compositeNode
	Name     *IdentNode
	Colon    *RuneNode
	TypeName *TypeNameNode
}

// NewUnionValueNode creates value node for enum or union.
// 'name' and 'Colon' is optional. 'typeName' is necessary.
// 'name' is actually an alias for 'typeName'.
func NewUnionValueNode(name *IdentNode, colon *RuneNode, typeName *TypeNameNode) *UnionValueNode {
	var children []Node
	if colon != nil {
		children = append(children, name, colon)
	}
	children = append(children, typeName) // typeName is necessary.
	return &UnionValueNode{
		compositeNode: compositeNode{children: children},
		Name:          name,
		Colon:         colon,
		TypeName:      typeName,
	}
}
