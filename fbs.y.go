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

// Code generated by goyacc -o fbs.y.go -p fbs fbs.y. DO NOT EDIT.

//line fbs.y:33

package fbs

import __yyfmt__ "fmt"

//line fbs.y:35

import "trpc.group/trpc-go/fbs/internal/ast"

//line fbs.y:41
type fbsSymType struct {
	yys    int
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

	v  ast.ValueNode
	iv ast.IntValueNode
	fv ast.FloatValueNode

	s   *ast.StringLiteralNode
	b   *ast.BoolLiteralNode
	i   *ast.UintLiteralNode
	f   *ast.FloatLiteralNode
	id  *ast.IdentNode
	r   *ast.RuneNode
	err error
}

const StrLit = 57346
const IntLit = 57347
const FloatLit = 57348
const Ident = 57349
const True = 57350
const False = 57351
const Attribute = 57352
const Bool = 57353
const Byte = 57354
const Double = 57355
const Enum = 57356
const FileExtension = 57357
const FileIdentifier = 57358
const Float = 57359
const Float32 = 57360
const Float64 = 57361
const Include = 57362
const Inf = 57363
const Int = 57364
const Int16 = 57365
const Int32 = 57366
const Int64 = 57367
const Int8 = 57368
const Long = 57369
const Namespace = 57370
const Nan = 57371
const RootType = 57372
const RPCService = 57373
const Short = 57374
const String = 57375
const Struct = 57376
const Table = 57377
const Ubyte = 57378
const Ushort = 57379
const Uint = 57380
const Uint16 = 57381
const Uint32 = 57382
const Uint64 = 57383
const Uint8 = 57384
const Ulong = 57385
const Union = 57386
const Error = 57387

var fbsToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"StrLit",
	"IntLit",
	"FloatLit",
	"Ident",
	"True",
	"False",
	"Attribute",
	"Bool",
	"Byte",
	"Double",
	"Enum",
	"FileExtension",
	"FileIdentifier",
	"Float",
	"Float32",
	"Float64",
	"Include",
	"Inf",
	"Int",
	"Int16",
	"Int32",
	"Int64",
	"Int8",
	"Long",
	"Namespace",
	"Nan",
	"RootType",
	"RPCService",
	"Short",
	"String",
	"Struct",
	"Table",
	"Ubyte",
	"Ushort",
	"Uint",
	"Uint16",
	"Uint32",
	"Uint64",
	"Uint8",
	"Ulong",
	"Union",
	"Error",
	"'='",
	"';'",
	"':'",
	"'{'",
	"'}'",
	"'\\\\'",
	"'/'",
	"'?'",
	"'.'",
	"','",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'('",
	"')'",
	"'['",
	"']'",
	"'*'",
	"'&'",
	"'^'",
	"'%'",
	"'$'",
	"'#'",
	"'@'",
	"'!'",
	"'~'",
	"'`'",
}

var fbsStatenames = [...]string{}

const fbsEofCode = 1
const fbsErrCode = 2
const fbsInitialStackSize = 16

//line fbs.y:592

//line yacctab:1
var fbsExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 8,
	-2, 0,
	-1, 5,
	1, 1,
	-2, 0,
}

const fbsPrivate = 57344

const fbsLast = 251

var fbsAct = [...]int{
	46, 65, 143, 114, 104, 117, 102, 94, 59, 91,
	129, 97, 63, 148, 121, 47, 150, 96, 93, 108,
	152, 32, 131, 45, 151, 121, 124, 157, 119, 120,
	47, 132, 130, 128, 89, 61, 48, 45, 50, 95,
	33, 125, 95, 92, 67, 68, 77, 56, 57, 126,
	74, 86, 87, 55, 62, 72, 80, 82, 84, 78,
	75, 154, 111, 100, 98, 70, 88, 160, 161, 69,
	71, 73, 81, 83, 85, 79, 76, 101, 122, 123,
	99, 168, 127, 4, 167, 109, 106, 115, 121, 124,
	49, 119, 120, 166, 155, 64, 54, 53, 52, 51,
	107, 110, 44, 43, 125, 31, 112, 110, 153, 33,
	133, 145, 126, 60, 105, 95, 134, 92, 67, 68,
	77, 42, 38, 37, 74, 86, 87, 36, 35, 72,
	80, 82, 84, 78, 75, 149, 34, 147, 146, 70,
	88, 122, 123, 69, 71, 73, 81, 83, 85, 79,
	76, 139, 135, 41, 156, 158, 162, 163, 164, 159,
	33, 7, 40, 165, 67, 68, 77, 30, 39, 64,
	74, 86, 87, 29, 103, 72, 80, 82, 84, 78,
	75, 3, 139, 140, 6, 70, 88, 144, 90, 69,
	71, 73, 81, 83, 85, 79, 76, 18, 141, 66,
	118, 116, 113, 135, 136, 27, 142, 58, 18, 22,
	25, 26, 17, 16, 15, 4, 27, 14, 13, 137,
	22, 25, 26, 19, 12, 24, 28, 138, 11, 21,
	20, 10, 9, 8, 19, 5, 24, 28, 2, 23,
	21, 20, 1, 0, 0, 0, 0, 0, 0, 0,
	23,
}

