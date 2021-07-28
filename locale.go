package humanize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

var (
	language = "en"
	currentLocale *Locale
	loadedLocales = map[string]*Locale{}
	availableLocales = map[string]string{
		"en": "locales/en.yml",
		"da": "locales/da.yml",
	}
)

type Locale struct {
	Language string
	Common CommonTranslations
	Time TimeTranslations
}

type CommonTranslations struct {
	Connector string
}

type TimeTranslations struct {
	Now string
	Future string
	Past string
	Approximation map[string]string
	Exact map[string]string
}

func init() {
	// set current locale to `en`
	locale, _ := loadLocale("en")
	currentLocale = locale
}

func SetLocale(language string) error {
	locale, ok := loadedLocales[language]

	if ok {
		currentLocale = locale
		return nil
	}

	loadedLocale, err := loadLocale(language)

	if err != nil {
		return err
	}

	currentLocale = loadedLocale

	return nil
}

func AddLocale(language string, locale *Locale) error {
	_, ok := loadedLocales[language]

	if ok {
		return fmt.Errorf("a locale for the language %s already exists", language)
	}

	loadedLocales[language] = locale

	return nil
}

func loadLocale(language string) (*Locale, error) {
	availableLocale, ok := availableLocales[language]

	if !ok {
		return nil, fmt.Errorf("no locale with the language %s found", language)
	}

	f, err := os.Open(availableLocale)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	locale := &Locale{}
	if err := yaml.UnmarshalStrict(data, locale); err != nil {
		return nil, err
	}

	return locale, nil
}