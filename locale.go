package humanize

import (
	"github.com/imbue11235/humanize/locale"
	"github.com/imbue11235/humanize/locale/en"
)

var (
	manager *locale.Manager
)

func init() {
	m, _ := locale.NewManager(locale.WithLocale(en.Code, en.Locale))

	UseManager(m)
}

// UseManager sets the global used manager to the given manager
// useful for setting up a custom manager
func UseManager(m *locale.Manager) {
	manager = m
}

// RegisterLocale registers a new locale with the translation manager.
//
//		humanize.RegisterLocale("my-locale", locale.Map{...})
//
func RegisterLocale(code string, translations locale.Map) error {
	return manager.RegisterLocale(code, translations)
}

// SetLocale sets the locale on the translation manager instance
//
//		humanize.SetLocale("en")
//
func SetLocale(code string) error {
	return manager.SetLocale(code)
}

// SetFallbackLocale sets the fallback locale on the translation manager instance
//
//		humanize.SetFallbackLocale("en")
//
func SetFallbackLocale(code string) error {
	return manager.SetFallbackLocale(code)
}

// translate finds a translation from the given path, via the translation manager
func translate(path string, args ...interface{}) string {
	return manager.Translate(path, args...)
}

// pluralize finds a translation from given path, and pluralized it, via the
// translation manager, based on the count
func pluralize(path string, count int) string {
	return manager.Pluralize(path, count)
}
