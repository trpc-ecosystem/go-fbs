package fbs

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestLexer(t *testing.T) {
	filename := "fbsfiles/lexer_test.fbs"
	file, err := os.Open(filename)
	assert.Nil(t, err)
	l := newLexer(file, "", newErrorHandler())
	var sym fbsSymType
	expected := []struct {
		t    int // token
		l    int // line
		c    int // column
		span int
		v    interface{} // value
	}{
		{t: FloatLit, l: 3, c: 5, span: 4, v: 0.666},
		{t: '.', l: 3, c: 10, span: 1},
		{t: Ident, l: 3, c: 11, span: 3, v: "app"},
		{t: FloatLit, l: 3, c: 15, span: 2, v: 0.7},
		{t: FloatLit, l: 3, c: 17, span: 2, v: 0.6},
		{t: FloatLit, l: 3, c: 20, span: 4, v: 90.0},
		{t: FloatLit, l: 3, c: 25, span: 5, v: 0.009},
		{t: FloatLit, l: 3, c: 31, span: 2, v: 0.9},
		{t: Ident, l: 3, c: 33, span: 2, v: "EF"},
		{t: FloatLit, l: 3, c: 36, span: 2, v: 0.9},
		{t: Ident, l: 3, c: 38, span: 1, v: "F"},
		{t: True, l: 3, c: 40, span: 4, v: "true"},
		{t: IntLit, l: 3, c: 45, span: 1, v: uint64(1)},
		{t: Ident, l: 3, c: 46, span: 2, v: "ee"},
		{t: FloatLit, l: 3, c: 49, span: 3, v: 10.0},
		{t: Ident, l: 3, c: 52, span: 1, v: "e"},
		{t: IntLit, l: 5, c: 5, span: 4, v: uint64(0xff)},
		{t: IntLit, l: 5, c: 10, span: 1, v: uint64(0)},
		{t: Ident, l: 5, c: 11, span: 3, v: "xzz"},
		{t: IntLit, l: 5, c: 15, span: 2, v: uint64(0)},
		{t: FloatLit, l: 5, c: 18, span: 3, v: 0.9},
		{t: FloatLit, l: 5, c: 22, span: 20, v: float64(18446744073709551616)},
		{t: StrLit, l: 7, c: 5, span: 25, v: "a simple string literal"},
		{t: StrLit, l: 9, c: 5, span: 31, v: "\xef \177 -Д- 😂"},
		{t: StrLit, l: 11, c: 5, span: 32, v: "\a \b \f \n \r \v \\ ' \" ? "},
		{t: Struct, l: 18, c: 5, span: 6, v: "struct"},
		{t: '{', l: 18, c: 12, span: 1},
		{t: '}', l: 18, c: 13, span: 1},
		{t: '/', l: 18, c: 17, span: 1},
		{t: '.', l: 20, c: 5, span: 1},
	}
	for i, exp := range expected {
		tok := l.Lex(&sym)
		if tok == 0 {
			t.Fatalf("lexer report EOF instead of %v", exp)
		}
		var n ast.Node
		var val interface{}
		switch tok {
		case True, Ident, Struct:
			n = sym.id
			val = sym.id.Val
		case StrLit:
			n = sym.s
			val = sym.s.Val
		case IntLit:
			n = sym.i
			val = sym.i.Val
		case FloatLit:
			n = sym.f
			val = sym.f.Val
		default:
			n = sym.r
			val = nil
		}
		assert.Equal(t, exp.t, tok, "case %d: wrong token type %v, expected %v", i, tok, exp.t)
		assert.Equal(t, exp.v, val, "case %d: wrong token value %v, expected %v", i, val, exp.v)
		assert.Equal(t, exp.l, n.Start().Line, "case %d: wrong line number %v, expected %v", i, n.Start().Line, exp.l)
		assert.Equal(t, exp.c, n.Start().Col, "case %d: wrong column number %v, expected %v", i, n.Start().Col, exp.c)
		assert.Equal(t, exp.l, n.End().Line, "case %d: wrong end line number %v, expected %v", i, n.End().Line, exp.l)
		assert.Equal(t, exp.c+exp.span, n.End().Col, "case %d: wrong end column number %v, expected %v", i, n.End().Col, exp.c+exp.span)
	}
}

func TestLexerError(t *testing.T) {
	expected := []struct {
		str    string
		errMsg string
	}{
		{str: ".1e1000", errMsg: "value out of range"},
		{str: "0xFFFFFFFFFFFFFFFFF", errMsg: "value out of range"},
		{str: "1.1e1000", errMsg: "value out of range"},
		{str: `"partial string literal`, errMsg: "unexpected EOF"},
		{str: "\"string with\nsdf", errMsg: "encounter end-of-line before end of string literal"},
		{str: "\"string\000", errMsg: "null char not allowed in string literal"},
		{str: `"\H"`, errMsg: "invalid escape sequence"},
		{str: `"\x`, errMsg: "EOF"},
		{str: `"\xx`, errMsg: "EOF"},
		{str: `"\x6x`, errMsg: "EOF"},
		{str: `"\xx9`, errMsg: "invalid hex escape"},
		{str: `"\0`, errMsg: "EOF"},
		{str: `"\00`, errMsg: "EOF"},
		{str: `"\09`, errMsg: "EOF"},
		{str: `"\019`, errMsg: "EOF"},
		{str: `"\777"`, errMsg: "octal escape out of range"},
		{str: `"\ukkkk`, errMsg: "invalid unicode escape"},
		{str: `"\uk`, errMsg: "EOF"},
		{str: `"\U`, errMsg: "EOF"},
		{str: `"\Ukkkkkkkk`, errMsg: "invalid unicode escape"},
		{str: `"\U00ffffff`, errMsg: "unicode escape out of range"},
		{str: `/*`, errMsg: "block comment never terminates"},
		{str: `/**`, errMsg: "block comment never terminates"},
		{str: `"\`, errMsg: "EOF"},
	}
	for i, c := range expected {
		l := newLexer(strings.NewReader(c.str), "", newErrorHandler())
		var sym fbsSymType
		tok := l.Lex(&sym)
		assert.Equal(t, Error, tok)
		assert.NotNil(t, sym.err)
		assert.Contains(t, sym.err.Error(), c.errMsg, "case %d: expected %q, got: %q", i, c.errMsg, sym.err.Error())
	}
}

func TestReadEOF(t *testing.T) {
	expected := []struct {
		str string
		t   int // token
	}{
		{str: "0", t: IntLit},
		{str: "0x", t: IntLit},
		{str: "/", t: '/'},
		{str: "ident", t: Ident},
		{str: "1e", t: IntLit},
		{str: "1e+", t: IntLit},
		{str: "1e+x", t: IntLit},
	}
	for i, exp := range expected {
		l := newLexer(strings.NewReader(exp.str), "", newErrorHandler())
		var sym fbsSymType
		tok := l.Lex(&sym)
		assert.Equal(t, exp.t, tok, "case %d: wrong token type %v, expected %v", i, tok, exp.t)
		assert.Nil(t, sym.err)
	}
}

func TestLexerErrorHandler(t *testing.T) {
	handler := newErrorHandler()
	handler.err = errors.New("handler error")
	l := newLexer(strings.NewReader(""), "", handler)
	var sym fbsSymType
	tok := l.Lex(&sym)
	assert.Equal(t, 0, tok)
}
