package humanize

import (
	"errors"
	"fmt"
	"time"

	"github.com/imbue11235/humanize/humantime"
)

type humanizerFunc func(duration time.Duration) (string, error)

// TimeBuilder ...
type TimeBuilder struct {
	origin        time.Time
	humanizerFunc humanizerFunc
}

// Time ...
func Time(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin:        origin,
		humanizerFunc: humanizeApproximateDifference,
	}
}

// ExactTime ...
func ExactTime(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin:        origin,
		humanizerFunc: humanizePreciseDifference,
	}
}

// From ...
func (t *TimeBuilder) From(from time.Time) (string, error) {
	return t.humanize(from, t.origin)
}

// FromNow ...
func (t *TimeBuilder) FromNow() (string, error) {
	return t.From(time.Now())
}

// To ...
func (t *TimeBuilder) To(to time.Time) (string, error) {
	return t.humanize(t.origin, to)
}

// ToNow ...
func (t *TimeBuilder) ToNow() (string, error) {
	return t.To(time.Now())
}

func (t *TimeBuilder) humanize(from, to time.Time) (string, error) {
	// find the difference between given times
	difference := to.Sub(from)

	// if there is zero seconds in difference,
	// we can treat it as happening right "now"
	if difference.Seconds() == 0 {
		return translate("time.now"), nil
	}

	humanization, err := t.humanizerFunc(difference)
	if err != nil {
		return "", err
	}

	// if the difference is positive, it's in the future
	if difference > 0 {
		return translate("time.future", humanization), nil
	}

	// else it's in the past
	return translate("time.past", humanization), nil
}

func humanizeApproximateDifference(duration time.Duration) (string, error) {
	// find the closest approximated "time distance" from given difference
	approximation := humantime.CalculateApproximateDuration(duration)
	if approximation == nil {
		return "", errors.New("could not calculate approximation")
	}

	path := fmt.Sprintf("time.approximation.%s", approximation.Symbol)
	return pluralize(path, approximation.Count), nil
}

func humanizePreciseDifference(duration time.Duration) (string, error) {
	results := humantime.CalculatePreciseDuration(duration)

	var output []string
	for _, result := range results {
		path := fmt.Sprintf("time.precision.%s", result.Symbol)
		output = append(output, pluralize(path, result.Count))
	}

	return Slice(output), nil
}
