package locale

import (
	"fmt"
)

// Manager ...
type Manager struct {
	currentLocaleCode  string
	currentTranslator  *translator
	fallbackTranslator *translator
	fallbackString     string
	translators        map[string]*translator
}

// NewManager ...
func NewManager(options ...option) *Manager {
	manager := &Manager{
		fallbackString: "",
		translators:    map[string]*translator{},
	}

	for _, opt := range options {
		opt.applyTo(manager)
	}

	return manager
}

func (m *Manager) getTranslationOrFallback(path string) string {
	translation, err := m.currentTranslator.getTranslation(path)

	if err != nil {
		// at this point, we will try to use the fallback translator
		// and as a last option, falling back to the default fallback string
		// if everything else fails
		return m.fallbackTranslator.getTranslationOrDefault(path, m.fallbackString)
	}

	return translation
}

func (m *Manager) getPluralizationOrFallback(path string, count int) string {
	pluralizer, err := m.currentTranslator.getPluralizer(path)

	if err != nil {
		// at this point, we will try to use the fallback translator
		// and as a last option, falling back to the default fallback string
		// if everything else fails
		pluralizer, err = m.fallbackTranslator.getPluralizer(path)

		if err != nil {
			return m.getTranslationOrFallback(path)
		}
	}

	return pluralizer.apply(count)
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
	return m.getPluralizationOrFallback(path, count)
}

func (m *Manager) addTranslator(code string, translations Map) *translator {
	translator := newTranslator(code, translations)

	m.translators[code] = translator

	return translator
}

// RegisterLocale ...
func (m *Manager) RegisterLocale(code string, translations Map) error {
	if _, ok := m.translators[code]; !ok {
		m.addTranslator(code, translations)

		return nil
	}

	return fmt.Errorf("a locale with code `%s` already exists", code)
}

func (m *Manager) setFallbackTranslator(fallbackTranslator *translator) {
	m.fallbackTranslator = fallbackTranslator
}

// SetFallbackLocale ...
func (m *Manager) SetFallbackLocale(code string) error {
	if foundTranslator, ok := m.translators[code]; ok {
		m.setFallbackTranslator(foundTranslator)

		return nil
	}

	return fmt.Errorf("could not find a locale with code `%s`", code)
}

func (m *Manager) setTranslator(nextTranslator *translator) {
	m.currentTranslator = nextTranslator
}

func (m *Manager) setLocaleCode(code string) {
	m.currentLocaleCode = code
}

// SetLocale ...
func (m *Manager) SetLocale(code string) error {
	if foundTranslator, ok := m.translators[code]; ok {
		m.setLocaleCode(code)
		m.setTranslator(foundTranslator)

		return nil
	}

	return fmt.Errorf("could not find a locale with code `%s`", code)
}
