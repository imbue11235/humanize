package language

import "fmt"

type Manager struct {
	currentLocale     *Locale
	registeredLocales map[string]*Locale
}

func NewLocaleManager() *Manager {
	return &Manager{
		registeredLocales: make(map[string]*Locale),
	}
}

func (m *Manager) Locale() *Locale {
	return m.currentLocale
}

func (m *Manager) RegisterLocale(locale *Locale) error {
	if _, ok := m.registeredLocales[locale.Code]; ok {
		return fmt.Errorf("locale with code '%s' is already registered", locale.Code)
	}

	m.registeredLocales[locale.Code] = locale

	// If no current language is set, we will make this the default one too
	if m.currentLocale == nil {
		return m.SetLocale(locale.Code)
	}

	return nil
}

func (m *Manager) SetLocale(code string) error {
	registeredLocale, ok := m.registeredLocales[code]

	if !ok {
		return fmt.Errorf("no locale with code '%s' is registered", code)
	}

	m.currentLocale = registeredLocale

	return nil
}
