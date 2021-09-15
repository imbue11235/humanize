package humanize

import (
	"testing"
)

func TestTruncFloat(t *testing.T) {
	tests := []struct {
		input, expected float64
		precision       int
	}{
		{1.242424, 1.2, 1},
		{5.320495, 5.32, 2},
		{8.324343, 8.324, 3},
		{10.2424266, 10.2424, 4},
		{555.32323335235, 555.32323, 5},
	}

	for _, test := range tests {
		if trunc := truncFloat(test.input, test.precision); trunc != test.expected {
			t.Errorf("expected `%f` but got `%f`", test.expected, trunc)
		}
	}
}
