package humanize

import (
	"github.com/imbue11235/humanize/language"
	"github.com/imbue11235/humanize/locales/en"
)

var (
	manager *language.Manager
)

func init() {
	manager = language.NewLocaleManager()
	manager.RegisterLocale(en.Locale)
}

func RegisterLocale(locale *language.Locale) error {
	return manager.RegisterLocale(locale)
}

func SetLocale(code string) error {
	return manager.SetLocale(code)
}
