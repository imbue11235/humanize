package era

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestDurationToPreciseTimeUnits(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{2 * time.Second, "2s"},
		{5*time.Hour + 2*time.Minute, "5h 2m"},
		{22*time.Minute + 50*time.Second, "22m 50s"},
		{120 * time.Second, "2m"},
		{47 * time.Hour, "1d 23h"},
		{month, "1M"},
		{year, "1y"},
		{120 * year, "120y"},
	}

	for _, test := range tests {
		actual := join(DurationToPreciseTimeUnits(test.input))

		if actual != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, actual)
		}
	}
}

func BenchmarkDurationToPreciseTimeUnits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DurationToPreciseTimeUnits(5 * 10 * 12 * 30 * 24 * time.Hour)
	}
}

func join(results []*Result) string {
	var cases []string
	for _, result := range results {
		cases = append(cases, fmt.Sprintf("%d%s", result.Volume, result.Symbol))
	}

	return strings.Join(cases, " ")
}
