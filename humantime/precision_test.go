package humantime

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCalculatePreciseDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{2 * time.Second, "2 second"},
		{5*time.Hour + 2*time.Minute, "5 hour 2 minute"},
		{22*time.Minute + 50*time.Second, "22 minute 50 second"},
		{120 * time.Second, "2 minute"},
	}

	for _, test := range tests {
		actual := join(CalculatePreciseDuration(test.input))

		if actual != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, actual)
		}
	}
}

func BenchmarkCalculatePreciseDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculatePreciseDuration(5 * 10 * 12 * 30 * 24 * time.Hour)
	}
}

func join(results []*Result) string {
	var cases []string
	for _, result := range results {
		cases = append(cases, fmt.Sprintf("%d %s", result.Count, result.Symbol))
	}

	return strings.Join(cases, " ")
}