var fbsPact = [...]int{
	63, -1000, 195, -1000, 169, 206, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 58, 102,
	129, 121, 120, 116, 115, 164, 158, 149, 114, 56,
	-1000, -1000, 55, -31, -45, -45, 42, -45, 52, 51,
	50, 49, 4, -1000, -1000, 102, -1, 106, -14, 33,
	-15, -1000, -1000, -1000, -1000, 110, -1000, 108, -44, -1000,
	16, 108, -45, -1000, 153, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 107,
	36, -1000, -41, 35, -1000, 14, -1000, 106, 83, 32,
	-16, -53, -18, -33, -1000, -17, -1000, -1000, 102, -1000,
	-1000, 33, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 198, 177, -1000, -1000, -1000, -1000, 104, -1000,
	-1000, 107, 33, -48, -30, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -26, -35, 62, -1000, -1000, 13, 47,
	20, -1000, 104, 9, 102, -1000, -45, -45, -1000, -1000,
	147, 146, -45, 46, 37, 34, -1000, -1000, -1000,
}

var fbsPgo = [...]int{
	0, 242, 181, 238, 161, 235, 233, 232, 231, 228,
	224, 218, 217, 214, 213, 212, 1, 4, 0, 7,
	18, 8, 207, 202, 3, 201, 5, 200, 12, 199,
	9, 188, 187, 2, 174, 6,
}

var fbsR1 = [...]int{
	0, 1, 3, 3, 3, 2, 5, 5, 5, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 6, 16, 16, 14, 7, 8, 9, 10, 11,
	12, 13, 15, 31, 31, 30, 33, 33, 33, 32,
	32, 35, 35, 35, 34, 34, 20, 20, 20, 19,
	19, 19, 18, 18, 22, 22, 22, 21, 21, 23,
	23, 24, 24, 24, 25, 25, 26, 26, 26, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 17, 17,
	28, 28, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29,
}

var fbsR2 = [...]int{
	0, 2, 2, 1, 0, 3, 2, 1, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 3, 1, 3, 3, 6, 6, 8, 6, 3,
	3, 3, 5, 1, 2, 8, 1, 3, 0, 1,
	3, 1, 3, 0, 1, 3, 1, 2, 0, 5,
	7, 7, 3, 0, 1, 3, 0, 1, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 2, 1,
	2, 2, 1, 2, 2, 1, 2, 2, 1, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1,
}

var fbsChk = [...]int{
	-1000, -1, -3, -2, 20, -5, -2, -4, -6, -7,
	-8, -9, -10, -11, -12, -13, -14, -15, 2, 28,
	35, 34, 14, 44, 30, 15, 16, 10, 31, 4,
	-4, 47, -16, 7, 7, 7, 7, 7, 7, 4,
	4, 4, 7, 47, 47, 54, -18, 60, -18, 48,
	-18, 47, 47, 47, 47, 49, -16, 49, -22, -21,
	7, 49, -17, -28, 62, -16, -29, 11, 12, 36,
	32, 37, 22, 38, 17, 27, 43, 13, 26, 42,
	23, 39, 24, 40, 25, 41, 18, 19, 33, 49,
	-31, -30, 7, -20, -19, 7, 61, 55, 48, -20,
	-18, -28, -35, -34, -17, 7, 50, -30, 60, 50,
	-19, 48, -21, -23, -24, 4, -25, -26, -27, 8,
	9, 5, 58, 59, 6, 21, 29, 50, 49, 63,
	50, 55, 48, -16, -17, 5, 6, 21, 29, 5,
	6, 21, 29, -33, -32, 7, -35, -17, 61, -18,
	46, 50, 55, 46, 48, 47, -24, 7, -33, -26,
	58, 59, -16, -18, -18, -18, 47, 47, 47,
}

