package main

import "testing"

func TestIsPossibleZeroSum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Case 1: 55 6 3 100",
			input:    "55 6 3 100",
			expected: false,
		},
		{
			name:     "Case 2: 0 0 0 20 5 15",
			input:    "0 0 0 20 5 15",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := extractData(tc.input)
			result := isPossibleZeroSum(data)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
