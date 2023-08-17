package fbs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestErrors(t *testing.T) {
	handler := newErrorHandler()
	msg := "this is a message"
	pos := &ast.Position{
		Filename: "file1.fbs",
		Line:     7,
		Col:      8,
		Offset:   9,
	}
	err := handler.handleErrorWithPos(pos, "an error occurred with msg: %v", msg)
	assert.NotNil(t, err)
	assert.Equal(t, "file1.fbs:7:8: an error occurred with msg: this is a message", err.Error())
	e, ok := err.(ErrorWithPos)
	assert.True(t, ok)
	assert.Equal(t, "an error occurred with msg: this is a message", e.Unwrap().Error())
	assert.Equal(t, *pos, e.GetPos())
	e1 := handler.handleError(err)
	assert.NotNil(t, e1)
	assert.Equal(t, err, e1)
}
