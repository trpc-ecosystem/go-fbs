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

package fbs

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"

	"trpc.group/trpc-go/fbs/internal/ast"
)

// Tabsize is assumed to be 4.
const Tabsize = 4

// fbsLex implements fbsLexer interface defined in fbs.y.go.
type fbsLex struct {
	filename string
	input    *runeReader
	handler  *errorHandler
	res      *ast.SchemaNode
	// line and column of current token.
	// offset is the number of rune in current token.
	line   int
	col    int
	offset int
	// presym stores previous token.
	preSym ast.TerminalNode
	eof    ast.TerminalNode
	// line, column and offset of previous token.
	preLine   int
	preCol    int
	preOffset int
	// comments and ws stores comments and whitespaces.
	comments []ast.Comment
	ws       []rune // white space
}

// newLexer creates a new lexer that reads in bytes and emits tokens.
func newLexer(in io.Reader, filename string, errHandler *errorHandler) *fbsLex {
	br := bufio.NewReader(in)
	return &fbsLex{
		input:    &runeReader{r: br},
		filename: filename,
		handler:  errHandler,
	}
}

// Lex implements fbsLexer interface defined in fbs.y.go.
func (f *fbsLex) Lex(lval *fbsSymType) int {
	if f.handler.getError() != nil {
		// error has already occurred, skip the rest of the input
		return 0
	}
	f.preLine = f.line
	f.preCol = f.col
	f.preOffset = f.offset
	f.comments = nil
	f.ws = nil
	f.input.endMark()
	return f.lex(lval)
}

// Error implements fbsLexer interface defined in fbs.y.go.
func (f *fbsLex) Error(s string) {
	f.wrappedError(errors.New(s))
}

// lex start the real processing of the lexer.
func (f *fbsLex) lex(lval *fbsSymType) int {
	for {
		c, n, err := f.input.readRune()
		if err != nil {
			f.setRune(lval, 0)
			f.eof = lval.r
			return 0
		}
		f.preLine = f.line
		f.preCol = f.col
		f.preOffset = f.offset
		f.offset += n
		f.adjustPos(c)
		cont, token := f.analyze(lval, c)
		if cont {
			continue
		}
		if token != -1 {
			return token
		}
		f.setRune(lval, c)
		return int(c)
	}
}

// analyze processes the single rune to dispatch the work to different method.
// 1. white spaces
// 2. start with a dot: floating point or a rune(dot)
// 3. start with alphabets or underscore: identifier
// 4. start with a number: number literal
// 5. start with quotation mark: string literal
// 6. start with back slash: line/block comments
func (f *fbsLex) analyze(lval *fbsSymType, c rune) (bool, int) {
	// Check white spaces.
	if strings.ContainsRune("\n\r\t ", c) {
		f.ws = append(f.ws, c)
		return true, -1
	}
	// Start processing the valid rune.
	f.input.startMark(c)
	if c == '.' {
		return false, f.analyzeDot(lval, c)
	}
	// Valid identifier usually starts with underscore or alphabet.
	if c == '_' || isAlpha(c) {
		return false, f.analyzeIdentifier(lval, c)
	}
	// Valid number literal usually starts with a number.
	if isNum(c) {
		return false, f.analyzeNumericLiteral(lval, c)
	}
	// Valid string literal starts with quotation mark.
	if c == '\'' || c == '"' {
		return false, f.analyzeStringLiteral(lval, c)
	}
	// Valid comment starts with a slash.
	if c == '/' {
		return f.analyzeComments(lval, c)
	}
	return false, -1
}

