package en

import (
	"github.com/imbue11235/humanize/language"
)

var Locale = &language.Locale{
	Slice: language.Slice{
		Connector: "and",
		Rest:      language.NewPluralizer("one other", "%d others"),
	},
	Time: language.Time{
		Now:    "just now",
		Future: "in %s",
		Past:   "%s ago",
		Approximation: map[string]*language.Pluralizer{
			"second": language.NewPluralizer("a second", "%d seconds"),
			"minute": language.NewPluralizer("a minute", "%d minutes"),
			"hour":   language.NewPluralizer("an hour", "%d hours"),
			"day":    language.NewPluralizer("a day", "%d hours"),
			"week":   language.NewPluralizer("a week", "%d weeks"),
			"month":  language.NewPluralizer("a month", "%d months"),
			"year":   language.NewPluralizer("a year", "%d years"),
			"decade": language.NewPluralizer("a decade", "%d decades"),
			"long":   language.NewPluralizer("a long time", "a long time"),
		},
		Exact: map[string]*language.Pluralizer{
			"second": language.NewPluralizer("1 second", "%d seconds"),
			"minute": language.NewPluralizer("1 minute", "%d minutes"),
			"hour":   language.NewPluralizer("1 hour", "%d hours"),
			"day":    language.NewPluralizer("1 day", "%d hours"),
			"month":  language.NewPluralizer("1 month", "%d months"),
			"year":   language.NewPluralizer("1 year", "%d years"),
		},
	},
	OrdinalStrategy: &OrdinalStrategy{},
}
