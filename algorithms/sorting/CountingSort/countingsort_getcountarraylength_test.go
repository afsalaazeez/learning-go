package CountingSort

import (
	"testing"
)

// TestgetCountArrayLength is a test function for getCountArrayLength function
func TestgetCountArrayLength(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Scenario 1: Empty array",
			input:    []int{},
			expected: 1,
		},
		{
			name:     "Scenario 2: Array with one element",
			input:    []int{5},
			expected: 6,
		},
		{
			name:     "Scenario 3: Array with multiple elements",
			input:    []int{5, 10, 15, 20},
			expected: 21,
		},
		{
			name:     "Scenario 4: Array with negative elements",
			input:    []int{-5, -10, -15, -20},
			expected: 1,
		},
	}

	// Execute test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call getCountArrayLength function
			result := getCountArrayLength(tc.input)

			// Check if the result is as expected
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			} else {
				t.Logf("Success: Expected %v and got %v", tc.expected, result)
			}
		})
	}
}
