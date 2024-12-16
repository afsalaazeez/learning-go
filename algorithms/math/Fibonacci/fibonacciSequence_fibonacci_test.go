package Fibonacci

import (
	"testing"
)

// Testfibonacci : Unit Test for fibonacci function
func Testfibonacci(t *testing.T) {
	// Table driven tests for different scenarios
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "Testing for 0",
			args: 0,
			want: 0,
		},
		{
			name: "Testing for 1",
			args: 1,
			want: 1,
		},
		{
			name: "Testing for 2",
			args: 2,
			want: 1,
		},
		{
			name: "Testing for 9",
			args: 9,
			want: 34,
		},
		{
			name: "Testing for 10",
			args: 10,
			want: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}

	// Edge case where fibonacci sequence cannot be generated
	t.Run("Negative Input", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		// This should cause the function to panic
		fibonacci(-1)
	})
}
