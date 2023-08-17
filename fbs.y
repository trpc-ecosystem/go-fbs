// This file is written according to 
// https://google.github.io/flatbuffers/flatbuffers_grammar.html
// 
// However, the grammar listed in the link is somewhat outdated and
// contains mistakes and ambiguities. 
// 
// We make some rearrangements, for example,
// 
// `object` related grammars are eliminated since no actual usage can be found:
// 
//     object = { commasep( ident : value ) }
//     value = single_value | object | [ commasep( value ) ]
// 
// 'values' and 'value' are also eliminated. 
//
// 'table' and 'struct' declarations are separated.
// 
// 'enum' and 'union' declarations are separated, also are 'enumVals' and 'unionVals'.
// 
// The second node of attribute declaration should be string literal rather than identifier.
// 
// The fileds of table counld be empty.
//
// 'typeName' can have dots inside.
// 
// For .y (yacc) file format, refer to http://dinosaur.compilertools.net/yacc/index.html
// and goyacc https://pkg.go.dev/golang.org/x/tools/cmd/goyacc 
//
// You may want to take a look at this .fbs example file: 
// https://github.com/google/flatbuffers/blob/master/tests/monster_test.fbs
// 
// `%{ }%` encloses declarations and definitions.
%{
package fbs

import "trpc.group/trpc-go/fbs/internal/ast"

%}

// %union define types used in yacc and map them to Go types.
%union{
	schema *ast.SchemaNode

	include  *ast.IncludeNode
	includes []*ast.IncludeNode
	decl     ast.DeclElement
	decls    []ast.DeclElement

	namespaceDecl *ast.NamespaceDeclNode
	tableDecl     *ast.TableDeclNode 
	structDecl    *ast.StructDeclNode 
	enumDecl      *ast.EnumDeclNode
	unionDecl     *ast.UnionDeclNode
	rootDecl      *ast.RootDeclNode
	fileExtDecl   *ast.FileExtDeclNode
	fileIdentDecl *ast.FileIdentDeclNode
	attrDecl      *ast.AttrDeclNode
	rpcDecl       *ast.RPCDeclNode

	idents          *ast.IdentList
	typeName        *ast.TypeNameNode 
	identLit        ast.IdentLiteralElement
	metadata        *ast.MetadataNode
	field           *ast.FieldNode
	fields          []*ast.FieldNode
	metadataEntry   *ast.MetadataEntryNode
	metadataEntries []*ast.MetadataEntryNode
	rpcMethod       *ast.RPCMethodNode
	rpcMethods      []*ast.RPCMethodNode
	enumVal         *ast.EnumValueNode
	enumVals        []*ast.EnumValueNode
	unionVal        *ast.UnionValueNode
	unionVals       []*ast.UnionValueNode

	v   ast.ValueNode 
	iv  ast.IntValueNode
	fv  ast.FloatValueNode

	s   *ast.StringLiteralNode
	b   *ast.BoolLiteralNode
	i   *ast.UintLiteralNode
	f   *ast.FloatLiteralNode
	id  *ast.IdentNode
	r   *ast.RuneNode
	err error
}

// non-terminals. %type is used to associate union member names with non-terminals.
%type <schema> schema
//     ^^^^^^  ^^^^^^
//       |        non-terminal
//  union member names (declared in %union)


%type <include> include 
%type <includes> includes
%type <decl> decl 
%type <decls> decls 

%type <namespaceDecl> namespaceDecl
%type <tableDecl> tableDecl
%type <structDecl> structDecl
%type <enumDecl> enumDecl 
%type <unionDecl> unionDecl 
%type <rootDecl> rootDecl 
%type <fileExtDecl> fileExtDecl
%type <fileIdentDecl> fileIdentDecl 
%type <attrDecl> attrDecl 
%type <rpcDecl> rpcDecl

%type <idents> idents 
%type <typeName> typeName
%type <metadata> metadata 
%type <field> field 
%type <fields> fields 
%type <metadataEntry> metadataEntry 
%type <metadataEntries> metadataEntries 
%type <v> singleVal scalar
%type <b> boolLit
%type <iv> intLit
%type <fv> floatLit
%type <identLit> typeLit
%type <id> builtinType
%type <rpcMethod> rpcMethod
%type <rpcMethods> rpcMethods
%type <enumVal> enumVal
%type <enumVals> enumVals
%type <unionVal> unionVal
%type <unionVals> unionVals

