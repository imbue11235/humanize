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
type Translator struct {
	code         string
	translations Map
}

// NewTranslator ...
func NewTranslator(code string, translations Map) *Translator {
	return &Translator{
		code:         code,
		translations: translations,
	}
}

// LanguageCode ...
func (t *Translator) LanguageCode() string {
	return t.code
}

func (t *Translator) getTranslation(path string) string {
	return t.getTranslationOrDefault(path, "")
}

func (t *Translator) getTranslationOrDefault(path, defaultValue string) string {
	value := t.get(path)
	if value == nil {
		return defaultValue
	}

	if stringValue, ok := value.(string); ok {
		return stringValue
	}

	return defaultValue
}

func (t *Translator) get(path string) interface{} {
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

// Translate ...
func (t *Translator) Translate(path string, args ...interface{}) string {
	translation := t.getTranslation(path)
	if len(args) > 0 {
		return fmt.Sprintf(translation, args...)
	}

	return translation
}

// Pluralize ...
func (t *Translator) Pluralize(path string, count int) string {
	translation := t.getTranslation(path)
	parts := strings.Split(translation, pluralSeparator)

	if count != 1 && len(parts) > 1 {
		return t.applyCountToTranslation(parts[1], count)
	}

	return parts[0]
}

func (t *Translator) applyCountToTranslation(translation string, count int) string {
	if strings.Contains(translation, "%d") {
		return fmt.Sprintf(translation, count)
	}

	return translation
}
