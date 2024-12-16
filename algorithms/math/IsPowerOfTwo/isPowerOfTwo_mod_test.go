package IsPowerOfTwo

import (
	"testing"
)

// Testmod is a testing function for mod function in the IsPowerOfTwo package
func Testmod(t *testing.T) {
	// Define test cases
	var tests = []struct {
		a, b, want int
	}{
		// Scenario 1: Adding two negative numbers
		{-3, -2, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := mod(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