var fbsDef = [...]int{
	4, -2, -2, 3, 0, -2, 2, 7, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 20, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	6, 19, 0, 22, 53, 53, 0, 53, 0, 0,
	0, 0, 0, 5, 21, 0, 0, 56, 0, 0,
	0, 29, 30, 31, 24, 0, 23, 48, 0, 54,
	57, 48, 53, 78, 0, 80, 81, 82, 83, 84,
	85, 86, 87, 88, 89, 90, 91, 92, 93, 94,
	95, 96, 97, 98, 99, 100, 101, 102, 103, 43,
	0, 33, 0, 0, 46, 0, 52, 0, 0, 0,
	0, 0, 0, 41, 44, 22, 32, 34, 0, 25,
	47, 0, 55, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 0, 0, 69, 72, 75, 26, 38, 79,
	28, 43, 0, 0, 53, 67, 71, 73, 76, 68,
	70, 74, 77, 0, 36, 39, 42, 45, 0, 0,
	0, 27, 38, 0, 0, 49, 53, 53, 37, 40,
	0, 0, 53, 0, 0, 0, 50, 51, 35,
}

var fbsTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 71, 3, 69, 68, 67, 65, 3,
	60, 61, 64, 58, 55, 59, 54, 52, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 48, 47,
	57, 46, 56, 53, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 62, 51, 63, 66, 3, 73, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 49, 3, 50, 72,
}

var fbsTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45,
}

var fbsTok3 = [...]int{
	0,
}

var fbsErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	fbsDebug        = 0
	fbsErrorVerbose = false
)

type fbsLexer interface {
	Lex(lval *fbsSymType) int
	Error(s string)
}

type fbsParser interface {
	Parse(fbsLexer) int
	Lookahead() int
}

type fbsParserImpl struct {
	lval  fbsSymType
	stack [fbsInitialStackSize]fbsSymType
	char  int
}

func (p *fbsParserImpl) Lookahead() int {
	return p.char
}

func fbsNewParser() fbsParser {
	return &fbsParserImpl{}
}

const fbsFlag = -1000

