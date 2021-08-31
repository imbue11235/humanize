package language

import (
	"testing"
)

func TestTranslatorTranslate(t *testing.T) {
	manager := createManagerInstance(Map{
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
	})

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
		translation := manager.Translate(test.path, test.args...)
		if test.expected != translation {
			t.Errorf("expected translation '%s', but got '%s'", test.expected, translation)
		}
	}
}

func TestTranslatorPluralize(t *testing.T) {
	manager := createManagerInstance(Map{
		"dollar": "I only have 1 dollar|I have %d dollars!!!!",
		"time": Map{
			"minutes": "one minute|%d minutes",
		},
		"normal": "Everything is normal",
	})

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
		pluralization := manager.Pluralize(test.path, test.amount)
		if test.expected != pluralization {
			t.Errorf("expected pluralization '%s', but got '%s'", test.expected, pluralization)
		}
	}
}

func createManagerInstance(language Map) *Manager {
	return NewManager(WithLanguage("default", language))
}
