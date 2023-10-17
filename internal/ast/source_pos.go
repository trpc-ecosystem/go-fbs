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

package ast

import "fmt"

// Position denotes a location in a .fbs file.
type Position struct {
	Filename string
	Line     int
	Col      int
	Offset   int
}

// String implements Stringer interface.
func (s Position) String() string {
	if s.Line <= 0 || s.Col <= 0 {
		return s.Filename
	}
	return fmt.Sprintf("%s:%d:%d", s.Filename, s.Line, s.Col)
}

// PosRange denotes a range of positions in a .fbs
// file which indicates some region of the file.
type PosRange struct {
	Start Position
	End   Position
}

// Comment denotes a line of block of comments.
type Comment struct {
	PosRange
	LeadingWhitespace string
	Text              string
}
