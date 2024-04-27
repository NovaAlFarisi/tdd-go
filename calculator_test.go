package main

import "testing"

func TestCalculate(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		result, err := Calculate("+", 10, 5)
		if err != nil {
			t.Errorf("Error occured: %v", err)
		}

		expected := 15.0
		if result != expected {
			t.Errorf("Expected: %.2f, Got: %.2f", expected, result)
		} else {
			t.Logf("Test passed! Expected: %.2f, Got: %.2f", expected, result)
		}
	})
}