// terminals. %token associates union member names with terminals. 
%token <s> StrLit        // string literal
%token <i> IntLit        // integer literal
%token <f> FloatLit      // floating point literal
%token <id> Ident         // identifiers
%token <id> True False   // true false
%token <id> Attribute Bool Byte Double Enum FileExtension
%token <id> FileIdentifier Float Float32 Float64 Include Inf
%token <id> Int Int16 Int32 Int64 Int8 Long Namespace Nan
%token <id> RootType RPCService Short String Struct Table
%token <id> Ubyte Ushort Uint Uint16 Uint32 Uint64 Uint8 Ulong Union
%token <err> Error
// Although some of the tokens defined below are not used directly, they can help
// improve error messages. 
%token <r>   '=' ';' ':' '{' '}' '\\' '/' '?' '.' ',' '>' '<' '+' '-' '(' ')' '[' ']' '*' '&' '^' '%' '$' '#' '@' '!' '~' '`'

%%  // A pair of `%%` `%%` encloses production rules.

// schema is the root node of AST, it is of type *ast.SchemaNode
// $$ represents left hand side of colon, in this case is `schema`
// $1 represents the first node on the right hand side of colon, `includes`
// $2 refers to the second node `decls`
// $3 $4 ..
schema: includes decls {
		$$ = ast.NewSchemaNode($1, $2)
		fbslex.(*fbsLex).res = $$  // store result into lexer
	}

// includes is of type []*ast.IncludeNode
includes: includes include {
		if $2 != nil {
			$$ = append($1, $2)
		} else {
			$$ = $1 
		}
	}
	| include {
		if $1 != nil {
			$$ = []*ast.IncludeNode{$1}
		} else {
			$$ = nil 
		}
	}
	| {
		$$ = nil 
	}

// include is of type *ast.IncludeNode
include: Include StrLit ';' {
		$$ = ast.NewIncludeNode($1.ToKeyword(), $2, $3)
	}

// decls is of type []ast.DeclElement
decls: decls decl {
		if $2 != nil {
			$$ = append($1, $2)
		} else {
			$$ = $1 
		}
	}
	| decl {
		if $1 != nil {
			$$ = []ast.DeclElement{$1}
		} else {
			$$ = nil 
		}
	}
	| {
		$$ = nil 
	}

// decl is of type ast.DeclElement (an interface)
decl: namespaceDecl {
		$$ = $1
	}
	| tableDecl {
		$$ = $1
	}
	| structDecl {
		$$ = $1
	}
	| enumDecl {
		$$ = $1
	}
	| unionDecl {
		$$ = $1 
	}
	| rootDecl {
		$$ = $1 
	}
	| fileExtDecl {
		$$ = $1 
	}
	| fileIdentDecl {
		$$ = $1 
	}
	| attrDecl {
		$$ = $1 
	}
	| rpcDecl {
		$$ = $1 
	}
	| error ';' {
		$$ = nil 
	}
	| error {
		$$ = nil 
	} 

// namespaceDecl is of type *ast.NamespaceDeclNode
namespaceDecl: Namespace idents ';' {
		$$ = ast.NewNamespaceDeclNode($1.ToKeyword(), $2.ToIdentValueNode(nil), $3)	
	}

// idents is of type *ast.IdentList
idents: Ident {
		$$ = &ast.IdentList{$1, nil, nil}
	}
	| Ident '.' idents {
		$$ = &ast.IdentList{$1, $2, $3}
	}

// attrDecl is of type *ast.AttrDeclNode
// The second node should be string literal rather than identifier, the original grammar is wrong.
attrDecl: Attribute StrLit ';' {
		$$ = ast.NewAttrDeclNode($1.ToKeyword(), $2, $3)
	}

// tableDecl is of type *ast.TableDeclNode
tableDecl: Table Ident metadata '{' fields '}' {
		var opts []ast.TableDeclOption
		opts = append(opts, ast.WithTableKeyword($1.ToKeyword()))
		opts = append(opts, ast.WithTableName($2))
		opts = append(opts, ast.WithTableMetadata($3))
		opts = append(opts, ast.WithTableOpenBrace($4))
		opts = append(opts, ast.WithTableFields($5))
		opts = append(opts, ast.WithTableCloseBrace($6))
		$$ = ast.NewTableDeclNode(opts...)
	}