func fbsTokname(c int) string {
	if c >= 1 && c-1 < len(fbsToknames) {
		if fbsToknames[c-1] != "" {
			return fbsToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func fbsStatname(s int) string {
	if s >= 0 && s < len(fbsStatenames) {
		if fbsStatenames[s] != "" {
			return fbsStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func fbsErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !fbsErrorVerbose {
		return "syntax error"
	}

	for _, e := range fbsErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + fbsTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := fbsPact[state]
	for tok := TOKSTART; tok-1 < len(fbsToknames); tok++ {
		if n := base + tok; n >= 0 && n < fbsLast && fbsChk[fbsAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if fbsDef[state] == -2 {
		i := 0
		for fbsExca[i] != -1 || fbsExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; fbsExca[i] >= 0; i += 2 {
			tok := fbsExca[i]
			if tok < TOKSTART || fbsExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if fbsExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += fbsTokname(tok)
	}
	return res
}

func fbslex1(lex fbsLexer, lval *fbsSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = fbsTok1[0]
		goto out
	}
	if char < len(fbsTok1) {
		token = fbsTok1[char]
		goto out
	}
	if char >= fbsPrivate {
		if char < fbsPrivate+len(fbsTok2) {
			token = fbsTok2[char-fbsPrivate]
			goto out
		}
	}
	for i := 0; i < len(fbsTok3); i += 2 {
		token = fbsTok3[i+0]
		if token == char {
			token = fbsTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = fbsTok2[1] /* unknown char */
	}
	if fbsDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", fbsTokname(token), uint(char))
	}
	return char, token
}

func fbsParse(fbslex fbsLexer) int {
	return fbsNewParser().Parse(fbslex)
}

func (fbsrcvr *fbsParserImpl) Parse(fbslex fbsLexer) int {
	var fbsn int
	var fbsVAL fbsSymType
	var fbsDollar []fbsSymType
	_ = fbsDollar // silence set and not used
	fbsS := fbsrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	fbsstate := 0
	fbsrcvr.char = -1
	fbstoken := -1 // fbsrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		fbsstate = -1
		fbsrcvr.char = -1
		fbstoken = -1
	}()
	fbsp := -1
	goto fbsstack

ret0:
	return 0

ret1:
	return 1

fbsstack:
	/* put a state and value onto the stack */
	if fbsDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", fbsTokname(fbstoken), fbsStatname(fbsstate))
	}

	fbsp++
	if fbsp >= len(fbsS) {
		nyys := make([]fbsSymType, len(fbsS)*2)
		copy(nyys, fbsS)
		fbsS = nyys
	}
	fbsS[fbsp] = fbsVAL
	fbsS[fbsp].yys = fbsstate

fbsnewstate:
	fbsn = fbsPact[fbsstate]
	if fbsn <= fbsFlag {
		goto fbsdefault /* simple state */
	}
	if fbsrcvr.char < 0 {
		fbsrcvr.char, fbstoken = fbslex1(fbslex, &fbsrcvr.lval)
	}
	fbsn += fbstoken
	if fbsn < 0 || fbsn >= fbsLast {
		goto fbsdefault
	}
	fbsn = fbsAct[fbsn]
	if fbsChk[fbsn] == fbstoken { /* valid shift */
		fbsrcvr.char = -1
		fbstoken = -1
		fbsVAL = fbsrcvr.lval
		fbsstate = fbsn
		if Errflag > 0 {
			Errflag--
		}
		goto fbsstack
	}

fbsdefault:
	/* default state action */
	fbsn = fbsDef[fbsstate]
	if fbsn == -2 {
		if fbsrcvr.char < 0 {
			fbsrcvr.char, fbstoken = fbslex1(fbslex, &fbsrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if fbsExca[xi+0] == -1 && fbsExca[xi+1] == fbsstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			fbsn = fbsExca[xi+0]
			if fbsn < 0 || fbsn == fbstoken {
				break
			}
		}
		fbsn = fbsExca[xi+1]
		if fbsn < 0 {
			goto ret0
		}
	}
	if fbsn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			fbslex.Error(fbsErrorMessage(fbsstate, fbstoken))
			Nerrs++
			if fbsDebug >= 1 {
				__yyfmt__.Printf("%s", fbsStatname(fbsstate))
				__yyfmt__.Printf(" saw %s\n", fbsTokname(fbstoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for fbsp >= 0 {
				fbsn = fbsPact[fbsS[fbsp].yys] + fbsErrCode
				if fbsn >= 0 && fbsn < fbsLast {
					fbsstate = fbsAct[fbsn] /* simulate a shift of "error" */
					if fbsChk[fbsstate] == fbsErrCode {
						goto fbsstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if fbsDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", fbsS[fbsp].yys)
				}
				fbsp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if fbsDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", fbsTokname(fbstoken))
			}
			if fbstoken == fbsEofCode {
				goto ret1
			}
			fbsrcvr.char = -1
			fbstoken = -1
			goto fbsnewstate /* try again in the same state */
		}
	}

	/* reduction by production fbsn */
	if fbsDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", fbsn, fbsStatname(fbsstate))
	}

	fbsnt := fbsn
	fbspt := fbsp
	_ = fbspt // guard against "declared and not used"

	fbsp -= fbsR2[fbsn]
	// fbsp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if fbsp+1 >= len(fbsS) {
		nyys := make([]fbsSymType, len(fbsS)*2)
		copy(nyys, fbsS)
		fbsS = nyys
	}
	fbsVAL = fbsS[fbsp+1]

	/* consult goto table to find next state */
	fbsn = fbsR1[fbsn]
	fbsg := fbsPgo[fbsn]
	fbsj := fbsg + fbsS[fbsp].yys + 1

	if fbsj >= fbsLast {
		fbsstate = fbsAct[fbsg]
	} else {
		fbsstate = fbsAct[fbsj]
		if fbsChk[fbsstate] != -fbsn {
			fbsstate = fbsAct[fbsg]
		}
	}
	// dummy call; replaced with literal code
	switch fbsnt {

	case 1:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:154
		{
			fbsVAL.schema = ast.NewSchemaNode(fbsDollar[1].includes, fbsDollar[2].decls)
			fbslex.(*fbsLex).res = fbsVAL.schema // store result into lexer
		}
	case 2:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:160
		{
			if fbsDollar[2].include != nil {
				fbsVAL.includes = append(fbsDollar[1].includes, fbsDollar[2].include)
			} else {
				fbsVAL.includes = fbsDollar[1].includes
			}
		}
	case 3:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:167
		{
			if fbsDollar[1].include != nil {
				fbsVAL.includes = []*ast.IncludeNode{fbsDollar[1].include}
			} else {
				fbsVAL.includes = nil
			}
		}
	case 4:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:174
		{
			fbsVAL.includes = nil
		}
	case 5:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:179
		{
			fbsVAL.include = ast.NewIncludeNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].s, fbsDollar[3].r)
		}
	case 6:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:184
		{
			if fbsDollar[2].decl != nil {
				fbsVAL.decls = append(fbsDollar[1].decls, fbsDollar[2].decl)
			} else {
				fbsVAL.decls = fbsDollar[1].decls
			}
		}
	case 7:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:191
		{
			if fbsDollar[1].decl != nil {
				fbsVAL.decls = []ast.DeclElement{fbsDollar[1].decl}
			} else {
				fbsVAL.decls = nil
			}
		}
	case 8:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:198
		{
			fbsVAL.decls = nil
		}
	case 9:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:203
		{
			fbsVAL.decl = fbsDollar[1].namespaceDecl
		}
	case 10:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:206
		{
			fbsVAL.decl = fbsDollar[1].tableDecl
		}
	case 11:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:209
		{
			fbsVAL.decl = fbsDollar[1].structDecl
		}
	case 12:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:212
		{
			fbsVAL.decl = fbsDollar[1].enumDecl
		}
	case 13:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:215
		{
			fbsVAL.decl = fbsDollar[1].unionDecl
		}
	case 14:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:218
		{
			fbsVAL.decl = fbsDollar[1].rootDecl
		}
	case 15:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:221
		{
			fbsVAL.decl = fbsDollar[1].fileExtDecl
		}
	case 16:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:224
		{
			fbsVAL.decl = fbsDollar[1].fileIdentDecl
		}
	case 17:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:227
		{
			fbsVAL.decl = fbsDollar[1].attrDecl
		}
	case 18:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:230
		{
			fbsVAL.decl = fbsDollar[1].rpcDecl
		}
	case 19:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:233
		{
			fbsVAL.decl = nil
		}
	case 20:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:236
		{
			fbsVAL.decl = nil
		}
	case 21:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:241
		{
			fbsVAL.namespaceDecl = ast.NewNamespaceDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].idents.ToIdentValueNode(nil), fbsDollar[3].r)
		}
	case 22:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:246
		{
			fbsVAL.idents = &ast.IdentList{fbsDollar[1].id, nil, nil}
		}
	case 23:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:249
		{
			fbsVAL.idents = &ast.IdentList{fbsDollar[1].id, fbsDollar[2].r, fbsDollar[3].idents}
		}
	case 24:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:255
		{
			fbsVAL.attrDecl = ast.NewAttrDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].s, fbsDollar[3].r)
		}
	case 25:
		fbsDollar = fbsS[fbspt-6 : fbspt+1]
//line fbs.y:260
		{
			var opts []ast.TableDeclOption
			opts = append(opts, ast.WithTableKeyword(fbsDollar[1].id.ToKeyword()))
			opts = append(opts, ast.WithTableName(fbsDollar[2].id))
			opts = append(opts, ast.WithTableMetadata(fbsDollar[3].metadata))
			opts = append(opts, ast.WithTableOpenBrace(fbsDollar[4].r))
			opts = append(opts, ast.WithTableFields(fbsDollar[5].fields))
			opts = append(opts, ast.WithTableCloseBrace(fbsDollar[6].r))
			fbsVAL.tableDecl = ast.NewTableDeclNode(opts...)
		}
	case 26:
		fbsDollar = fbsS[fbspt-6 : fbspt+1]
//line fbs.y:272
		{
			var opts []ast.StructDeclOption
			opts = append(opts, ast.WithStructKeyword(fbsDollar[1].id.ToKeyword()))
			opts = append(opts, ast.WithStructName(fbsDollar[2].id))
			opts = append(opts, ast.WithStructMetadata(fbsDollar[3].metadata))
			opts = append(opts, ast.WithStructOpenBrace(fbsDollar[4].r))
			opts = append(opts, ast.WithStructFields(fbsDollar[5].fields))
			opts = append(opts, ast.WithStructCloseBrace(fbsDollar[6].r))
			fbsVAL.structDecl = ast.NewStructDeclNode(opts...)
		}
	case 27:
		fbsDollar = fbsS[fbspt-8 : fbspt+1]
//line fbs.y:284
		{
			var opts []ast.EnumDeclOption
			opts = append(opts, ast.WithEnumKeyword(fbsDollar[1].id.ToKeyword()))
			opts = append(opts, ast.WithEnumName(fbsDollar[2].id))
			opts = append(opts, ast.WithEnumColon(fbsDollar[3].r))
			opts = append(opts, ast.WithEnumTypeName(fbsDollar[4].typeName))
			opts = append(opts, ast.WithEnumMetadata(fbsDollar[5].metadata))
			opts = append(opts, ast.WithEnumOpenBrace(fbsDollar[6].r))
			opts = append(opts, ast.WithEnumDecls(fbsDollar[7].enumVals))
			opts = append(opts, ast.WithEnumCloseBrace(fbsDollar[8].r))
			fbsVAL.enumDecl = ast.NewEnumDeclNode(opts...)
		}
	case 28:
		fbsDollar = fbsS[fbspt-6 : fbspt+1]
//line fbs.y:298
		{
			var opts []ast.UnionDeclOption
			opts = append(opts, ast.WithUnionKeyword(fbsDollar[1].id.ToKeyword()))
			opts = append(opts, ast.WithUnionName(fbsDollar[2].id))
			opts = append(opts, ast.WithUnionMetadata(fbsDollar[3].metadata))
			opts = append(opts, ast.WithUnionOpenBrace(fbsDollar[4].r))
			opts = append(opts, ast.WithUnionDecls(fbsDollar[5].unionVals))
			opts = append(opts, ast.WithUnionCloseBrace(fbsDollar[6].r))
			fbsVAL.unionDecl = ast.NewUnionDeclNode(opts...)
		}
	case 29:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:310
		{
			fbsVAL.rootDecl = ast.NewRootDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].id, fbsDollar[3].r)
		}
	case 30:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:315
		{
			fbsVAL.fileExtDecl = ast.NewFileExtDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].s, fbsDollar[3].r)
		}
	case 31:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:320
		{
			fbsVAL.fileIdentDecl = ast.NewFileIdentDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].s, fbsDollar[3].r)
		}
	case 32:
		fbsDollar = fbsS[fbspt-5 : fbspt+1]
