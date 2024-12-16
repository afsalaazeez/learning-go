package calculator_test

import (
  "testing"
  "github.com/yourname/yourproject/calculator"
)

func TestSum(t *testing.T) {
    t.Run("sum of two positive integers", func(t *testing.T) {
        got := calculator.Sum(5, 6)
        want := 11
        if got != want {
            t.Errorf("got %d, want %d", got, want)
        }
    })

    t.Run("sum of two negative integers", func(t *testing.T) {
        got := calculator.Sum(-5, -6)
        want := -11
        if got != want {
            t.Errorf("got %d, want %d", got, want)
        }
    })

    t.Run("sum of positive and negative integer", func(t *testing.T) {
        got := calculator.Sum(-5, 6)
        want := 1
        if got != want {
            t.Errorf("got %d, want %d", got, want)
        }
    })

    t.Run("sum of zero and integer", func(t *testing.T) {
        got := calculator.Sum(0, 6)
        want := 6
        if got != want {
            t.Errorf("got %d, want %d", got, want)
        }
    })
}
