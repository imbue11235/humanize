package humanize

import (
	"github.com/imbue11235/humanize/locale"
	"github.com/imbue11235/humanize/locale/en"
)

var (
	manager *locale.Manager
)

func init() {
	UseManager(locale.NewManager(locale.WithLocale("en", en.Locale)))
}

// UseManager sets the global used manager to the given manager
// useful for setting up a custom manager
func UseManager(m *locale.Manager) {
	manager = m
}

// RegisterLocale registers a new locale with the translation manager.
func RegisterLocale(code string, translations locale.Map) error {
	return manager.RegisterLocale(code, translations)
}

// SetLocale sets the locale on the translation manager instance
func SetLocale(code string) error {
	return manager.SetLocale(code)
}

func translate(path string, args ...interface{}) string {
	return manager.Translate(path, args...)
}

func pluralize(path string, count int) string {
	return manager.Pluralize(path, count)
}
