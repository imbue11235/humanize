package language

import (
	"fmt"
	"strings"
)

const (
	pluralSeparator = "|"
	pathSeparator   = "."
)

// Translator ...
type translator struct {
	code         string
	translations Map
}

// NewTranslator ...
func newTranslator(code string, translations Map) *translator {
	return &translator{
		code:         code,
		translations: translations,
	}
}

func (t *translator) getTranslation(path string) (string, error) {
	translation := t.get(path)
	if translation == nil {
		return "", fmt.Errorf("could not find translation with path `%s`", path)
	}

	if stringValue, ok := translation.(string); ok {
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
