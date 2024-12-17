package Fibonacci

import (
	"testing"
)

// Fibonacci function
func fibonacci(n int) int {
	a := 0
	b := 1

	// return 0 for negative values
	if n < 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func Testfibonacci(t *testing.T) {
	// Define table-driven test cases
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "Testing Fibonacci sequence with a positive number",
			input: 5,
			want:  5, // The 5th number in the Fibonacci sequence is 5
		},
		{
			name:  "Testing Fibonacci sequence with zero",
			input: 0,
			want:  0, // The 0th number in the Fibonacci sequence is 0
		},
		{
			name:  "Testing Fibonacci sequence with a negative number",
			input: -5,
			want:  0, // The Fibonacci sequence is not defined for negative numbers
		},
	}

	// Execute test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := fibonacci(test.input); got != test.want {
				t.Errorf("fibonacci() = %v, want %v", got, test.want)
			} else {
				t.Logf("Success: %s", test.name)
			}
		})
	}
}