//line fbs.y:325
		{
			fbsVAL.rpcDecl = ast.NewRPCDeclNode(fbsDollar[1].id.ToKeyword(), fbsDollar[2].id, fbsDollar[3].r, fbsDollar[4].rpcMethods, fbsDollar[5].r)
		}
	case 33:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:330
		{
			fbsVAL.rpcMethods = []*ast.RPCMethodNode{fbsDollar[1].rpcMethod}
		}
	case 34:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:333
		{
			fbsVAL.rpcMethods = append(fbsDollar[1].rpcMethods, fbsDollar[2].rpcMethod)
		}
	case 35:
		fbsDollar = fbsS[fbspt-8 : fbspt+1]
//line fbs.y:338
		{
			var opts []ast.MethodOption
			opts = append(opts, ast.WithMethodName(fbsDollar[1].id))
			opts = append(opts, ast.WithMethodOpenParen(fbsDollar[2].r))
			opts = append(opts, ast.WithMethodReqName(fbsDollar[3].idents.ToIdentValueNode(nil)))
			opts = append(opts, ast.WithMethodCloseParen(fbsDollar[4].r))
			opts = append(opts, ast.WithMethodColon(fbsDollar[5].r))
			opts = append(opts, ast.WithMethodRspName(fbsDollar[6].idents.ToIdentValueNode(nil)))
			opts = append(opts, ast.WithMethodMetadata(fbsDollar[7].metadata))
			opts = append(opts, ast.WithMethodSemicolon(fbsDollar[8].r))
			fbsVAL.rpcMethod = ast.NewRPCMethodNode(opts...)
		}
	case 36:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:352
		{
			fbsVAL.enumVals = []*ast.EnumValueNode{fbsDollar[1].enumVal}
		}
	case 37:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:355
		{ // so the last item can have trailing ','
			if fbsDollar[3].enumVals != nil {
				fbsVAL.enumVals = append(fbsDollar[3].enumVals, fbsDollar[1].enumVal)
			} else {
				fbsVAL.enumVals = []*ast.EnumValueNode{fbsDollar[1].enumVal}
			}
		}
	case 38:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:362
		{
			fbsVAL.enumVals = nil
		}
	case 39:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:367
		{
			fbsVAL.enumVal = ast.NewEnumValueNode(fbsDollar[1].id, nil, nil)
		}
	case 40:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:370
		{
			fbsVAL.enumVal = ast.NewEnumValueNode(fbsDollar[1].id, fbsDollar[2].r, fbsDollar[3].iv)
		}
	case 41:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:375
		{
			fbsVAL.unionVals = []*ast.UnionValueNode{fbsDollar[1].unionVal}
		}
	case 42:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:378
		{ // so the last item can have trailing ','
			if fbsDollar[3].unionVals != nil {
				fbsVAL.unionVals = append(fbsDollar[3].unionVals, fbsDollar[1].unionVal)
			} else {
				fbsVAL.unionVals = []*ast.UnionValueNode{fbsDollar[1].unionVal}
			}
		}
	case 43:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:385
		{
			fbsVAL.unionVals = nil
		}
	case 44:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:390
		{
			fbsVAL.unionVal = ast.NewUnionValueNode(nil, nil, fbsDollar[1].typeName)
		}
	case 45:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:393
		{
			fbsVAL.unionVal = ast.NewUnionValueNode(fbsDollar[1].id, fbsDollar[2].r, fbsDollar[3].typeName)
		}
	case 46:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:398
		{
			fbsVAL.fields = []*ast.FieldNode{fbsDollar[1].field}
		}
	case 47:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:401
		{
			fbsVAL.fields = append(fbsDollar[1].fields, fbsDollar[2].field)
		}
	case 48:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:404
		{
			fbsVAL.fields = nil // allow empty (different from original grammar, but is needed in monster.fbs)
		}
	case 49:
		fbsDollar = fbsS[fbspt-5 : fbspt+1]
