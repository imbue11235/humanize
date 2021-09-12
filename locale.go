package humanize

import (
	"github.com/imbue11235/humanize/locale"
	"github.com/imbue11235/humanize/locale/en"
)

var (
	manager *locale.Manager
)

func init() {
	UseManager(locale.NewManager(
		locale.WithLocale("en", en.Locale),
		locale.WithFallbackLocale("en", en.Locale),
	))
}

// UseManager ...
func UseManager(m *locale.Manager) {
	manager = m
}

// RegisterLocale ...
func RegisterLocale(code string, translations locale.Map) {
	manager.RegisterLocale(code, translations)
}

// SetLocale ...
func SetLocale(code string) error {
	return manager.SetLocale(code)
}

func translate(path string, args ...interface{}) string {
	return manager.Translate(path, args...)
}

func pluralize(path string, count int) string {
	return manager.Pluralize(path, count)
}
