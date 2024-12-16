package SquareRoot

import (
	"testing"
	"math"
)

// Testnewton function implements the table-driven tests
func Testnewton(t *testing.T) {
	// Define test data structures
	type args struct {
		z float64
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Adding two negative numbers",
			args: args{z: -2, x: -2},
			want: -1, // Expected result based on the Newton method formula
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function newton with test parameters
			if got := newton(tt.args.z, tt.args.x); math.Abs(got-tt.want) > 1e-6 {
				t.Errorf("newton() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: newton() = %v, want %v", got, tt.want)
			}
		})
	}
}