// analyzeDot starts from the dot rune.
// The token could be:
// 1. floating point
// 2. a rune (a single dot)
func (f *fbsLex) analyzeDot(lval *fbsSymType, c rune) int {
	// could be decimal literals
	cn, _, err := f.input.readRune()
	if err != nil {
		f.setRune(lval, c)
		return int(c)
	}
	if isNum(cn) {
		f.adjustPos(cn)
		token := []rune{c, cn}
		token = f.readNumber(token, false, true)
		num, err := strconv.ParseFloat(string(token), 64)
		if err != nil {
			f.setError(lval, err)
			return Error
		}
		f.setFloat(lval, num)
		return FloatLit
	}
	f.input.unreadRune(cn)
	f.setRune(lval, c)
	return int(c)
}

// analyzeIdentifier starts from an underscore or an alphabet.
// The token could be:
// 1. keyword
// 2. plain identifier
func (f *fbsLex) analyzeIdentifier(lval *fbsSymType, c rune) int {
	// identifier
	token := []rune{c}
	token = f.readIdentifier(token)
	s := string(token)
	if t, ok := keywords[s]; ok {
		f.setIdent(lval, s)
		return t
	}
	f.setIdent(lval, s)
	return Ident
}

// analyzeNumericLiteral starts from a number.
// The presentation could be
// 1. hexadecimal: 0x3f3f3f3f
// 2. decimal literal: 666 9.9
func (f *fbsLex) analyzeNumericLiteral(lval *fbsSymType, c rune) int {
	// integer or float literal
	if c == '0' {
		if token := f.analyzeStartsWithZero(lval); token != -1 {
			return token
		}
	}
	if token := f.analyzeDecimal(lval, c); token != -1 {
		return token
	}
	return -1
}

// analyzeStartsWithZero process number that starts with zero.
// The token could be:
// 1. 0
// 2. hexadecimal
func (f *fbsLex) analyzeStartsWithZero(lval *fbsSymType) int {
	cn, _, err := f.input.readRune()
	if err != nil {
		f.setInt(lval, 0)
		return IntLit
	}
	if isHexStart(cn) {
		return f.analyzeHexNumber(lval, cn)
	}
	f.input.unreadRune(cn)
	return -1
}

// analyzeHexNumber reads numbers that form a hexadecimal.
func (f *fbsLex) analyzeHexNumber(lval *fbsSymType, cn rune) int {
	cnn, _, err := f.input.readRune()
	if err != nil {
		f.input.unreadRune(cn)
		f.setInt(lval, 0)
		return IntLit
	}
	if isHexNum(cnn) {
		// hexadecimal
		f.adjustPos(cn, cnn)
		token := []rune{cnn}
		token = f.readHexNumber(token)
		ui, err := strconv.ParseUint(string(token), 16, 64)
		if err != nil {
			f.setError(lval, err)
			return Error
		}
		f.setInt(lval, ui)
		return IntLit
	}
	f.input.unreadRune(cnn)
	f.input.unreadRune(cn)
	f.setInt(lval, 0)
	return IntLit
}

// analyzeDecimal tries to form integer or floating point value.
func (f *fbsLex) analyzeDecimal(lval *fbsSymType, c rune) int {
	runes := []rune{c}
	//                      allowDot allowExp
	runes = f.readNumber(runes, true, true)
	if token := f.analyzeFloat(lval, runes); token != -1 {
		return token
	}
	return f.analyzeInteger(lval, runes)
}

// analyzeFloat tries to form floating points.
func (f *fbsLex) analyzeFloat(lval *fbsSymType, runes []rune) int {
	number := string(runes)
	if strings.Contains(number, ".") || strings.Contains(number, "e") || strings.Contains(number, "E") {
		// floating point
		num, err := strconv.ParseFloat(number, 64)
		if err != nil {
			f.setError(lval, err)
			return Error
		}
		f.setFloat(lval, num)
		return FloatLit
	}
	return -1
}

