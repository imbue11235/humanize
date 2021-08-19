package humanize

import (
	"errors"
	"fmt"
	"time"

	"github.com/imbue11235/humanize/humantime"
	"github.com/imbue11235/humanize/language"
)

type humanizerFunc func(builder *TimeBuilder, duration time.Duration) (string, error)

type TimeBuilder struct {
	translations  language.Time
	origin        time.Time
	humanizerFunc humanizerFunc
}

func newTimeBuilder(origin time.Time, humanizerFunc humanizerFunc) *TimeBuilder {
	return &TimeBuilder{
		origin:        origin,
		humanizerFunc: humanizerFunc,
		translations:  manager.Locale().Time,
	}
}

func Time(origin time.Time) *TimeBuilder {
	return newTimeBuilder(origin, humanizeApproximateDifference)
}

func ExactTime(origin time.Time) *TimeBuilder {
	return newTimeBuilder(origin, humanizeExactDifference)
}

func (t *TimeBuilder) From(from time.Time) (string, error) {
	return t.humanize(from, t.origin)
}

func (t *TimeBuilder) FromNow() (string, error) {
	return t.From(time.Now())
}

func (t *TimeBuilder) To(to time.Time) (string, error) {
	return t.humanize(t.origin, to)
}

func (t *TimeBuilder) ToNow() (string, error) {
	return t.To(time.Now())
}

func (t *TimeBuilder) humanize(from, to time.Time) (string, error) {
	// find the difference between given times
	difference := to.Sub(from)

	// if there is zero seconds in difference,
	// we can treat it as happening right "now"
	if difference.Seconds() == 0 {
		return t.translations.Now, nil
	}

	humanization, err := t.humanizerFunc(t, difference)

	if err != nil {
		return "", err
	}

	// if the difference is positive, it's in the future
	if difference > 0 {
		return fmt.Sprintf(t.translations.Future, humanization), nil
	}

	// else it's in the past
	return fmt.Sprintf(t.translations.Past, humanization), nil
}

func humanizeApproximateDifference(builder *TimeBuilder, duration time.Duration) (string, error) {
	// find the closest approximated "time distance" from given difference
	approximation := humantime.CalculateApproximateDuration(duration)

	if approximation == nil {
		return "", errors.New("could not calculate approximation")
	}

	text, err := builder.translations.GetApproximateTimeText(approximation.Threshold.Symbol)
	if err != nil {
		return "", err
	}

	return text.Pluralize(approximation.Count), err
}

func humanizeExactDifference(builder *TimeBuilder, duration time.Duration) (string, error) {
	results := humantime.CalculateExactDuration(duration)

	var output []string
	for _, result := range results {
		text, err := builder.translations.GetExactTimeText(result.Threshold.Symbol)
		if err != nil {
			return "", err
		}

		output = append(output, text.Pluralize(result.Count))
	}

	return Slice(output), nil
}
