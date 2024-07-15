package bint

import "testing"

func TestAdder(t *testing.T) {
	t.Run("2 + 2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", sum, expected)
		}
	})
}
