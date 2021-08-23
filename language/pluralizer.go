package language

import (
	"fmt"
	"strings"
)

type Pluralizer struct {
	one  string
	many string
}

func NewPluralizer(one, many string) *Pluralizer {
	return &Pluralizer{
		one:  one,
		many: many,
	}
}

// applyAmountToTemplate conditionally applies the given amount
// to the string
func (p *Pluralizer) applyAmountToTemplate(template string, amount int) string {
	if strings.Contains(template, "%d") {
		return fmt.Sprintf(template, amount)
	}

	return template
}

func (p *Pluralizer) Pluralize(amount int) string {
	if amount > 1 {
		return p.applyAmountToTemplate(p.many, amount)
	}

	return p.one
}
