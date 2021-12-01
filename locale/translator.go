package locale

import (
	"fmt"
	"strings"
)

const (
	pathSeparator = "."
)

// Translator ...
type translator struct {
	code               string
	translations       Map
	pluralizationCache map[string]*pluralizer
	translationCache   map[string]string
}

// NewTranslator ...
func newTranslator(code string, translations Map) *translator {
	return &translator{
		code:               code,
		translations:       translations,
		pluralizationCache: map[string]*pluralizer{},
		translationCache:   map[string]string{},
	}
}

func (t *translator) getPluralizer(path string) (*pluralizer, error) {
	// check if a pluralizer has been cached
	// and return it if it has
	if pluralizer, ok := t.pluralizationCache[path]; ok {
		return pluralizer, nil
	}

	translation, err := t.getTranslation(path)
	if err != nil {
		return nil, err
	}

	pluralizer, err := createPluralizer(translation)

	if err != nil {
		return nil, err
	}

	// save pluralizer to cache
	t.pluralizationCache[path] = pluralizer

	return pluralizer, nil
}

func (t *translator) getTranslation(path string) (string, error) {
	// check if a translation has been cached
	// and return it if it has
	if translation, ok := t.translationCache[path]; ok {
		return translation, nil
	}

	translation := t.get(path)
	if translation == nil {
		return "", fmt.Errorf("could not find translation with path `%s`", path)
	}

	if stringValue, ok := translation.(string); ok {
		// save translation to cache
		t.translationCache[path] = stringValue

		return stringValue, nil
	}

	return "", fmt.Errorf("could not cast translation to string with path `%s`", path)
}

func (t *translator) getTranslationOrDefault(path, defaultValue string) string {
	translation, err := t.getTranslation(path)
	if err != nil {
		return defaultValue
	}

	return translation
}

func (t *translator) get(path string) interface{} {
	current := t.translations
	parts := strings.Split(path, pathSeparator)

	for index, part := range parts {
		if index == len(parts)-1 {
			return current[part]
		}

		switch current[part].(type) {
		case Map:
			current = current[part].(Map)
		}
	}

	return current
}