// analyzeInteger tries to form an integer value. It will be parsed as floating point
// if it is too large.
func (f *fbsLex) analyzeInteger(lval *fbsSymType, runes []rune) int {
	number := string(runes)
	// integer (decimal or octal)
	ui, err := strconv.ParseUint(number, 0, 64)
	if err != nil {
		if numErr, ok := err.(*strconv.NumError); ok && numErr.Err == strconv.ErrRange {
			// parse to float if too large
			num, err := strconv.ParseFloat(number, 64)
			if err == nil {
				f.setFloat(lval, num)
				return FloatLit
			}
		}
		f.setError(lval, err)
		return Error
	}
	f.setInt(lval, ui)
	return IntLit
}

// analyzeStringLiteral tries to form a valid string literal.
func (f *fbsLex) analyzeStringLiteral(lval *fbsSymType, c rune) int {
	// string literal
	s, err := f.analyzeString(c)
	if err != nil {
		f.setError(lval, err)
		return Error
	}
	f.setString(lval, s)
	return StrLit
}

// analyzeComments tries to form line/block comments.
func (f *fbsLex) analyzeComments(lval *fbsSymType, c rune) (bool, int) {
	// comment
	cn, _, err := f.input.readRune()
	if err != nil {
		f.setRune(lval, '/')
		return false, int(c)
	}
	if cont, token, ok := f.analyzeLineComment(cn); ok {
		return cont, token
	}
	if cont, token, ok := f.analyzeBlockComment(lval, cn); ok {
		return cont, token
	}
	f.input.unreadRune(cn)
	return false, -1
}

// analyzeLineComment tries to form line comment.
func (f *fbsLex) analyzeLineComment(c rune) (cont bool, token int, ok bool) {
	if c == '/' {
		// line comment
		f.adjustPos(c)
		hitNewline := f.skipToEndOfLineComment()
		comment := f.newComment()
		comment.PosRange.End.Col++
		if hitNewline {
			f.adjustPos('\n')
		}
		f.comments = append(f.comments, comment)
		return true, -1, true
	}
	return false, -1, false
}

// analyzeBlockComment tries to form block comment.
func (f *fbsLex) analyzeBlockComment(lval *fbsSymType, c rune) (cont bool, token int, ok bool) {
	if c == '*' {
		// block of comments
		f.adjustPos(c)
		if ok := f.skipToEndOfBlockComment(); !ok {
			f.setError(lval, errors.New("block comment never terminates, unexpected EOF"))
			return false, Error, true
		}
		f.comments = append(f.comments, f.newComment())
		return true, -1, true
	}
	return false, -1, false
}

