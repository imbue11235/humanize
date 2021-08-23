package language

import "testing"

func TestManagerRegisterLocale(t *testing.T) {
	manager := NewLocaleManager()
	locale := &Locale{Code: "test"}

	if err := manager.RegisterLocale(locale); err != nil {
		t.Errorf("expected no error, but got error: %s", err.Error())
	}

	if len(manager.registeredLocales) != 1 {
		t.Errorf("expected registered locales to be 1 but got %d", len(manager.registeredLocales))
	}

	if _, ok := manager.registeredLocales[locale.Code]; !ok {
		t.Errorf("expected to find registered locale '%s'", locale.Code)
	}

	if manager.Locale() == nil {
		t.Error("expected locale to be set")
	}

	if manager.Locale().Code != locale.Code {
		t.Errorf(
			"expected to find locale with code %s, but got locale with code %s",
			locale.Code,
			manager.Locale().Code,
		)
	}
}

func TestManagerSetLocale(t *testing.T) {
	manager := NewLocaleManager()
	locale := &Locale{Code: "test"}
	locale2 := &Locale{Code: "test2"}

	manager.RegisterLocale(locale)
	manager.RegisterLocale(locale2)

	// should not be able to set unregistered locales
	if err := manager.SetLocale("nb"); err == nil {
		t.Error("expected an error, but got none")
	}

	// check if switches work
	if err := manager.SetLocale(locale.Code); err != nil {
		t.Errorf("expected no error, but got error: %s", err.Error())
	}

	if manager.Locale().Code != locale.Code {
		t.Errorf("expected current locale to be %s, but it was %s", locale.Code, manager.Locale().Code)
	}

	if err := manager.SetLocale(locale2.Code); err != nil {
		t.Errorf("expected no error, but got error: %s", err.Error())
	}

	if manager.Locale().Code != locale2.Code {
		t.Errorf("expected current locale to be %s, but it was %s", locale2.Code, manager.Locale().Code)
	}
}
