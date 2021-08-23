package da

import (
	"github.com/imbue11235/humanize/language"
)

var Locale = &language.Locale{
	Code: "da",
	Slice: language.Slice{
		Connector: "og",
		Rest:      language.NewPluralizer("én anden", "%d andre"),
	},
	Time: language.Time{
		Now:    "lige nu",
		Future: "om %s",
		Past:   "%s siden",
		Approximation: map[string]*language.Pluralizer{
			"second": language.NewPluralizer("et sekund", "%d sekunder"),
			"minute": language.NewPluralizer("et minut", "%d minutter"),
			"hour":   language.NewPluralizer("en time", "%d timer"),
			"day":    language.NewPluralizer("en dag", "%d dage"),
			"week":   language.NewPluralizer("en uge", "%d uger"),
			"month":  language.NewPluralizer("en måned", "%d måneder"),
			"year":   language.NewPluralizer("et år", "%d år"),
			"decade": language.NewPluralizer("et årti", "%d årtier"),
			"long":   language.NewPluralizer("lang tid", "lang tid"),
		},
		Exact: map[string]*language.Pluralizer{
			"second": language.NewPluralizer("1 sekund", "%d sekunder"),
			"minute": language.NewPluralizer("1 minut", "%d minutter"),
			"hour":   language.NewPluralizer("1 time", "%d timer"),
			"day":    language.NewPluralizer("1 dag", "%d dage"),
			"month":  language.NewPluralizer("1 måned", "%d måneder"),
			"year":   language.NewPluralizer("1 år", "%d år"),
		},
	},
	OrdinalStrategy: &OrdinalStrategy{},
}
