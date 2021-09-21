package humanize

import (
	"testing"

	"github.com/imbue11235/humanize/locale"
	"github.com/imbue11235/humanize/locale/da"
	"github.com/imbue11235/humanize/locale/de"
	"github.com/imbue11235/humanize/locale/en"
)

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
