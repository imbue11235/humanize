package language

import (
	"fmt"
	"strings"
)

// Manager ...
type Manager struct {
	currentLanguage    string
	currentTranslator  *translator
	fallbackTranslator *translator
	fallbackString     string
	translators        map[string]*translator
}

// NewManager ...
func NewManager(options ...Option) *Manager {
	manager := &Manager{
		fallbackString: "",
		translators:    map[string]*translator{},
	}

	for _, option := range options {
		option.apply(manager)
	}

	return manager
}

func (m *Manager) getTranslationOrFallback(path string) string {
	translation, err := m.currentTranslator.getTranslation(path)
	if err != nil {
		// at this point, we will just fall back to a default value of an empty string
		// as no suitable translation was found
		return m.fallbackTranslator.getTranslationOrDefault(path, m.fallbackString)
	}

	return translation
}

// Translate ...
func (m *Manager) Translate(path string, args ...interface{}) string {
	translation := m.getTranslationOrFallback(path)

	if len(args) > 0 {
		return fmt.Sprintf(translation, args...)
	}

	return translation
}

// Pluralize ...
func (m *Manager) Pluralize(path string, count int) string {
	translation := m.getTranslationOrFallback(path)
	parts := strings.Split(translation, pluralSeparator)

	if count != 1 && len(parts) > 1 {
		return m.applyCountToTranslation(parts[1], count)
	}

	return parts[0]
}

func (m *Manager) applyCountToTranslation(translation string, count int) string {
	if strings.Contains(translation, "%d") {
		return fmt.Sprintf(translation, count)
	}

	return translation
}

// RegisterLocale ...
func (m *Manager) RegisterLanguage(code string, translations Map) {
	m.translators[code] = newTranslator(code, translations)
}

// SetFallbackLanguage ...
func (m *Manager) SetFallbackLanguage(code string) error {
	if translator, ok := m.translators[code]; ok {
		m.fallbackTranslator = translator

		return nil
	}

	return fmt.Errorf("could not find a language with code `%s`", code)
}

// SetLanguage ...
func (m *Manager) SetLanguage(code string) error {
	if translator, ok := m.translators[code]; ok {
		m.currentLanguage = code
		m.currentTranslator = translator

		return nil
	}

	return fmt.Errorf("could not find a language with code `%s`", code)
}
