package era

import (
	"math"
	"testing"
	"time"
)

func TestDurationToEstimation(t *testing.T) {
	tests := []struct {
		input          time.Duration
		expectedSymbol Symbol
		expectedCount  int
	}{
		{2 * time.Millisecond, SymbolSecond, 0},
		{2 * time.Second, SymbolSecond, 2},
		{22 * time.Minute, SymbolMinute, 22},
		{10 * time.Hour, SymbolHour, 10},
		{time.Hour + 25*time.Minute, SymbolHour, 1},
		{time.Hour + 50*time.Minute, SymbolHour, 2},
		{48 * time.Hour, SymbolDay, 2},
		{3 * 7 * 24 * time.Hour, SymbolWeek, 3},
		{3 * 6 * 24 * time.Hour, SymbolWeek, 2},
		{30 * 24 * time.Hour, SymbolMonth, 1},
		{58 * 24 * time.Hour, SymbolMonth, 2},
		{2 * 12 * 30 * 24 * time.Hour, SymbolYear, 2},
		{11 * 12 * 30 * 24 * time.Hour, SymbolDecade, 1},
		{9 * 12 * 30 * 24 * time.Hour, SymbolDecade, 1},
		{5 * 10 * 12 * 30 * 24 * time.Hour, SymbolLong, 1},
		{time.Duration(math.NaN()), SymbolSecond, 0},
	}

	for _, test := range tests {
		estimation := DurationToEstimation(test.input)
		if estimation.Symbol != test.expectedSymbol {
			t.Errorf("expected symbol '%s', but got '%s'", test.expectedSymbol, estimation.Symbol)
		}

		if estimation.Volume != test.expectedCount {
			t.Errorf("expected count '%d', but got '%d'", test.expectedCount, estimation.Volume)
		}
	}
}

func BenchmarkDurationToEstimation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DurationToEstimation(5 * 10 * 12 * 30 * 24 * time.Hour)
	}
}
