package humanize

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		input    int64
		expected string
	}{
		{1000, "1 thousand"},
		{100000, "100 thousand"},
		{125000, "125 thousand"},
		{1000000, "1 million"},
		{10000000, "10 million"},
		{15600000, "15.6 million"},
		{1000000000, "1 billion"},
		{1000000000000, "1 trillion"},
		{1000000000000000, "1 quadrillion"},
		{1000000000000000, "1 quadrillion"},
		{10000000000000000, "10 quadrillion"},
		{9200000000000000000, "9.2 quintillion"},
	}

	for _, test := range tests {
		if actual := Int(test.input); actual != test.expected {
			t.Errorf("expected '%s' but got '%s'", test.expected, actual)
		}
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(3250000000)
	}
}
