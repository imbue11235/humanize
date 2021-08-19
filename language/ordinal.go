package language

type OrdinalStrategy interface {
	Format() string
	Indicator(number int) string
}
