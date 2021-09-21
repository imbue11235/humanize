package humanize

import (
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0"},
		{5, "5"},
		{999, "999"},
		{1000, "1 thousand"},
		{100000, "100 thousand"},
		{125000, "125 thousand"},
		{785030, "785 thousand"},
		{1000000, "1 million"},
		{1589035, "1.6 million"},
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

func TestIntWithSuffix(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{125000, "125K"},
		{1000000, "1M"},
		{10000000, "10M"},
		{15600000, "15.6M"},
		{1000000000, "1B"},
		{1000000000000, "1T"},
		{1000000000000000, "1Q"},
		{1000000000000000, "1Q"},
		{10000000000000000, "10Q"},
		{9200000000000000000, "9.2Qi"},
	}

	for _, test := range tests {
		if actual := IntWithSymbol(test.input); actual != test.expected {
			t.Errorf("expected '%s' but got '%s'", test.expected, actual)
		}
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int(3250000000)
	}
}
