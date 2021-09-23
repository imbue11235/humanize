package locale

// Option ...
type option interface {
	applyTo(m *Manager)
}

// RegisterLocaleOption ...
type registerLocaleOption struct {
	code         string
	translations Map
}

func (r *registerLocaleOption) applyTo(m *Manager) {
	addedTranslator := m.addTranslator(r.code, r.translations)

	// set this as default locale, if none is set
	if m.currentTranslator == nil {
		m.setLocaleCode(r.code)
		m.setTranslator(addedTranslator)
	}

	// fallback will be set to this language
	// if it's not already set
	if m.fallbackTranslator == nil {
		m.setFallbackTranslator(addedTranslator)
	}
}

// WithLocale ...
func WithLocale(code string, translations Map) *registerLocaleOption {
	return &registerLocaleOption{code, translations}
}

type fallbackStringOption string

func (f fallbackStringOption) applyTo(m *Manager) {
	m.fallbackString = string(f)
}

// WithFallbackString ...
func WithFallbackString(value string) fallbackStringOption {
	return fallbackStringOption(value)
}
