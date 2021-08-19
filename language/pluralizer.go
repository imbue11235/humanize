package language

import (
	"fmt"
	"strings"
)

type Pluralizer struct {
	one string
	many string
}

func NewPluralizer(one, many string) *Pluralizer {
	return &Pluralizer{
		one: one,
		many: many,
	}
}

// applyAmountToTemplate conditionally applies the given amount
// to the string, if it contains a digit format `%d`
func (p *Pluralizer) applyAmountToTemplate(templateText string, amount int) string {
	if strings.Contains(templateText, "%d") {
		return fmt.Sprintf(templateText, amount)
	}

	return templateText
}

func (p *Pluralizer) Pluralize(amount int) string {
	if amount > 1 {
		return p.applyAmountToTemplate(p.many, amount)
	}

	return p.one
}

