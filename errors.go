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
	"fmt"

	"trpc.group/trpc-go/fbs/internal/ast"
)

// ErrorWithPos wrap error with position information in the
// source file.
type ErrorWithPos struct {
	Err error
	Pos *ast.Position
}

// Error implements the error interface.
func (e ErrorWithPos) Error() string {
	sourcePos := e.GetPos()
	return fmt.Sprintf("%s: %v", sourcePos, e.Err)
}

// GetPos returns position information of error.
func (e ErrorWithPos) GetPos() ast.Position {
	if e.Pos == nil {
		return ast.Position{Filename: "<input>"}
	}
	return *e.Pos
}

// Unwrap retrieves the original error.
func (e ErrorWithPos) Unwrap() error {
	return e.Err
}

// errorHandler stores error to be handled.
type errorHandler struct {
	err error
}

// newErrorHandler creates an error handler.
func newErrorHandler() *errorHandler {
	return &errorHandler{}
}

// handleErrorWithPos is used mostly by parser and linker to mark error position.
func (e *errorHandler) handleErrorWithPos(pos *ast.Position, format string, args ...interface{}) error {
	if e.err != nil {
		return e.err
	}
	err := errorWithPos(pos, format, args...)
	e.err = err
	return err
}

// handleError is used mostly by lexer. The passed in error err is already set with
// position information.
func (e *errorHandler) handleError(err error) error {
	if e.err != nil {
		return e.err
	}
	e.err = err
	return err
}

// getError returns the underlying error in error handler.
func (e *errorHandler) getError() error {
	return e.err
}

// errorWithPos create an ErrorWithPos out of position information and customized message.
func errorWithPos(pos *ast.Position, format string, args ...interface{}) ErrorWithPos {
	return ErrorWithPos{Pos: pos, Err: fmt.Errorf(format, args...)}
}
