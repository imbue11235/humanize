package language

// Option ...
type Option interface {
	apply(m *Manager)
}

// LanguageOption ...
type RegisterLanguageOption struct {
	code         string
	translations Map
}

func (o *RegisterLanguageOption) apply(m *Manager) {
	m.RegisterLanguage(o.code, o.translations)
	m.SetLanguage(o.code)

	// fallback will be set to this language
	// if it's not already set
	if m.fallbackTranslator == nil {
		m.SetFallbackLanguage(o.code)
	}
}

// WithInitialLanguage ...
func WithLanguage(code string, translations Map) *RegisterLanguageOption {
	return &RegisterLanguageOption{code, translations}
}

// FallbackLanguageOption ...
type FallbackLanguageOption struct {
	code         string
	translations Map
}

func (o *FallbackLanguageOption) apply(m *Manager) {
	m.RegisterLanguage(o.code, o.translations)
	m.SetFallbackLanguage(o.code)
}

// WithFallbackLanguage ...
func WithFallbackLanguage(code string, translations Map) *FallbackLanguageOption {
	return &FallbackLanguageOption{code, translations}
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