// analyzeString reads until another quotation mark to form a string literal.
func (f *fbsLex) analyzeString(quote rune) (string, error) {
	buf := &bytes.Buffer{}
	for {
		c, err := f.readRune()
		if err != nil {
			return "", err
		}
		f.adjustPos(c)
		if c == quote { // end quote encountered.
			break
		}
		if err := f.analyzeRune(buf, c); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

// analyzeRune processes any escape sequences occurred in a string literal.
func (f *fbsLex) analyzeRune(buf *bytes.Buffer, c rune) error {
	if c == 0 {
		return errors.New("null char not allowed in string literal")
	}
	if c == '\\' {
		if err := f.analyzeEscapeSequence(buf); err != nil {
			return err
		}
		return nil
	}
	buf.WriteRune(c)
	return nil
}

// analyzeEscapeSequence returns valid escape sequence occurred in a string literal.
// Could be:
// 1. octal: \233
// 2. hex:   \x7f
// 3. short unicode: \u2333
// 4. long unicode:  \U23332333
// 5. common escape sequence: \a \n \r ..
func (f *fbsLex) analyzeEscapeSequence(buf *bytes.Buffer) error {
	// escape sequence
	c, _, err := f.input.readRune()
	if err != nil {
		return err
	}
	f.adjustPos(c)
	if ok, err := f.readOctHexInString(buf, c); ok {
		return err
	}
	if ok, err := f.readUnicodeInString(buf, c); ok {
		return err
	}
	if err := f.readCommonEscapeSequence(buf, c); err != nil {
		return err
	}
	return nil
}

// setPre stores the previous token.
func (f *fbsLex) setPre(n ast.TerminalNode) {
	f.preSym = n
}

func (f *fbsLex) setRune(lval *fbsSymType, val rune) {
	lval.r = ast.NewRuneNode(val, f.newTokenInfo())
	f.setPre(lval.r)
}

func (f *fbsLex) setError(lval *fbsSymType, err error) {
	lval.err = f.wrappedError(err)
}

func (f *fbsLex) setInt(lval *fbsSymType, val uint64) {
	lval.i = ast.NewUintLiteralNode(val, f.newTokenInfo())
	f.setPre(lval.i)
}

func (f *fbsLex) setFloat(lval *fbsSymType, val float64) {
	lval.f = ast.NewFloatLiteralNode(val, f.newTokenInfo())
	f.setPre(lval.f)
}

func (f *fbsLex) setIdent(lval *fbsSymType, val string) {
	lval.id = ast.NewIdentNode(val, f.newTokenInfo())
	f.setPre(lval.id)
}

func (f *fbsLex) setString(lval *fbsSymType, val string) {
	lval.s = ast.NewStringLiteralNode(val, f.newTokenInfo())
	f.setPre(lval.s)
}

// adjustPos adjust position information such as line, column
// according to the given runes.
func (f *fbsLex) adjustPos(rs ...rune) {
	for _, r := range rs {
		switch r {
		case '\n': // new line
			f.col = 0 // back to beginning
			f.line++
		case '\r': // no action
		case '\t': // to next tab stop
			m := f.col % Tabsize
			f.col += Tabsize - m
		default:
			f.col++
		}
	}
}

func (f *fbsLex) pre() *ast.Position {
	if f.preSym == nil {
		return &ast.Position{
			Filename: f.filename,
			Offset:   0,
			Line:     1,
			Col:      1,
		}
	}
	return f.preSym.Start()
}

func (f *fbsLex) cur() ast.Position {
	return ast.Position{
		Filename: f.filename,
		Offset:   f.offset,
		Line:     f.line + 1,
		Col:      f.col + 1,
	}
}

func (f *fbsLex) posRange() ast.PosRange {
	return ast.PosRange{
		Start: ast.Position{
			Filename: f.filename,
			Offset:   f.preOffset,
			Line:     f.preLine + 1,
			Col:      f.preCol + 1,
		},
		End: f.cur(),
	}
}

func (f *fbsLex) newTokenInfo() ast.Token {
	return ast.Token{
		PosRange: f.posRange(),
		RawText:  f.input.endMark(),
	}
}

func (f *fbsLex) newComment() ast.Comment {
	ws := string(f.ws)
	f.ws = f.ws[:0]
	return ast.Comment{
		PosRange:          f.posRange(),
		LeadingWhitespace: ws,
		Text:              f.input.endMark(),
	}
}

func (f *fbsLex) skipToEndOfLineComment() bool {
	for {
		c, _, err := f.input.readRune()
		if err != nil {
			return false
		}
		if c == '\n' {
			return true
		}
		f.adjustPos(c)
	}
}

func (f *fbsLex) skipToEndOfBlockComment() bool {
	for {
		c, _, err := f.input.readRune()
		if err != nil {
			return false
		}
		f.adjustPos(c)
		if success, end := f.skipToEnd(c); end {
			return success
		}
	}
}

func (f *fbsLex) skipToEnd(c rune) (success, end bool) {
	if c == '*' {
		c, _, err := f.input.readRune()
		if err != nil {
			return false, true
		}
		if c == '/' {
			f.adjustPos(c)
			return true, true
		}
		f.input.unreadRune(c)
	}
	return true, false
}

func (f *fbsLex) wrappedError(err error) ErrorWithPos {
	ewp, ok := err.(ErrorWithPos)
	if !ok {
		ewp = ErrorWithPos{Pos: f.pre(), Err: err}
	}
	f.handler.handleError(ewp)
	return ewp
}
