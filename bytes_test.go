package humanize

import (
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{100, "100 B"},
		{10000, "10 kB"},
		{10000000, "10 MB"},
		{35324355, "35 MB"},
		{2000000000, "2.0 GB"},
		{1000000000000, "1.0 TB"},
	}

	for _, test := range tests {
		output := Bytes(test.input)
		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}
