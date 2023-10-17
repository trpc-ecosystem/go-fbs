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
	"fmt"
	"io"
	"strconv"
	"unicode/utf8"
)

// keywords stores the keywords of flatbuffers.
var keywords = map[string]int{
	"true":            True,
	"false":           False,
	"attribute":       Attribute,
	"bool":            Bool,
	"byte":            Byte,
	"double":          Double,
	"enum":            Enum,
	"file_extension":  FileExtension,
	"file_identifier": FileIdentifier,
	"float":           Float,
	"float32":         Float32,
	"float64":         Float64,
	"include":         Include,
	"inf":             Inf,
	"int":             Int,
	"int16":           Int16,
	"int32":           Int32,
	"int64":           Int64,
	"int8":            Int8,
	"long":            Long,
	"namespace":       Namespace,
	"nan":             Nan,
	"root_type":       RootType,
	"rpc_service":     RPCService,
	"short":           Short,
	"string":          String,
	"struct":          Struct,
	"table":           Table,
	"ubyte":           Ubyte,
	"uint":            Uint,
	"uint16":          Uint16,
	"uint32":          Uint32,
	"uint64":          Uint64,
	"uint8":           Uint8,
	"ulong":           Ulong,
	"union":           Union,
	"ushort":          Ushort,
}

// runeReader performs read on a single character.
type runeReader struct {
	r      *bufio.Reader
	marked []rune
	unread []rune
	err    error
}

// readRune reads a rune and returns its size.
func (r *runeReader) readRune() (rune, int, error) {
	if len(r.unread) > 0 {
		c := r.unread[len(r.unread)-1]
		r.unread = r.unread[:len(r.unread)-1]
		if r.marked != nil {
			r.marked = append(r.marked, c)
		}
		return c, utf8.RuneLen(c), nil
	}
	c, sz, err := r.r.ReadRune()
	if err != nil {
		r.err = err
	} else if r.marked != nil {
		r.marked = append(r.marked, c)
	}
	return c, sz, err
}

// unreadRune unreads a previously read rune.
// It is usually used after readRune to perform the
// same functionality of peek.
func (r *runeReader) unreadRune(c rune) {
	if r.marked != nil {
		if r.marked[len(r.marked)-1] != c {
			panic("unread rune should be the same with last marked rune")
		}
		r.marked = r.marked[:len(r.marked)-1]
	}
	r.unread = append(r.unread, c)
}

// startMark is used together with endMark to mark
// a series of runes which finally becomes a token.
func (r *runeReader) startMark(c rune) {
	r.marked = []rune{c}
}

// endMark is used together with startMark to mark
// a series of runes.
func (r *runeReader) endMark() string {
	m := string(r.marked)
	r.marked = r.marked[:0]
	return m
}

// readRune is a higher level wrapper of (*runeReader).readRune
// which reads runes as part of a string literal.
func (f *fbsLex) readRune() (rune, error) {
	c, _, err := f.input.readRune()
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return -1, err
	}
	if c == '\n' {
		return -1, errors.New("encounter end-of-line before end of string literal")
	}
	return c, err
}

// readIdentifier reads runes which finally form an identifier.
func (f *fbsLex) readIdentifier(sofar []rune) []rune {
	token := sofar
	for {
		c, _, err := f.input.readRune()
		if err != nil {
			break
		}
		if c != '_' && !isAlphaNum(c) {
			f.input.unreadRune(c)
			break
		}
		f.adjustPos(c)
		token = append(token, c)
	}
	return token
}

// readOctHexInString tries to read an octal or hexadecimal.
func (f *fbsLex) readOctHexInString(buf *bytes.Buffer, c rune) (bool, error) {
	if isHexStart(c) {
		if err := f.readHexInString(buf); err != nil {
			return true, err
		}
		return true, nil
	}
	if isOct(c) {
		if err := f.readOctInString(buf, c); err != nil {
			return true, err
		}
		return true, nil
	}
	return false, nil
}

// readUnicodeInString tries to read an unicode in string.
func (f *fbsLex) readUnicodeInString(buf *bytes.Buffer, c rune) (bool, error) {
	if c == 'u' {
		if err := f.readShortUnicodeInString(buf); err != nil {
			return true, err
		}
		return true, nil
	}
	if c == 'U' {
		if err := f.readLongUnicodeInString(buf); err != nil {
			return true, err
		}
		return true, nil
	}
	return false, nil
}

