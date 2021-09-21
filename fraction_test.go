package humanize

import (
	"math"
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
		{0.25, "1/4"},
		{0.3, "3/10"},
		{0.333, "333/1000"},
		{1.66, "1 33/50"},
		{0, "0"},
		{1, "1"},
		{-0.1, "-1/10"},
		{-1.1, "-1 1/10"},
		{-1.625, "-1 5/8"},
		{-6.89, "-6 89/100"},
		{0.5, "1/2"},
		{math.NaN(), "NaN"},
	}

	for _, test := range tests {
		if actual := Fraction(test.input); actual != test.expected {
			t.Errorf("expected '%s' to be '%s'", actual, test.expected)
		}
	}
}

func BenchmarkFraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fraction(1.35)
	}
}