//line fbs.y:409
		{
			var opts []ast.FieldOption
			opts = append(opts, ast.WithFieldName(fbsDollar[1].id))
			opts = append(opts, ast.WithFieldColon(fbsDollar[2].r))
			opts = append(opts, ast.WithFieldTypeName(fbsDollar[3].typeName))
			opts = append(opts, ast.WithFieldMetadata(fbsDollar[4].metadata))
			opts = append(opts, ast.WithFieldSemicolon(fbsDollar[5].r))
			fbsVAL.field = ast.NewFieldNode(opts...)
		}
	case 50:
		fbsDollar = fbsS[fbspt-7 : fbspt+1]
//line fbs.y:418
		{
			var opts []ast.FieldOption
			opts = append(opts, ast.WithFieldName(fbsDollar[1].id))
			opts = append(opts, ast.WithFieldColon(fbsDollar[2].r))
			opts = append(opts, ast.WithFieldTypeName(fbsDollar[3].typeName))
			opts = append(opts, ast.WithFieldEqual(fbsDollar[4].r))
			opts = append(opts, ast.WithFieldScalar(fbsDollar[5].v))
			opts = append(opts, ast.WithFieldMetadata(fbsDollar[6].metadata))
			opts = append(opts, ast.WithFieldSemicolon(fbsDollar[7].r))
			fbsVAL.field = ast.NewFieldNode(opts...)
		}
	case 51:
		fbsDollar = fbsS[fbspt-7 : fbspt+1]
