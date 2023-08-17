package ast_test

import (
	"testing"

	"trpc.group/trpc-go/fbs/internal/ast"
)

func TestSourcePos(t *testing.T) {
	tests := []struct {
		name string
		s    ast.Position
		want string
	}{
		{
			name: "normal case",
			s: ast.Position{
				Filename: "file1.fbs",
				Line:     9,
				Col:      8,
				Offset:   3,
			},
			want: "file1.fbs:9:8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("Position.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