// readCommonEscapeSequence tries to read a common escape sequence.
func (f *fbsLex) readCommonEscapeSequence(buf *bytes.Buffer, c rune) error {
	switch c {
	case 'a':
		buf.WriteByte('\a')
		return nil
	case 'b':
		buf.WriteByte('\b')
		return nil
	case 'f':
		buf.WriteByte('\f')
		return nil
	case 'n':
		buf.WriteByte('\n')
		return nil
	case 'r':
		buf.WriteByte('\r')
		return nil
	case 'v':
		buf.WriteByte('\v')
		return nil
	case '\\':
		buf.WriteByte('\\')
		return nil
	case '\'':
		buf.WriteByte('\'')
		return nil
	case '"':
		buf.WriteByte('"')
		return nil
	case '?':
		buf.WriteByte('?')
		return nil
	default:
		return fmt.Errorf("invalid escape sequence: %q", "\\"+string(c))
	}
}

// readHexInString tries to read a hexadecimal in string.
func (f *fbsLex) readHexInString(buf *bytes.Buffer) error {
	// \x hex escape
	c, _, err := f.input.readRune()
	if err != nil {
		return err
	}
	f.adjustPos(c)
	c2, _, err := f.input.readRune()
	if err != nil {
		return err
	}
	var hex string
	if !isHexNum(c2) {
		f.input.unreadRune(c2)
		hex = string(c)
	} else {
		f.adjustPos(c2)
		hex = string([]rune{c, c2})
	}
	i, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return fmt.Errorf("invalid hex escape: \\x%q", hex)
	}
	buf.WriteByte(byte(i))
	return nil
}

// readOctInString tries to read an octcal in string.
func (f *fbsLex) readOctInString(buf *bytes.Buffer, c rune) error {
	// \2 octal escape
	c2, _, err := f.input.readRune()
	if err != nil {
		return err
	}
	if err := f.readOctal(buf, c, c2); err != nil {
		return err
	}
	return nil
}

// readOctal tries to read an octcal.
func (f *fbsLex) readOctal(buf *bytes.Buffer, c, c2 rune) error {
	var octal string
	if !isOct(c2) {
		f.input.unreadRune(c2)
		octal = string(c)
	} else {
		f.adjustPos(c2)
		var err error
		octal, err = f.readRuneForOctal(c, c2)
		if err != nil {
			return nil
		}
	}
	i, err := parseOctal(octal)
	if err != nil {
		return err
	}
	buf.WriteByte(byte(i))
	return nil
}

// readRuneForOctal retries to read a rune for octal.
func (f *fbsLex) readRuneForOctal(c, c2 rune) (string, error) {
	var octal string
	c3, _, err := f.input.readRune()
	if err != nil {
		return "", err
	}
	if !isOct(c3) {
		f.input.unreadRune(c3)
		octal = string([]rune{c, c2})
	} else {
		f.adjustPos(c3)
		octal = string([]rune{c, c2, c3})
	}
	return octal, nil
}

// readShortUnicodeInString tries to read a short unicode in string.
func (f *fbsLex) readShortUnicodeInString(buf *bytes.Buffer) error {
	// \u short unicode escape
	u := make([]rune, 4)
	for i := range u {
		c, _, err := f.input.readRune()
		if err != nil {
			return err
		}
		f.adjustPos(c)
		u[i] = c
	}
	i, err := strconv.ParseInt(string(u), 16, 32)
	if err != nil {
		return fmt.Errorf("invalid unicode escape: \\u%q", string(u))
	}
	buf.WriteRune(rune(i))
	return nil
}

// readLongUnicodeInString tries to read a long unicode in string.
func (f *fbsLex) readLongUnicodeInString(buf *bytes.Buffer) error {
	// \U long unicode escape
	u, err := f.readLongUnicodeRunes()
	if err != nil {
		return err
	}
	i, err := strconv.ParseInt(string(u), 16, 32)
	if err != nil {
		return fmt.Errorf("invalid unicode escape: \\U%q", string(u))
	}
	if i > 0x10ffff || i < 0 {
		return fmt.Errorf("unicode escape out of range, should in [0,0x10ffff]: \\U%q", string(u))
	}
	buf.WriteRune(rune(i))
	return nil
}