//line fbs.y:429
		{ // case: "color: Color = Green;"
			var opts []ast.FieldOption
			opts = append(opts, ast.WithFieldName(fbsDollar[1].id))
			opts = append(opts, ast.WithFieldColon(fbsDollar[2].r))
			opts = append(opts, ast.WithFieldTypeName(fbsDollar[3].typeName))
			opts = append(opts, ast.WithFieldEqual(fbsDollar[4].r))
			opts = append(opts, ast.WithFieldScalar(fbsDollar[5].id))
			opts = append(opts, ast.WithFieldMetadata(fbsDollar[6].metadata))
			opts = append(opts, ast.WithFieldSemicolon(fbsDollar[7].r))
			fbsVAL.field = ast.NewFieldNode(opts...)
		}
	case 52:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:442
		{
			fbsVAL.metadata = ast.NewMetadataNode(fbsDollar[1].r, fbsDollar[2].metadataEntries, fbsDollar[3].r)
		}
	case 53:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:445
		{
			fbsVAL.metadata = nil
		}
	case 54:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:450
		{
			fbsVAL.metadataEntries = []*ast.MetadataEntryNode{fbsDollar[1].metadataEntry}
		}
	case 55:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:453
		{
			fbsVAL.metadataEntries = append(fbsDollar[1].metadataEntries, fbsDollar[3].metadataEntry)
		}
	case 56:
		fbsDollar = fbsS[fbspt-0 : fbspt+1]
