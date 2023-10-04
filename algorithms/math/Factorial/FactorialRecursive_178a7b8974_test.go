package Factorial

import "testing"

func TestFactorialRecursive_178a7b8974(t *testing.T) {
	// Test case 1: Factorial of 0 should be 1
	input1 := 0
	expectedOutput1 := 1
	actualOutput1 := FactorialRecursive(input1)
	if actualOutput1 != expectedOutput1 {
		t.Errorf("Test case 1 failed: FactorialRecursive(%d) = %d, expected %d", input1, actualOutput1, expectedOutput1)
	} else {
		t.Logf("Test case 1 passed: FactorialRecursive(%d) = %d", input1, actualOutput1)
	}

	// Test case 2: Factorial of a positive number
	input2 := 5
	expectedOutput2 := 120
	actualOutput2 := FactorialRecursive(input2)
	if actualOutput2 != expectedOutput2 {
		t.Errorf("Test case 2 failed: FactorialRecursive(%d) = %d, expected %d", input2, actualOutput2, expectedOutput2)
	} else {
		t.Logf("Test case 2 passed: FactorialRecursive(%d) = %d", input2, actualOutput2)
	}
}
