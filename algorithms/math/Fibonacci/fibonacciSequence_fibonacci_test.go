package Fibonacci

import (
	"testing"
)

// Fibonacci function for generating Fibonacci sequence
func fibonacci(n int) int {
	a := 0
	b := 1

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
			want:  5,
		},
		{
			name:  "Testing Fibonacci sequence with zero",
			input: 0,
			want:  0,
		},
		{
			name:  "Testing Fibonacci sequence with a negative number",
			input: -5,
			want:  0,
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
