package language

import "fmt"

// Manager ...
type Manager struct {
	currentLanguage   string
	currentTranslator *Translator
	translators       map[string]*Translator
}

// NewManager ...
func NewManager() *Manager {
	return &Manager{
		translators: map[string]*Translator{},
	}
}

// Translate ...
func (m *Manager) Translate(path string, args ...interface{}) string {
	return m.currentTranslator.Translate(path, args...)
}

// Pluralize ...
func (m *Manager) Pluralize(path string, count int) string {
	return m.currentTranslator.Pluralize(path, count)
}

// RegisterLocale ...
func (m *Manager) RegisterLanguage(code string, translations Map) {
	m.translators[code] = NewTranslator(code, translations)
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
