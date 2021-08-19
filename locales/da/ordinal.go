package da

type OrdinalStrategy struct{}

func (o *OrdinalStrategy) Format() string {
	return "%d%s"
}

func (o *OrdinalStrategy) Indicator(number int) string {
	return ""
}
