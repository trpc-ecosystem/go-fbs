package ast

// TypeNameNode represents a type name used in table/struct/enum definitions. Examples:
//
// enum Race:byte {..}     table Stat { id:string; }    struct Ability { id:uint(key); }
//           ^^^^                          ^^^^^^                           ^^^^
//         type name                     type name                       type name
//
// When enclosed inside brackets, it represents vector of types. Example:
//
// table Monster { inventory:[ubyte] (id: 5); }
//                           ^^^^^^^
//                        vector of ubyte
type TypeNameNode struct {
	compositeNode
	OpenBracket  *RuneNode
	TypeName     IdentLiteralElement
	CloseBracket *RuneNode
}

// NewTypeNameNode creates a type name node.
func NewTypeNameNode(openBracket *RuneNode, typeName IdentLiteralElement, closeBracket *RuneNode) *TypeNameNode {
	var children []Node
	if openBracket != nil {
		children = append(children, openBracket)
	}
	children = append(children, typeName)
	if closeBracket != nil {
		children = append(children, closeBracket)
	}
	return &TypeNameNode{
		compositeNode: compositeNode{children: children},
		OpenBracket:   openBracket,
		TypeName:      typeName,
		CloseBracket:  closeBracket,
	}
}
