package humanize

import (
	"testing"

	"github.com/imbue11235/humanize/locale"
	"github.com/imbue11235/humanize/locale/da"
	"github.com/imbue11235/humanize/locale/de"
	"github.com/imbue11235/humanize/locale/en"
)

func TestLocaleManager(t *testing.T) {
	if translation := translate("test"); translation != "" {
		t.Errorf("expected path `test` to be empty")
	}

	if err := RegisterLocale("test", locale.Map{"test": "response"}); err != nil {
		t.Errorf("expected `RegisterLocale` with `test` to not return an error. Error: %s", err)
	}

	if err := SetLocale("test"); err != nil {
		t.Errorf("expected `SetLocale` to not return an error. Error: %s", err)
	}

	if translation := translate("test"); translation == "" {
		t.Errorf("expected path `test` to be `response`")
	}

	if err := SetFallbackLocale("en"); err != nil {
		t.Errorf("expected `SetFallbackLocale` to not return an error. Error: %s", err)
	}

	if translation := translate("time.now"); translation == "" {
		t.Errorf("expected path `time.now` to fallback to `en` and resolve to `response`")
	}

	if err := SetLocale("en"); err != nil {
		t.Errorf("expected `SetLocale` to not return an error. Error: %s", err)
	}
}

func TestLocales(t *testing.T) {
	tests := []struct {
		code   string
		locale locale.Map
	}{
		{"en", en.Locale},
		{"da", da.Locale},
		{"de", de.Locale},
	}

	requiredPaths := []string{
		"slice.connector",
		"slice.rest",
		"time.now",
		"time.future",
		"time.past",
		"time.estimation.s",
		"time.estimation.m",
		"time.estimation.h",
		"time.estimation.d",
		"time.estimation.w",
		"time.estimation.M",
		"time.estimation.y",
		"time.estimation.D",
		"time.estimation.l",
		"time.precision.s",
		"time.precision.m",
		"time.precision.h",
		"time.precision.d",
		"time.precision.M",
		"time.precision.y",
		"int.K",
		"int.M",
		"int.B",
		"int.T",
		"int.Q",
		"int.Qi",
		"int.Sx",
		"int.Sp",
	}

	for _, test := range tests {
		// using a manager to test the paths
		testManager := locale.NewManager(locale.WithLocale(test.code, test.locale))

		for _, path := range requiredPaths {
			if value := testManager.Translate(path); value == "" {
				t.Errorf("locale `%s`: expected to find something at path `%s`, but got an empty string", test.code, path)
			}
		}
	}
}
