package humanize

import (
	"github.com/imbue11235/humanize/language"
	"github.com/imbue11235/humanize/locales/en"
)

var (
	manager *language.Manager
)

func init() {
	manager = language.NewManager()
	manager.RegisterLanguage("en", en.Language)
	manager.SetLanguage("en")
}

// RegisterLanguage ...
func RegisterLanguage(code string, translations language.Map) {
	manager.RegisterLanguage(code, translations)
}

// SetLanguage ...
func SetLanguage(code string) error {
	return manager.SetLanguage(code)
}

func translate(path string, args ...interface{}) string {
	return manager.Translate(path, args...)
}

func pluralize(path string, count int) string {
	return manager.Pluralize(path, count)
}