//line fbs.y:456
		{
			fbsVAL.metadataEntries = nil
		}
	case 57:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:461
		{
			fbsVAL.metadataEntry = ast.NewMetadataEntryNode(fbsDollar[1].id, nil, nil)
		}
	case 58:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:464
		{
			fbsVAL.metadataEntry = ast.NewMetadataEntryNode(fbsDollar[1].id, fbsDollar[2].r, fbsDollar[3].v)
		}
	case 59:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:469
		{ // ast.ValueNode
			fbsVAL.v = fbsDollar[1].v
		}
	case 60:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:472
		{ // *ast.StringLiteralNode
			fbsVAL.v = fbsDollar[1].s
		}
	case 61:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:490
		{ // *ast.BoolLiteralNode
			fbsVAL.v = fbsDollar[1].b
		}
	case 62:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:493
		{ // ast.IntValueNode (an interface)
			fbsVAL.v = fbsDollar[1].iv
		}
	case 63:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:496
		{ // ast.FloatValueNode (an interface)
			fbsVAL.v = fbsDollar[1].fv
		}
	case 64:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:501
		{
			fbsVAL.b = ast.NewBoolLiteralNode(fbsDollar[1].id.ToKeyword())
		}
	case 65:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:504
		{
			fbsVAL.b = ast.NewBoolLiteralNode(fbsDollar[1].id.ToKeyword())
		}
	case 66:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:509
		{ // *ast.UintLiteralNode
			fbsVAL.iv = fbsDollar[1].i
		}
	case 67:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:512
		{ // *ast.PositiveUintLiteralNode
			fbsVAL.iv = ast.NewPositiveUintLiteralNode(fbsDollar[1].r, fbsDollar[2].i)
		}
	case 68:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:515
		{ // *ast.NegativeIntLiteralNode
			fbsVAL.iv = ast.NewNegativeIntLiteralNode(fbsDollar[1].r, fbsDollar[2].i)
		}
	case 69:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:520
		{ // *ast.FloatLiteralNode
			fbsVAL.fv = fbsDollar[1].f
		}
	case 70:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:523
		{ // *ast.SignedFloatLiteralNode
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, fbsDollar[2].f)
		}
	case 71:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:526
		{ // *ast.SignedFloatLiteralNode
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, fbsDollar[2].f)
		}
	case 72:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:529
		{ // *ast.SpecialFloatLiteralNode
			fbsVAL.fv = ast.NewSpecialFloatLiteralNode(fbsDollar[1].id.ToKeyword())
		}
	case 73:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:532
		{ // *ast.SignedFloatLiteralNode
			f := ast.NewSpecialFloatLiteralNode(fbsDollar[2].id.ToKeyword())
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, f)
		}
	case 74:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:536
		{ // *ast.SignedFloatLiteralNode
			f := ast.NewSpecialFloatLiteralNode(fbsDollar[2].id.ToKeyword())
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, f)
		}
	case 75:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:540
		{ // *ast.SpecialFloatLiteralNode
			fbsVAL.fv = ast.NewSpecialFloatLiteralNode(fbsDollar[1].id.ToKeyword())
		}
	case 76:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:543
		{ // *ast.SignedFloatLiteralNode
			f := ast.NewSpecialFloatLiteralNode(fbsDollar[2].id.ToKeyword())
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, f)
		}
	case 77:
		fbsDollar = fbsS[fbspt-2 : fbspt+1]
//line fbs.y:547
		{ // *ast.SignedFloatLiteralNode
			f := ast.NewSpecialFloatLiteralNode(fbsDollar[2].id.ToKeyword())
			fbsVAL.fv = ast.NewSignedFloatLiteralNode(fbsDollar[1].r, f)
		}
	case 78:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:553
		{
			fbsVAL.typeName = ast.NewTypeNameNode(nil, fbsDollar[1].identLit, nil)
		}
	case 79:
		fbsDollar = fbsS[fbspt-3 : fbspt+1]
//line fbs.y:556
		{ // [typeName] means vector of types
			fbsVAL.typeName = ast.NewTypeNameNode(fbsDollar[1].r, fbsDollar[2].identLit, fbsDollar[3].r)
		}
	case 80:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:561
		{ // *ast.IdentList => ast.IdentLiteralElement
			fbsVAL.identLit = fbsDollar[1].idents.ToIdentValueNode(nil)
		}
	case 81:
		fbsDollar = fbsS[fbspt-1 : fbspt+1]
//line fbs.y:564
		{
			fbsVAL.identLit = ast.IdentLiteralElement(fbsDollar[1].id)
		}
	}
	goto fbsstack /* stack new state and value */
}
