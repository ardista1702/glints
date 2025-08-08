package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"italem irad irigayaj", "melati dari jayagiri"},
		{"iadab itsap ulalreb", "badai pasti berlalu"},
		{"nalub kusutret gnalali", "bulan tertusuk ilalang"},
	}

	for _, tt := range tests {
		got := solve(tt.input)
		if got != tt.expected {
			t.Errorf("solve(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}