// readLongUnicodeRunes tries to read long unicode runes.
func (f *fbsLex) readLongUnicodeRunes() ([]rune, error) {
	u := make([]rune, 8)
	for i := range u {
		c, _, err := f.input.readRune()
		if err != nil {
			return nil, err
		}
		f.adjustPos(c)
		u[i] = c
	}
	return u, nil
}

// readHexNumber tries to read a hexadecimal number.
func (f *fbsLex) readHexNumber(sofar []rune) []rune {
	token := sofar
	for {
		c, _, err := f.input.readRune()
		if err != nil {
			break
		}
		if !isHexNum(c) {
			f.input.unreadRune(c)
			break
		}
		f.adjustPos(c)
		token = append(token, c)
	}
	return token
}

// readNumber tries to read a number.
func (f *fbsLex) readNumber(sofar []rune, allowDot, allowExp bool) []rune {
	runes := sofar
	for {
		c, _, err := f.input.readRune()
		if err != nil {
			break
		}
		var brk bool
		runes, allowDot, allowExp, brk = f.checkRunesForNumber(runes, c, allowDot, allowExp)
		if brk {
			break
		}
	}
	return runes
}

// checkRunesForNumber checks runes for a number.
func (f *fbsLex) checkRunesForNumber(runes []rune, c rune, allowDot, allowExp bool) (rs []rune, dot, exp, brk bool) {
	if c == '.' {
		if !allowDot {
			f.input.unreadRune(c)
			return runes, allowDot, allowExp, true
		}
		allowDot = false
	} else if isExp(c) {
		runes, c, allowExp, brk = f.checkHexRunesForNumber(runes, c, allowExp)
		if brk {
			return runes, allowDot, allowExp, brk
		}
	} else if !isNum(c) {
		f.input.unreadRune(c)
		return runes, allowDot, allowExp, true
	}
	f.adjustPos(c)
	runes = append(runes, c)
	return runes, allowDot, allowExp, false
}

// checkHexRunesForNumber checks hexadecimal for a number.
func (f *fbsLex) checkHexRunesForNumber(runes []rune, c rune, allowExp bool) ([]rune, rune, bool, bool) {
	if !allowExp {
		f.input.unreadRune(c)
		return runes, c, allowExp, true
	}
	allowExp = false
	cn, _, err := f.input.readRune()
	if err != nil {
		f.input.unreadRune(c)
		return runes, c, allowExp, true
	}
	if cn == '-' || cn == '+' {
		cnn, _, err := f.input.readRune()
		if err != nil {
			f.input.unreadRune(cn)
			f.input.unreadRune(c)
			return runes, c, allowExp, true
		}
		if !isNum(cnn) {
			f.input.unreadRune(cnn)
			f.input.unreadRune(cn)
			f.input.unreadRune(c)
			return runes, c, allowExp, true
		}
		f.adjustPos(c)
		runes = append(runes, c)
		c, cn = cn, cnn
	} else if !isNum(cn) {
		f.input.unreadRune(cn)
		f.input.unreadRune(c)
		return runes, c, allowExp, true
	}
	f.adjustPos(c)
	runes = append(runes, c)
	c = cn
	return runes, c, allowExp, false
}

// isNum checks whether the given rune is a number.
func isNum(c rune) bool {
	return c >= '0' && c <= '9'
}

// isAlpha checks whether the given rune is an alphabet.
func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// isAlphaNum checks whether the given rune is an alphabet or a number.
func isAlphaNum(c rune) bool {
	return isNum(c) || isAlpha(c)
}

// isOct checks whether the given rune is an octal.
func isOct(c rune) bool {
	return c >= '0' && c <= '7'
}

// isHex checks whether the given rune is a hexadecimal.
func isHex(c rune) bool {
	return (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'f')
}

// isHexNum checks whether the given rune is a hexadecimal or decimal number.
func isHexNum(c rune) bool {
	return isNum(c) || isHex(c)
}

// isHexStart checks whether the given rune can be a start of a hexadecimal number.
func isHexStart(c rune) bool {
	return c == 'x' || c == 'X'
}

// isExp checks whether the given rune can be an exponential.
func isExp(c rune) bool {
	return c == 'e' || c == 'E'
}

// parseOctal parses an octal number.
func parseOctal(octal string) (int64, error) {
	i, err := strconv.ParseInt(octal, 8, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid octal escape: \\%q", octal)
	}
	if i > 0xff {
		return 0, fmt.Errorf("octal escape out of range, should in [0,377]: \\%q", octal)
	}
	return i, nil
}
