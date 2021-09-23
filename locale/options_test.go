package locale

import "testing"

func TestWithLocale(t *testing.T) {
	manager := NewManager()

	localeMap := Map{
		"test": "123",
	}

	WithLocale("test", localeMap).applyTo(manager)

	if manager.currentLocaleCode != "test" {
		t.Error("expected current locale to be `test`")
	}

	if len(manager.translators) != 1 {
		t.Errorf("expected manager to have 1 registered translator, found `%d`", len(manager.translators))
	}

	if manager.fallbackTranslator == nil {
		t.Error("expected fallback translator to be set")
	}
}

func TestWithFallbackString(t *testing.T) {
	manager := NewManager()

	WithFallbackString("hey").applyTo(manager)

	if manager.fallbackString != "hey" {
		t.Errorf("expected fallback string to be `hey`, but got %s", manager.fallbackString)
	}
}
