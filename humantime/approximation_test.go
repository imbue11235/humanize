package humantime

import (
	"testing"
	"time"
)

func TestCalculateApproximateDuration(t *testing.T) {
	tests := []struct {
		input          time.Duration
		expectedSymbol string
		expectedCount  int
	}{
		{2 * time.Millisecond, "second", 1},
		{2 * time.Second, "second", 2},
		{22 * time.Minute, "minute", 22},
		{10 * time.Hour, "hour", 10},
		{time.Hour + 25*time.Minute, "hour", 1},
		{time.Hour + 50*time.Minute, "hour", 2},
		{48 * time.Hour, "day", 2},
		{3 * 7 * 24 * time.Hour, "week", 3},
		{3 * 6 * 24 * time.Hour, "week", 2},
		{30 * 24 * time.Hour, "month", 1},
		{58 * 24 * time.Hour, "month", 2},
		{2 * 12 * 30 * 24 * time.Hour, "year", 2},
		{11 * 12 * 30 * 24 * time.Hour, "decade", 1},
		{9 * 12 * 30 * 24 * time.Hour, "decade", 1},
		{5 * 10 * 12 * 30 * 24 * time.Hour, "long", 1},
	}

	for _, test := range tests {
		approximation := CalculateApproximateDuration(test.input)
		if approximation.Symbol != test.expectedSymbol {
			t.Errorf("expected symbol '%s', but got '%s'", test.expectedSymbol, approximation.Symbol)
		}

		if approximation.Count != test.expectedCount {
			t.Errorf("expected count '%d', but got '%d'", test.expectedCount, approximation.Count)
		}
	}
}

func BenchmarkCalculateApproximateDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateApproximateDuration(5 * 10 * 12 * 30 * 24 * time.Hour)
	}
}
