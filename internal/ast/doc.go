// Package ast provides definitions for nodes' types of
// the constructed abstract syntax tree.
//
// It provides:
// 1. definitions for terminal nodes
// 2. constructions for terminal nodes
// 3. definitions for non-terminal nodes
// 4. constructions for non-terminal nodes
//
// `fbs.y` and its generated file `fbs.y.go` will use 1,3,4 to construct
// non-terminal nodes using terminals and non-terminals up to the formation
// of the root node.
//
// lexer will use 1,2 to read the raw source file, emit tokens and construct
// terminal nodes. These terminal nodes will be fed into fbsParser.Parse defined
// in `fbs.y.go`.
//
// Production rules can be summarized by the following trees (partially):
//
//                      schema (*SchemaNode) (this is the root of the AST)
//                      /                 \
//  includes ([]*IncludeNode)             decls ([]DeclElement)
//       /   ..   \                           /    ..    \
// include (*IncludeNode)                  decl (DeclElement)
//                                                 |
//         namespaceDecl or tableDecl or structDecl or enumDecl or unionDecl or
//           rootDecl or fileExtDecl or fileIdentDecl or attrDecl or rpcDecl
//
//
//                    tableDecl (*TableDeclNode)
//                  /   |     |       |    |    \
//            Table Ident metadata '{' fields '}'
//                                       / .. \
//                                 field (*FieldNode)
//
//
//                     enumDecl (*EnumDeclNode)
//           /     |    |     |        |      |     |     \
//        Enum Ident ':' typeName metadata '{' enumVals '}'
//
//
//                      rpcDecl (*RPCDeclNode)
//                  /        |    |       |     \
//           RPCService Ident '{' rpcMethods '}'
//                                    /  ..  \
//                           rpcMethod (*RPCMethodNode)
//                      / |     |    |   |     |      |     \
//                Ident '(' Ident ')' ':' Ident metadata ';'
//
//
//          metadata (*MetadataNode)  (metadata could be nil)
//            /                   |                    \
//          '(' metadataEntries ([]*MetadataEntryNode) ')'
//                            /  ..  \
//               metadataEntry (*MetadataEntryNode)
//                    /          |         \
//                 Ident       ':'        singleVal (ValueNode)
//
//
//                           singleVal (ValueNode)
//                                     |
//                scalar (ValueNode)  or  StrLit (*StringLiteralNode)
//                           |
//         boolLit    or    intLit    or  floatLit
//   (*BoolLiteralNode)  (IntValueNode)  (FloatValueNode)
//          |                  |                    \
//  True or False       (*UintLiteralNode)        (*FloatLiteralNode)
//  (*IdentNode)    (*NegativeIntLiteralNode)    (*SpecialFloatLiteralNode)
//                  (*PositiveUintLiteralNode)   (*SignedFloatLiteralNode)
//
// Note: types without '*' are all `interface`s.
package ast
