package BubbleSort

import (
	"reflect"
	"testing"
)

// TestSwap is a test function for the Swap function in the BubbleSort package
func TestSwap(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		arr      []int
		i        int
		j        int
		expected []int
	}{
		{
			name:     "Test 1: Swap elements at indices 0 and 1",
			arr:      []int{2, 1, 3, 4, 5},
			i:        0,
			j:        1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test 2: Swap elements at indices 2 and 3",
			arr:      []int{1, 2, 4, 3, 5},
			i:        2,
			j:        3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Test 3: Swap elements at indices 3 and 4",
			arr:      []int{1, 2, 3, 5, 4},
			i:        3,
			j:        4,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function with test inputs
			swap(tt.arr, tt.i, tt.j)
			// Check if the output matches the expected result
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("swap() = %v, want %v", tt.arr, tt.expected)
			}
		})
	}
}
