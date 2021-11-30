package humanize

import (
	"fmt"
	"time"

	"github.com/imbue11235/humanize/era"
)

type timeHumanizerFunc func(from, to time.Time) string

// TimeBuilder is a struct to build a time query
type TimeBuilder struct {
	origin        time.Time
	timeHumanizer timeHumanizerFunc
}

// Time creates a time builder using estimation as a
// humanizer
func Time(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin:        origin,
		timeHumanizer: humanizeEstimation,
	}
}

// ExactTime creates a time builder using exact difference
// as a humanizer
func ExactTime(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin:        origin,
		timeHumanizer: humanizeDifference,
	}
}

// Duration converts a duration into a string
// representation, loosely estimating the duration
// e.g. 1 month and 2 days = 1 month
func Duration(duration time.Duration) string {
	// find the closest estimated "time distance" from given difference
	estimation := era.DurationToEstimation(duration)

	path := fmt.Sprintf("time.estimation.%s", estimation.Symbol)
	return pluralize(path, estimation.Volume)
}

// ExactDuration converts a duration into a string
// representation, calculating the exact duration
// e.g. 28 hours = 1 day and 4 hours
func ExactDuration(duration time.Duration) string {
	return concatResults("time.precision", era.DurationToPreciseTimeUnits(duration))
}

// From defines the `origin` time
// on the time builder
func (t *TimeBuilder) From(from time.Time) string {
	return t.humanize(from, t.origin)
}

// FromNow is a utility function, which is like calling
// `Time(a).From(b)`, but where `b` is automatically set to `time.Now()`
func (t *TimeBuilder) FromNow() string {
	return t.From(time.Now())
}

// To defines the `destination` time
// on the time builder
func (t *TimeBuilder) To(to time.Time) string {
	return t.humanize(t.origin, to)
}

// ToNow is a utility function, which is like calling
// `Time(a).To(b)`, but where `b` is automatically set to `time.Now()`
func (t *TimeBuilder) ToNow() string {
	return t.To(time.Now())
}

// normalizeTimezone ensures that from and to time are in the same
// timezone by converting one to the other's timezone if needed
func (t *TimeBuilder) normalizeTimezone(from, to time.Time) (time.Time, time.Time) {
	if from.Location() != to.Location() {
		from.In(to.Location())
	}

	return from, to
}

func (t *TimeBuilder) humanize(from, to time.Time) string {
	// if there is zero seconds in difference,
	// we can treat it as happening right "now"
	if to.Equal(from) {
		return translate("time.now")
	}

	humanization := t.timeHumanizer(t.normalizeTimezone(from, to))

	// if the difference is positive, it's in the future
	if to.After(from) {
		return translate("time.future", humanization)
	}

	// else it's in the past
	return translate("time.past", humanization)
}

func humanizeEstimation(from, to time.Time) string {
	return Duration(to.Sub(from))
}

func humanizeDifference(from, to time.Time) string {
	return concatResults("time.precision", era.Difference(from, to))
}

func concatResults(path string, results []*era.Result) string {
	var output []string
	for _, result := range results {
		output = append(output, pluralize(fmt.Sprintf("%s.%s", path, result.Symbol), result.Volume))
	}

	return Slice(output)
}
