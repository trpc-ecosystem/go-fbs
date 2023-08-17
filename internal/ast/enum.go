package ast

// EnumDeclNode represents enumerations in flatbuffers.
// Keyword is 'enum'. Example:
//
// enum Color:ubyte (bit_flags) {
//   Red = 0, // color Red = (1u << 0)
//   /// \brief color Green
//   /// Green is bit_flag with value (1u << 1)
//   Green,
//   /// \brief color Blue (1u << 3)
//   Blue = 3,
// }
type EnumDeclNode struct {
	compositeNode
	Keyword    *KeywordNode
	Name       *IdentNode
	Colon      *RuneNode
	TypeName   *TypeNameNode
	Metadata   *MetadataNode
	OpenBrace  *RuneNode
	Decls      []*EnumValueNode
	CloseBrace *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*EnumDeclNode) AsDeclElement() {}

// EnumDeclOption provides functional options pattern.
type EnumDeclOption func(*EnumDeclNode)

// WithEnumKeyword create option to set enum keyword.
func WithEnumKeyword(keyword *KeywordNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.Keyword = keyword
		node.compositeNode.children = append(node.compositeNode.children, keyword)
	}
}

// WithEnumName creates option to set enum name.
func WithEnumName(name *IdentNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithEnumColon creates option to set enum colon.
func WithEnumColon(colon *RuneNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.Colon = colon
		node.compositeNode.children = append(node.compositeNode.children, colon)
	}
}

// WithEnumTypeName creates option to set enum typename.
func WithEnumTypeName(typeName *TypeNameNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.TypeName = typeName
		node.compositeNode.children = append(node.compositeNode.children, typeName)
	}
}

// WithEnumMetadata creates option to set enum metadata.
func WithEnumMetadata(metadata *MetadataNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithEnumOpenBrace creates option to set enum open brace.
func WithEnumOpenBrace(openBrace *RuneNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.OpenBrace = openBrace
		node.compositeNode.children = append(node.compositeNode.children, openBrace)
	}
}

// WithEnumDecls creates option to set enum decls.
func WithEnumDecls(decls []*EnumValueNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.Decls = decls
		for _, decl := range decls {
			node.compositeNode.children = append(node.compositeNode.children, decl)
		}
	}
}

// WithEnumCloseBrace creates option to set enum close brace.
func WithEnumCloseBrace(closeBrace *RuneNode) EnumDeclOption {
	return func(node *EnumDeclNode) {
		node.CloseBrace = closeBrace
		node.compositeNode.children = append(node.compositeNode.children, closeBrace)
	}
}

// NewEnumDeclNode creates enum node. Note: metadata could be nil.
func NewEnumDeclNode(opts ...EnumDeclOption) *EnumDeclNode {
	res := &EnumDeclNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

// EnumValueNode represents a field inside enum or union. Example:
//
// Red = 0
// Green
type EnumValueNode struct {
	compositeNode
	Name   *IdentNode
	Equal  *RuneNode
	IntVal IntValueNode
}

// NewEnumValueNode creates value node for enum or union.
func NewEnumValueNode(name *IdentNode, equal *RuneNode, intVal IntValueNode) *EnumValueNode {
	var children []Node
	children = append(children, name)
	if equal != nil {
		children = append(children, equal, intVal)
	}
	return &EnumValueNode{
		compositeNode: compositeNode{children: children},
		Name:          name,
		Equal:         equal,
		IntVal:        intVal,
	}
}