// structDecl is of type *ast.StructDeclNode
structDecl: Struct Ident metadata '{' fields '}' {
		var opts []ast.StructDeclOption
		opts = append(opts, ast.WithStructKeyword($1.ToKeyword()))
		opts = append(opts, ast.WithStructName($2))
		opts = append(opts, ast.WithStructMetadata($3))
		opts = append(opts, ast.WithStructOpenBrace($4))
		opts = append(opts, ast.WithStructFields($5))
		opts = append(opts, ast.WithStructCloseBrace($6))
		$$ = ast.NewStructDeclNode(opts...)
	}

// enumDecl is of type *ast.EnumDeclNode
enumDecl: Enum Ident ':' typeName metadata '{' enumVals '}' {
		var opts []ast.EnumDeclOption
		opts = append(opts, ast.WithEnumKeyword($1.ToKeyword()))
		opts = append(opts, ast.WithEnumName($2))
		opts = append(opts, ast.WithEnumColon($3))
		opts = append(opts, ast.WithEnumTypeName($4))
		opts = append(opts, ast.WithEnumMetadata($5))
		opts = append(opts, ast.WithEnumOpenBrace($6))
		opts = append(opts, ast.WithEnumDecls($7))
		opts = append(opts, ast.WithEnumCloseBrace($8))
		$$ = ast.NewEnumDeclNode(opts...)
	}

// unionDecl is of type *ast.UnionDeclNode
unionDecl: Union Ident metadata '{' unionVals '}' {
		var opts []ast.UnionDeclOption
		opts = append(opts, ast.WithUnionKeyword($1.ToKeyword()))
		opts = append(opts, ast.WithUnionName($2))
		opts = append(opts, ast.WithUnionMetadata($3))
		opts = append(opts, ast.WithUnionOpenBrace($4))
		opts = append(opts, ast.WithUnionDecls($5))
		opts = append(opts, ast.WithUnionCloseBrace($6))
		$$ = ast.NewUnionDeclNode(opts...)
	}

// rootDecl is of type *ast.RootDeclNode
rootDecl: RootType Ident ';' {
		$$ = ast.NewRootDeclNode($1.ToKeyword(), $2, $3)
	}

// fileExtDecl is of type *ast.FileExtDeclNode
fileExtDecl: FileExtension StrLit ';' {
		$$ = ast.NewFileExtDeclNode($1.ToKeyword(), $2, $3)
	}

// fileIdentDecl is of type *ast.FileIdentDeclNode
fileIdentDecl: FileIdentifier StrLit ';' {
		$$ = ast.NewFileIdentDeclNode($1.ToKeyword(), $2, $3)
	}

// rpcDecl is of type *ast.RPCDeclNode
rpcDecl: RPCService Ident '{' rpcMethods '}' {
		$$ = ast.NewRPCDeclNode($1.ToKeyword(), $2, $3, $4, $5)
	}

// rpcMethods is of type []*ast.RPCMethodNode
rpcMethods: rpcMethod {
		$$ = []*ast.RPCMethodNode{$1}
	}
	| rpcMethods rpcMethod {
		$$ = append($1, $2)
	}

// rpcMethod is of type *ast.RPCMethodNode
rpcMethod: Ident '(' idents ')' ':' idents metadata ';' {
		var opts []ast.MethodOption 
		opts = append(opts, ast.WithMethodName($1))
		opts = append(opts, ast.WithMethodOpenParen($2))
		opts = append(opts, ast.WithMethodReqName($3.ToIdentValueNode(nil)))
		opts = append(opts, ast.WithMethodCloseParen($4))
		opts = append(opts, ast.WithMethodColon($5))
		opts = append(opts, ast.WithMethodRspName($6.ToIdentValueNode(nil)))
		opts = append(opts, ast.WithMethodMetadata($7))
		opts = append(opts, ast.WithMethodSemicolon($8))
		$$ = ast.NewRPCMethodNode(opts...)
	}

