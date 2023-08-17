package ast

// SchemaNode represents the whole .fbs file. It is the
// root node of the AST (Abstract Syntax Tree).
type SchemaNode struct {
	compositeNode
	Includes []*IncludeNode
	Decls    []DeclElement
}

// NewSchemaNode creates a new schema node.
func NewSchemaNode(includes []*IncludeNode, decls []DeclElement) *SchemaNode {
	// Calculate the number of children to preallocate the slice.
	// Since we may have a lot of `includes` and `declarations`, we'd
	// better preallocate here.
	numChildren := len(includes) + len(decls)
	children := make([]Node, 0, numChildren)
	for _, incl := range includes {
		children = append(children, incl)
	}
	for _, decl := range decls {
		children = append(children, decl)
	}
	return &SchemaNode{
		compositeNode: compositeNode{children: children},
		Includes:      includes,
		Decls:         decls,
	}
}

// IncludeNode represents an include statement. Example:
//
// include "include_test.fbs";
type IncludeNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      *StringLiteralNode
	Semicolon *RuneNode
}

// NewIncludeNode create an include node.
func NewIncludeNode(keyword *KeywordNode, name *StringLiteralNode, semicolon *RuneNode) *IncludeNode {
	var children []Node
	children = append(children, keyword, name, semicolon)
	return &IncludeNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}

// DeclElement is an interface implemented by all AST nodes that can
// serve as top-level declarations in the file.
type DeclElement interface {
	Node
	AsDeclElement()
}

var _ DeclElement = (*NamespaceDeclNode)(nil)
var _ DeclElement = (*TableDeclNode)(nil)
var _ DeclElement = (*StructDeclNode)(nil)
var _ DeclElement = (*EnumDeclNode)(nil)
var _ DeclElement = (*UnionDeclNode)(nil)
var _ DeclElement = (*RootDeclNode)(nil)
var _ DeclElement = (*FileExtDeclNode)(nil)
var _ DeclElement = (*FileIdentDeclNode)(nil)
var _ DeclElement = (*AttrDeclNode)(nil)
var _ DeclElement = (*RPCDeclNode)(nil)

// NamespaceDeclNode represents a namespace statement,
// just like package statement in protobuf. Example:
//
// namespace MyAPP.Example;
type NamespaceDeclNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      IdentLiteralElement
	Semicolon *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*NamespaceDeclNode) AsDeclElement() {}

// NewNamespaceDeclNode creates namespace node.
func NewNamespaceDeclNode(keyword *KeywordNode, name IdentLiteralElement,
	semicolon *RuneNode) *NamespaceDeclNode {
	children := []Node{keyword, name, semicolon}
	return &NamespaceDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}

// RootDeclNode represents the root type of a .fbs file. Example:
//
// root_type Monster;
type RootDeclNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      *IdentNode
	Semicolon *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*RootDeclNode) AsDeclElement() {}

// NewRootDeclNode creates root declaration node.
func NewRootDeclNode(keyword *KeywordNode, name *IdentNode, semicolon *RuneNode) *RootDeclNode {
	children := []Node{keyword, name, semicolon}
	return &RootDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}

// FileExtDeclNode represents a file extension statement. Example:
//
// file_extension "mon";
type FileExtDeclNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      *StringLiteralNode
	Semicolon *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*FileExtDeclNode) AsDeclElement() {}

// NewFileExtDeclNode creates file extension declaration node.
func NewFileExtDeclNode(keyword *KeywordNode, name *StringLiteralNode, semicolon *RuneNode) *FileExtDeclNode {
	children := []Node{keyword, name, semicolon}
	return &FileExtDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}

// FileIdentDeclNode represents a file identifier statement. Example:
//
// file_identifier "MONS";
type FileIdentDeclNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      *StringLiteralNode
	Semicolon *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*FileIdentDeclNode) AsDeclElement() {}

// NewFileIdentDeclNode creates file identifier declaration node.
func NewFileIdentDeclNode(keyword *KeywordNode, name *StringLiteralNode, semicolon *RuneNode) *FileIdentDeclNode {
	children := []Node{keyword, name, semicolon}
	return &FileIdentDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}

// AttrDeclNode represents an attribute statement. Example:
//
// attribute "priority";
type AttrDeclNode struct {
	compositeNode
	Keyword   *KeywordNode
	Name      *StringLiteralNode
	Semicolon *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*AttrDeclNode) AsDeclElement() {}

// NewAttrDeclNode creates attribute declaration node.
func NewAttrDeclNode(keyword *KeywordNode, name *StringLiteralNode, semicolon *RuneNode) *AttrDeclNode {
	children := []Node{keyword, name, semicolon}
	return &AttrDeclNode{
		compositeNode: compositeNode{children: children},
		Keyword:       keyword,
		Name:          name,
		Semicolon:     semicolon,
	}
}
