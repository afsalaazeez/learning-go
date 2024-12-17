package main

import (
	"testing"
)

func mod(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Testmod(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name        string
		a           int
		b           int
		expect      int
		expectPanic bool
	}{
		{
			name:        "Test with positive a and b values",
			a:           10,
			b:           3,
			expect:      1,
			expectPanic: false,
		},
		{
			name:        "Test with negative a and positive b values",
			a:           -10,
			b:           3,
			expect:      -1,
			expectPanic: false,
		},
		{
			name:        "Test with positive a and negative b values",
			a:           10,
			b:           -3,
			expect:      1,
			expectPanic: false,
		},
		{
			name:        "Test with negative a and b values",
			a:           -10,
			b:           -3,
			expect:      -1,
			expectPanic: false,
		},
		{
			name:        "Test with b equal to zero",
			a:           10,
			b:           0,
			expect:      0,
			expectPanic: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Errorf("mod() panics with %v - want panic = %v", r, tc.expectPanic)
					}
				} else if tc.expectPanic {
					t.Errorf("mod() didn't panic - want panic = %v", tc.expectPanic)
				}
			}()

			if got := mod(tc.a, tc.b); got != tc.expect && !tc.expectPanic {
				t.Errorf("mod(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.expect)
			}
		})
	}
}