// enumVals is of type []*ast.EnumValueNode
enumVals: enumVal {
		$$ = []*ast.EnumValueNode{$1}
	}
	| enumVal ',' enumVals { // so the last item can have trailing ','
		if $3 != nil {
			$$ = append($3, $1)
		} else {
			$$ = []*ast.EnumValueNode{$1}
		}
	}
	| {
		$$ = nil 
	}

// enumVal is of type *ast.EnumValueNode
enumVal: Ident {
		$$ = ast.NewEnumValueNode($1, nil, nil)
	}
	| Ident '=' intLit {
		$$ = ast.NewEnumValueNode($1, $2, $3)
	}

// unionVals is of type []*ast.UnionValueNode
unionVals: unionVal {
		$$ = []*ast.UnionValueNode{$1}
	}
	| unionVal ',' unionVals { // so the last item can have trailing ','
		if $3 != nil {
			$$ = append($3, $1)
		} else {
			$$ = []*ast.UnionValueNode{$1}
		}
	}
	| {
		$$ = nil 
	}

// unionVal is of type *ast.UnionValueNode
unionVal: typeName {
		$$ = ast.NewUnionValueNode(nil, nil, $1)
	}
	| Ident ':' typeName {
		$$ = ast.NewUnionValueNode($1, $2, $3)
	}

// fields is of type []*ast.FieldNode
fields: field {
		$$ = []*ast.FieldNode{$1}
	}
	| fields field {
		$$ = append($1, $2)
	} 
	| {
		$$ = nil // allow empty (different from original grammar, but is needed in monster.fbs)
	}

// field is of type *ast.FieldNode
field: Ident ':' typeName metadata ';' {
		var opts []ast.FieldOption 
		opts = append(opts, ast.WithFieldName($1))
		opts = append(opts, ast.WithFieldColon($2))
		opts = append(opts, ast.WithFieldTypeName($3))
		opts = append(opts, ast.WithFieldMetadata($4))
		opts = append(opts, ast.WithFieldSemicolon($5))
		$$ = ast.NewFieldNode(opts...)
	}
	| Ident ':' typeName '=' scalar metadata ';' {
		var opts []ast.FieldOption 
		opts = append(opts, ast.WithFieldName($1))
		opts = append(opts, ast.WithFieldColon($2))
		opts = append(opts, ast.WithFieldTypeName($3))
		opts = append(opts, ast.WithFieldEqual($4))
		opts = append(opts, ast.WithFieldScalar($5))
		opts = append(opts, ast.WithFieldMetadata($6))
		opts = append(opts, ast.WithFieldSemicolon($7))
		$$ = ast.NewFieldNode(opts...)
	}
	| Ident ':' typeName '=' Ident metadata ';' { // case: "color: Color = Green;"
		var opts []ast.FieldOption 
		opts = append(opts, ast.WithFieldName($1))
		opts = append(opts, ast.WithFieldColon($2))
		opts = append(opts, ast.WithFieldTypeName($3))
		opts = append(opts, ast.WithFieldEqual($4))
		opts = append(opts, ast.WithFieldScalar($5))
		opts = append(opts, ast.WithFieldMetadata($6))
		opts = append(opts, ast.WithFieldSemicolon($7))
		$$ = ast.NewFieldNode(opts...)
	}

// metadata is of type *ast.MetadataNode
metadata: '(' metadataEntries ')' {
		$$ = ast.NewMetadataNode($1, $2, $3)	
	}
	| {
		$$ = nil 
	}

// metadataEntries is of type []*ast.MetadataEntryNode
metadataEntries: metadataEntry {
		$$ = []*ast.MetadataEntryNode{$1}
	}
	| metadataEntries ',' metadataEntry {
		$$ = append($1, $3)
	}
	| {
		$$ = nil 
	}

// metadataEntry is of type *ast.MetadataEntryNode 
metadataEntry: Ident {
		$$ = ast.NewMetadataEntryNode($1, nil, nil)
	}
	| Ident ':' singleVal {
		$$ = ast.NewMetadataEntryNode($1, $2, $3)
	}

// singleVal is of type ast.ValueNode (an interface)
singleVal: scalar { // ast.ValueNode 
		$$ = $1 
	}
	| StrLit { // *ast.StringLiteralNode
		$$ = $1 
	}

