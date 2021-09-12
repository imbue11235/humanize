package humanize

import (
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0 B"},
		{2, "2 B"},
		{100, "100 B"},
		{1000, "1.0 kB"},
		{10000, "10 kB"},
		{2500000, "2.5 MB"},
		{10000000, "10 MB"},
		{35324355, "35 MB"},
		{350000000, "350 MB"},
		{2000000000, "2.0 GB"},
		{1000000000000, "1.0 TB"},
		{1000000000000000, "1.0 PB"},
		{1000000000000000000, "1.0 EB"},
		{13400000000000000000, "13 EB"},
	}

	for _, test := range tests {
		output := Bytes(test.input)
		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}

func TestBinaryBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0 B"},
		{2, "2 B"},
		{100, "100 B"},
		{1024, "1.0 KiB"},
		{10000, "9.8 KiB"},
		{2500000, "2.4 MiB"},
		{10000000, "9.5 MiB"},
		{35324355, "34 MiB"},
		{350000000, "334 MiB"},
		{2000000000, "1.9 GiB"},
		{1000000000000, "931 GiB"},
		{1000000000000000, "909 TiB"},
		{1000000000000000000, "888 PiB"},
		{13400000000000000000, "12 EiB"},
	}

	for _, test := range tests {
		output := BinaryBytes(test.input)
		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}

func TestShortFormBinaryBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{0, "0B"},
		{2, "2B"},
		{100, "100B"},
		{1024, "1.0K"},
		{10000, "9.8K"},
		{10000000, "9.5M"},
		{35324355, "34M"},
		{350000000, "334M"},
		{2000000000, "1.9G"},
		{1000000000000, "931G"},
		{1000000000000000, "909T"},
		{1000000000000000000, "888P"},
		{13400000000000000000, "12E"},
	}

	for _, test := range tests {
		output := ShortFormBinaryBytes(test.input)
		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}
