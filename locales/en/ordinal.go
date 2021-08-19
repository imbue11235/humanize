package en

type OrdinalStrategy struct{}

func (o *OrdinalStrategy) Format() string {
	return "%d%s"
}

func (o *OrdinalStrategy) Indicator(number int) string {
	switch number % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}

	return ""
}
