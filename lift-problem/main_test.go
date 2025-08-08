package main

import (
	"testing"
)

func TestElevator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Case 1 - request di lantai 1",
			input:    "0 1 0 0 0",
			expected: 1, 
		},
		{
			name:     "Case 2 - beberapa request",
			input:    "2 5 0 0 4",
			expected: 4,
		},
		{
			name:     "Case 3 - tidak ada request",
			input:    "0 0 0 0",
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			filtered := filterRequest(tc.input)
			got := elevator(filtered)
			if got != tc.expected {
				t.Errorf("input %q → got %d, want %d", tc.input, got, tc.expected)
			}
		})
	}
}
