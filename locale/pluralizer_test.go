package locale

import (
	"testing"
)

func TestPluralizer(t *testing.T) {
	tests := []struct {
		text        string
		boundaries  []*boundary
		shouldError bool
	}{
		{
			"[1] Some text|[2] some other text",
			[]*boundary{
				{1, 1, "Some text"},
				{2, 2, "some other text"},
			},
			false,
		},
		{
			"[1] Text 1|[2-*] Text 2",
			[]*boundary{
				{1, 1, "Text 1"},
				{2, -1, "Text 2"},
			},
			false,
		},
		{
			"[1] Text 1",
			[]*boundary{},
			true,
		},
		{
			"[0] Text 0|[1-*] Test 2",
			[]*boundary{
				{0, 0, "Text 0"},
				{1, -1, "Test 2"},
			},
			false,
		},
		{
			"[1-*] Text 1|[2-*] Text 2",
			[]*boundary{
				{1, -1, "Text 1"},
			},
			false,
		},
		{
			"text with no boundaries",
			[]*boundary{},
			true,
		},
		{
			"[*-*] Text 1|[2-*] Text 2",
			[]*boundary{},
			true,
		},
		{
			"[1-@] Text 1|[2-*] Text 2",
			[]*boundary{},
			true,
		},
		{
			"[-100, 5] Text 1|[2] Text 2",
			[]*boundary{},
			true,
		},
		{
			"[5-1] Text 1|[1] Text 2",
			[]*boundary{},
			true,
		},
		{
			"[2] Text 1|[1] Text 2",
			[]*boundary{
				{2, 2, "Text 1"},
				{3, 3, "Text 2"},
			},
			false,
		},
		{
			"[2-100] Text 1|[5-200] Text 2",
			[]*boundary{
				{2, 100, "Text 1"},
				{101, 200, "Text 2"},
			},
			false,
		},
		{
			"[2-100] Text 1|[5-95] Text 2",
			[]*boundary{
				{2, 100, "Text 1"},
				{101, 101, "Text 2"},
			},
			false,
		},
	}

	for _, test := range tests {
		instance, err := createPluralizer(test.text)

		if err == nil && test.shouldError {
			t.Error("Expected error, but got none")
			continue
		}

		if err != nil {
			if !test.shouldError {
				t.Errorf("Expected no error, got %s", err)
				continue
			}

			continue
		}

		if len(test.boundaries) != len(instance.boundaries) {
			t.Errorf("Expected %d boundaries, got %d", len(test.boundaries), len(instance.boundaries))
			continue
		}

		for i, boundary := range instance.boundaries {
			checkBoundaryEquality(t, test.boundaries[i], boundary)
		}
	}
}

func TestPluralizerApply(t *testing.T) {
	tests := []struct {
		actual, expected string
		count            int
	}{
		{"[1] Text %d|[2] No", "Text 1", 1},
		{"[1] Test 1|[2-*] %d bananas", "8 bananas", 8},
		{"[1-5] A few seconds|[6-100] A lot of seconds|[101-500] %d seconds", "A few seconds", 3},
		{"[1-5] A few seconds|[6-100] A lot of seconds|[101-500] %d seconds", "A lot of seconds", 88},
		{"[1-5] A few seconds|[6-100] A lot of seconds|[101-500] %d seconds", "404 seconds", 404},
		{"[1-5] A few seconds|[6-100] A lot of seconds|[101-500] %d seconds", "904 seconds", 904},
	}

	for _, test := range tests {
		instance, err := createPluralizer(test.actual)

		if err != nil {
			t.Fatal(err)
		}

		result := instance.apply(test.count)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func checkBoundaryEquality(t *testing.T, a, b *boundary) {
	if a.min != b.min {
		t.Errorf("expected to get min %d, got %d", a.min, b.min)
	}

	if a.max != b.max {
		t.Errorf("expected to get max %d, got %d", a.max, b.max)
	}

	if a.text != b.text {
		t.Errorf("expected to get text %s, got %s", a.text, b.text)
	}
}
