package locale

// Option ...
type Option interface {
	apply(m *Manager)
}

// RegisterLocaleOption ...
type RegisterLocaleOption struct {
	code         string
	translations Map
}

func (o *RegisterLocaleOption) apply(m *Manager) {
	m.RegisterLocale(o.code, o.translations)

	// set this as default locale, if none is set
	if m.currentTranslator == nil {
		m.SetLocale(o.code)
	}

	// fallback will be set to this language
	// if it's not already set
	if m.fallbackTranslator == nil {
		m.SetFallbackLocale(o.code)
	}
}

// WithLocale ...
func WithLocale(code string, translations Map) *RegisterLocaleOption {
	return &RegisterLocaleOption{code, translations}
}

// FallbackLocaleOption ...
type FallbackLocaleOption struct {
	code         string
	translations Map
}

func (o *FallbackLocaleOption) apply(m *Manager) {
	m.RegisterLocale(o.code, o.translations)
	m.SetFallbackLocale(o.code)
}

// WithFallbackLocale ...
func WithFallbackLocale(code string, translations Map) *FallbackLocaleOption {
	return &FallbackLocaleOption{code, translations}
}

// FallbackStringOption ...
type FallbackStringOption struct {
	value string
}

func (o *FallbackStringOption) apply(m *Manager) {
	m.fallbackString = o.value
}

// WithFallbackString ...
func WithFallbackString(value string) *FallbackStringOption {
	return &FallbackStringOption{value}
}
