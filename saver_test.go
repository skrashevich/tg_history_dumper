package main

import (
	"runtime"
	"testing"
)

func Test_escapeNameForFS(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{"Unix Forward Slash", "file/name", "file_name"},
		{"Colon", "file:name", "file_name"},
		{"No Special Characters", "filename", "filename"},
		// Add more test cases as needed
	}
	if runtime.GOOS == "windows" {
		tests = append(tests, struct {
			name string
			args string
			want string
		}{"Windows Characters", `file\<>:"|*?name`, "file_________name"})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeNameForFS(tt.args); got != tt.want {
				t.Errorf("escapeNameForFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
