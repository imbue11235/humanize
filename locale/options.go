package locale

// Option ...
type option interface {
	apply(m *Manager) error
}

// RegisterLocaleOption ...
type registerLocaleOption struct {
	code         string
	translations Map
}

func (r *registerLocaleOption) apply(m *Manager) error {
	if err := m.RegisterLocale(r.code, r.translations); err != nil {
		return err
	}

	// set this as default locale, if none is set
	if m.currentTranslator == nil {
		if err := m.SetLocale(r.code); err != nil {
			return err
		}
	}

	// fallback will be set to this language
	// if it's not already set
	if m.fallbackTranslator == nil {
		if err := m.SetFallbackLocale(r.code); err != nil {
			return err
		}
	}

	return nil
}

// WithLocale ...
func WithLocale(code string, translations Map) *registerLocaleOption {
	return &registerLocaleOption{code, translations}
}

type fallbackStringOption string

func (f fallbackStringOption) apply(m *Manager) error {
	m.fallbackString = string(f)

	return nil
}

// WithFallbackString ...
func WithFallbackString(value string) fallbackStringOption {
	return fallbackStringOption(value)
}
