package locale

import (
	"testing"
)

func TestManagerTranslate(t *testing.T) {
	manager := NewManager(WithLocale("default", Map{
		"flowers": Map{
			"rose":   "Rose",
			"lily":   "Lily",
			"tulip":  "Tulip",
			"orchid": "Orchid",
		},
		"trees": Map{
			"birch": "Birch",
			"oak":   "Oak",
		},
		"timber": "We have %d pieces of timber",
	}))

	tests := []struct {
		path, expected string
		args           []interface{}
	}{
		{"flowers", "", nil},
		{"flowers.lily", "Lily", nil},
		{"flowers.tulip", "Tulip", nil},
		{"trees.oak", "Oak", nil},
		{"bushes.hortensia", "", nil},
		{"a.non.existing.long.path", "", nil},
		{"timber", "We have 5 pieces of timber", []interface{}{5}},
	}

	for _, test := range tests {
		if translation := manager.Translate(test.path, test.args...); test.expected != translation {
			t.Errorf("expected translation `%s`, but got `%s`", test.expected, translation)
		}
	}
}

func TestManagerPluralize(t *testing.T) {
	manager := NewManager(WithLocale("default", Map{
		"dollar": "I only have 1 dollar|I have %d dollars!!!!",
		"time": Map{
			"minutes": "one minute|%d minutes",
		},
		"normal": "Everything is normal",
	}))

	tests := []struct {
		path, expected string
		amount         int
	}{
		{"dollar", "I only have 1 dollar", 1},
		{"dollar", "I have 231 dollars!!!!", 231},
		{"time.minutes", "one minute", 1},
		{"time.minutes", "2 minutes", 2},
		{"normal", "Everything is normal", 1},
		{"normal", "Everything is normal", 100},
	}

	for _, test := range tests {
		if pluralization := manager.Pluralize(test.path, test.amount); test.expected != pluralization {
			t.Errorf("expected pluralization `%s`, but got `%s`", test.expected, pluralization)
		}
	}
}

func TestManagerSetLocale(t *testing.T) {
	manager := NewManager(
		WithLocale("1", Map{}),
		WithLocale("2", Map{}),
	)

	if manager.currentLocale != "1" {
		t.Errorf("expected current locale to be `1`, but it was `%s`", manager.currentLocale)
	}

	manager.SetLocale("2")

	if manager.currentLocale != "2" {
		t.Errorf("expected current locale to be `2`, but it was `%s`", manager.currentLocale)
	}
}