//                             singleVal(ValueNode)
//                                     |         
//                 scalar(ValueNode)  or StrLit(*StringLiteralNode)
//                           |
//         boolLit    or    intLit    or  floatLit 
//   (*BoolLiteralNode)  (IntValueNode)  (FloatValueNode)
//          |                  |                    \
//  True or False       (*UintLiteralNode)        (*FloatLiteralNode)
//  (*IdentNode)    (*NegativeIntLiteralNode)    (*SpecialFloatLiteralNode)
//                  (*PositiveUintLiteralNode)   (*SignedFloatLiteralNode)
//
// Note: types without '*' are all `interface`s.

// scalar is of type ast.ValueNode 
scalar: boolLit { // *ast.BoolLiteralNode 
		$$ = $1 
	}
	| intLit { // ast.IntValueNode (an interface)
		$$ = $1 
	}
	| floatLit { // ast.FloatValueNode (an interface)
		$$ = $1
	}

// boolLit is of type *ast.BoolLiteralNode 
boolLit: True {
		$$ = ast.NewBoolLiteralNode($1.ToKeyword())
	} 
	| False {
		$$ = ast.NewBoolLiteralNode($1.ToKeyword())
	}

// intLit is of type ast.IntValueNode
intLit: IntLit { // *ast.UintLiteralNode
		$$ = $1 
	}
	| '+' IntLit { // *ast.PositiveUintLiteralNode
		$$ = ast.NewPositiveUintLiteralNode($1, $2)
	}
	| '-' IntLit { // *ast.NegativeIntLiteralNode
		$$ = ast.NewNegativeIntLiteralNode($1, $2)
	}

// floatLit is of type ast.FloatValueNode (an interface)
floatLit: FloatLit { // *ast.FloatLiteralNode
		$$ = $1 
	}
	| '-' FloatLit { // *ast.SignedFloatLiteralNode
		$$ = ast.NewSignedFloatLiteralNode($1, $2)
	}
	| '+' FloatLit { // *ast.SignedFloatLiteralNode
		$$ = ast.NewSignedFloatLiteralNode($1, $2)
	}
	| Inf { // *ast.SpecialFloatLiteralNode
		$$ = ast.NewSpecialFloatLiteralNode($1.ToKeyword())
	}
	| '+' Inf { // *ast.SignedFloatLiteralNode
		f := ast.NewSpecialFloatLiteralNode($2.ToKeyword())
		$$ = ast.NewSignedFloatLiteralNode($1, f)
	}
	| '-' Inf { // *ast.SignedFloatLiteralNode
		f := ast.NewSpecialFloatLiteralNode($2.ToKeyword())
		$$ = ast.NewSignedFloatLiteralNode($1, f)
	} 
	| Nan { // *ast.SpecialFloatLiteralNode
		$$ = ast.NewSpecialFloatLiteralNode($1.ToKeyword())
	}
	| '+' Nan { // *ast.SignedFloatLiteralNode
		f := ast.NewSpecialFloatLiteralNode($2.ToKeyword())
		$$ = ast.NewSignedFloatLiteralNode($1, f)
	}
	| '-' Nan { // *ast.SignedFloatLiteralNode
		f := ast.NewSpecialFloatLiteralNode($2.ToKeyword())
		$$ = ast.NewSignedFloatLiteralNode($1, f)
	} 

// typeName is of type *ast.TypeNameNode
typeName: typeLit {
		$$ = ast.NewTypeNameNode(nil, $1, nil)
	}
	| '[' typeLit ']' { // [typeName] means vector of types
		$$ = ast.NewTypeNameNode($1, $2, $3)
	}

// typeLit is of type ast.IdentLiteralElement
typeLit: idents { // *ast.IdentList => ast.IdentLiteralElement
		$$ = $1.ToIdentValueNode(nil)
	}
	| builtinType {
		$$ = ast.IdentLiteralElement($1)
	}

// builtinType is of type *ast.IdentNode 
builtinType: Bool
	| Byte
	| Ubyte
	| Short
	| Ushort
	| Int
	| Uint
	| Float
	| Long
	| Ulong
	| Double
	| Int8
	| Uint8
	| Int16
	| Uint16
	| Int32
	| Uint32
	| Int64
	| Uint64
	| Float32
	| Float64
	| String

%%
