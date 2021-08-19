package language

import "fmt"

type Time struct {
	Now           string
	Future        string
	Past          string
	Approximation map[string]*Pluralizer
	Exact         map[string]*Pluralizer
}

func (t Time) getTimeText(symbol string, dictionary map[string]*Pluralizer) (*Pluralizer, error) {
	text, ok := dictionary[symbol]
	if !ok {
		return nil, fmt.Errorf("could not find pluralizer from symbol %s", symbol)
	}

	return text, nil
}

func (t Time) GetApproximateTimeText(symbol string) (*Pluralizer, error) {
	return t.getTimeText(symbol, t.Approximation)
}

func (t Time) GetExactTimeText(symbol string) (*Pluralizer, error) {
	return t.getTimeText(symbol, t.Exact)
}
