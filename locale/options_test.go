package locale

import "testing"

func TestWithLocale(t *testing.T) {
	manager, _ := NewManager()

	localeMap := Map{
		"test": "123",
	}

	localeOption := WithLocale("test", localeMap)

	if err := localeOption.apply(manager); err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if manager.currentLocale != "test" {
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
	manager, _ := NewManager()

	fallbackOption := WithFallbackString("hey")

	if err := fallbackOption.apply(manager); err != nil {
		t.Errorf("expected no error, but got: %s", err)
	}

	if manager.fallbackString != "hey" {
		t.Errorf("expected fallback string to be `hey`, but got %s", manager.fallbackString)
	}
}
