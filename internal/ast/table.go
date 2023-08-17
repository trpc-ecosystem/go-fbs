package ast

// TableDeclNode represents table definition in flatbuffers using 'table'. Example:
//
// table HelloRequest {
//   name:string;
// }
//
type TableDeclNode struct {
	compositeNode
	Keyword    *KeywordNode
	Name       *IdentNode
	Metadata   *MetadataNode
	OpenBrace  *RuneNode
	Fields     []*FieldNode
	CloseBrace *RuneNode
}

// AsDeclElement implements DeclElement interface.
func (*TableDeclNode) AsDeclElement() {}

// TableDeclOption provides functional options pattern.
type TableDeclOption func(*TableDeclNode)

// WithTableKeyword create option to set table keyword.
func WithTableKeyword(keyword *KeywordNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.Keyword = keyword
		node.compositeNode.children = append(node.compositeNode.children, keyword)
	}
}

// WithTableName creates option to set table name.
func WithTableName(name *IdentNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithTableMetadata creates option to set table metadata.
func WithTableMetadata(metadata *MetadataNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithTableOpenBrace creates option to set table open brace.
func WithTableOpenBrace(openBrace *RuneNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.OpenBrace = openBrace
		node.compositeNode.children = append(node.compositeNode.children, openBrace)
	}
}

// WithTableFields creates option to set table fields.
func WithTableFields(fields []*FieldNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.Fields = fields
		for _, decl := range fields {
			node.compositeNode.children = append(node.compositeNode.children, decl)
		}
	}
}

// WithTableCloseBrace creates option to set table close brace.
func WithTableCloseBrace(closeBrace *RuneNode) TableDeclOption {
	return func(node *TableDeclNode) {
		node.CloseBrace = closeBrace
		node.compositeNode.children = append(node.compositeNode.children, closeBrace)
	}
}

// NewTableDeclNode creates table declaration node. Note: metadata could be nil.
func NewTableDeclNode(opts ...TableDeclOption) *TableDeclNode {
	res := &TableDeclNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

// FieldNode represents a single field inside table/struct declaration.
type FieldNode struct {
	compositeNode
	Name      *IdentNode
	Colon     *RuneNode
	TypeName  *TypeNameNode
	Equal     *RuneNode
	Scalar    ValueNode
	Metadata  *MetadataNode
	Semicolon *RuneNode
}

// FieldOption provides functional options pattern.
type FieldOption func(*FieldNode)

// WithFieldName creates option to set field name.
func WithFieldName(name *IdentNode) FieldOption {
	return func(node *FieldNode) {
		node.Name = name
		node.compositeNode.children = append(node.compositeNode.children, name)
	}
}

// WithFieldColon creates option to set field colon.
func WithFieldColon(colon *RuneNode) FieldOption {
	return func(node *FieldNode) {
		node.Colon = colon
		node.compositeNode.children = append(node.compositeNode.children, colon)
	}
}

// WithFieldTypeName creates option to set field typename.
func WithFieldTypeName(typeName *TypeNameNode) FieldOption {
	return func(node *FieldNode) {
		node.TypeName = typeName
		node.compositeNode.children = append(node.compositeNode.children, typeName)
	}
}

// WithFieldEqual creates option to set field equal rune.
func WithFieldEqual(equal *RuneNode) FieldOption {
	return func(node *FieldNode) {
		node.Equal = equal
		node.compositeNode.children = append(node.compositeNode.children, equal)
	}
}

// WithFieldScalar creates option to set field scalar.
func WithFieldScalar(scalar ValueNode) FieldOption {
	return func(node *FieldNode) {
		node.Scalar = scalar
		node.compositeNode.children = append(node.compositeNode.children, scalar)
	}
}

// WithFieldMetadata creates option to set field metadata.
func WithFieldMetadata(metadata *MetadataNode) FieldOption {
	return func(node *FieldNode) {
		node.Metadata = metadata
		node.compositeNode.children = append(node.compositeNode.children, metadata)
	}
}

// WithFieldSemicolon creates option to set field semicolon rune.
func WithFieldSemicolon(semicolon *RuneNode) FieldOption {
	return func(node *FieldNode) {
		node.Semicolon = semicolon
		node.compositeNode.children = append(node.compositeNode.children, semicolon)
	}
}

// NewFieldNode creates a field node. Note: metadata could be nil.
func NewFieldNode(opts ...FieldOption) *FieldNode {
	res := &FieldNode{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}
