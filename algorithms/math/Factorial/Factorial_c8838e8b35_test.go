package Factorial

import "testing"

func TestFactorial(t *testing.T) {
	t.Run("Positive number", func(t *testing.T) {
		num := 5
		expected := 120

		result := Factorial(num)

		if result != expected {
			t.Errorf("Factorial(%d) = %d; expected %d", num, result, expected)
		} else {
			t.Logf("Factorial(%d) = %d; expected %d", num, result, expected)
		}
	})

	t.Run("Zero", func(t *testing.T) {
		num := 0
		expected := 1

		result := Factorial(num)

		if result != expected {
			t.Errorf("Factorial(%d) = %d; expected %d", num, result, expected)
		} else {
			t.Logf("Factorial(%d) = %d; expected %d", num, result, expected)
		}
	})

	t.Run("Negative number", func(t *testing.T) {
		num := -5
		expected := 0

		result := Factorial(num)

		if result != expected {
			t.Errorf("Factorial(%d) = %d; expected %d", num, result, expected)
		} else {
			t.Logf("Factorial(%d) = %d; expected %d", num, result, expected)
		}
	})
}
