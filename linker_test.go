package fbs

import (
	"reflect"
	"testing"
)

func TestLinker(t *testing.T) {
	tests := []struct {
		name      string
		filenames []string
		wantLen   int
		wantErr   bool
	}{
		{
			name:      "duplicate table name",
			filenames: []string{"./fbsfiles/error_test/link_test1.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate table field name",
			filenames: []string{"./fbsfiles/error_test/link_test2.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate struct name",
			filenames: []string{"./fbsfiles/error_test/link_test3.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate struct field name",
			filenames: []string{"./fbsfiles/error_test/link_test4.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate enum name",
			filenames: []string{"./fbsfiles/error_test/link_test5.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate enum value name",
			filenames: []string{"./fbsfiles/error_test/link_test6.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate union name",
			filenames: []string{"./fbsfiles/error_test/link_test7.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate rpc_service name",
			filenames: []string{"./fbsfiles/error_test/link_test8.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "duplicate rpc method name",
			filenames: []string{"./fbsfiles/error_test/link_test9.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "table field type name undefined",
			filenames: []string{"./fbsfiles/error_test/link_test10.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "table field type name resolved to package name",
			filenames: []string{"./fbsfiles/error_test/link_test11.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "table field type name resolved to rpc service name",
			filenames: []string{"./fbsfiles/error_test/link_test12.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name: "duplicate symbol defined in two files",
			filenames: []string{
				"./fbsfiles/error_test/link_test13.fbs",
				"./fbsfiles/error_test/link_test14.fbs",
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name: "error message still the same even if given different file order",
			filenames: []string{
				"./fbsfiles/error_test/link_test14.fbs",
				"./fbsfiles/error_test/link_test13.fbs",
			},
			wantLen: 0,
			wantErr: true,
		},
		{
			name:      "struct field type name undefined",
			filenames: []string{"./fbsfiles/error_test/link_test15.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "method input type name undefined",
			filenames: []string{"./fbsfiles/error_test/link_test16.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "method input type is not table type",
			filenames: []string{"./fbsfiles/error_test/link_test17.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "method output type name resolved to package name",
			filenames: []string{"./fbsfiles/error_test/link_test18.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "method output type is not table type",
			filenames: []string{"./fbsfiles/error_test/link_test19.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
		{
			name:      "method input type name is already fully qualified",
			filenames: []string{"./fbsfiles/error_test/link_test20.fbs"},
			wantLen:   0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser()
			got, err := p.ParseFiles(tt.filenames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("ParseFiles() len(got) = %v, wantLen %v", got, tt.wantLen)
			}
		})
	}
}

func TestGetFullNamespaces(t *testing.T) {
	type args struct {
		nss []string
	}
	tests := []struct {
		name string
		args args
		want map[string]struct{}
	}{
		{name: "input is nil", args: args{nss: nil}, want: nil},
		{
			name: "normal case",
			args: args{
				nss: []string{"", "namespace1", "rpc.app.server"},
			},
			want: map[string]struct{}{
				"":               {},
				"namespace1":     {},
				"rpc":            {},
				"rpc.app":        {},
				"rpc.app.server": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllNamespaces(tt.args.nss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllNamespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
