package humanize

import (
	"testing"
)

func TestFraction(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{2.65, "2 13/20"},
		{2.625, "2 5/8"},
		{2, "2"},
		{7, "7"},
		{8.9, "8 9/10"},
		{0.3, "3/10"},
		{1.66, "1 33/50"},
		{0, "0"},
		{1, "1"},
		{-0.1, "-1/10"},
		{-1.1, "-1 1/10"},
		{-1.625, "-1 5/8"},
		{-6.89, "-6 89/100"},
	}

	for _, test := range tests {
		fraction := Fraction(test.input)

		if fraction != test.expected {
			t.Errorf("expected '%s' to be '%s'", fraction, test.expected)
		}
	}
}

func BenchmarkFraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fraction(1.35)
	}
}
