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
