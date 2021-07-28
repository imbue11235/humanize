package humanize

import (
	"errors"
	"fmt"
	"github.com/imbue11235/humanize/pkg/humantime"
	"strconv"
	"strings"
	"time"
)

const intMax = int(^uint(0) >> 1)

type TimeBuilder struct {
	origin time.Time
	humanizerFunc func(duration time.Duration) (string, error)
}

func Time(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin: origin,
		humanizerFunc: humanizeApproximateDifference,
	}
}

func ExactTime(origin time.Time) *TimeBuilder {
	return &TimeBuilder{
		origin: origin,
		humanizerFunc: humanizeExactDifference,
	}
}

func (t *TimeBuilder) From(from time.Time) (string, error) {
	return t.humanize(from, t.origin)
}

func (t *TimeBuilder) FromNow() (string, error) {
	return t.humanize(time.Now(), t.origin)
}

func (t *TimeBuilder) To(to time.Time) (string, error) {
	return t.humanize(t.origin, to)
}

func (t *TimeBuilder) ToNow() (string, error) {
	return t.humanize(t.origin, time.Now())
}

func (t *TimeBuilder) humanize(from, to time.Time) (string, error) {
	// find the difference between given times
	difference := to.Sub(from)

	// if there is zero seconds in difference,
	// we can treat it as happening right "now"
	if difference.Seconds() == 0 {
		return currentLocale.Time.Now, nil
	}

	humanization, err := t.humanizerFunc(difference)

	if err != nil {
		return "", err
	}

	// if the difference is positive, it's in the future
	if difference > 0 {
		return fmt.Sprintf(currentLocale.Time.Future, humanization), nil
	}

	// else it's in the past
	return fmt.Sprintf(currentLocale.Time.Past, humanization), nil
}

func parseBoundsFromTemplate(templateText string) (lowerBound int, upperBound int, rest string, err error) {
	// split on last bracket
	splitted := strings.Split(templateText, "]")

	// if len is not 2 (the bounds, and the "rest")
	// we can safely assume that the format is incorrect
	if len(splitted) != 2 {
		err = errors.New("template format is incorrect")
		return
	}

	// define rest of the template string without the bounds
	rest = strings.TrimSpace(splitted[1])

	// parse lower bound
	lowerBound, err = strconv.Atoi(string(splitted[0][1]))
	if err != nil {
		return
	}

	// parse upper bound
	upperBound = intMax
	upperBoundString := string(splitted[0][3])
	if upperBoundString == "*" {
		return
	}

	// if upperbound not infinite aka "*"
	// we will try to parse it
	upperBound, err = strconv.Atoi(upperBoundString)

	return
}

func applyResultToTemplate(result *humantime.Result, templateText string) string {
	if strings.Contains(templateText, "%d") {
		return fmt.Sprintf(templateText, result.Count)
	}

	return templateText
}

func getHumanizationFromResult(result *humantime.Result, dictionary map[string]string) (string, error) {
	unitDescription, ok := dictionary[result.Threshold.Symbol]

	if !ok {
		return "", fmt.Errorf("could not find unit translation with symbol %s", result.Threshold.Symbol)
	}

	templates := strings.Split(unitDescription, "|")
	for _, template := range templates {
		lowerBound, upperBound, rest, err := parseBoundsFromTemplate(template)

		if err != nil {
			return "", err
		}

		if result.Count >= lowerBound && result.Count <= upperBound {
			return applyResultToTemplate(result, rest), nil
		}
	}

	return "", nil
}

func humanizeApproximateDifference(duration time.Duration) (string, error) {
	// find the closest approximated "time distance" from given difference
	approximation := humantime.CalculateApproximateDuration(duration)

	if approximation == nil {
		return "", errors.New("could not calculate approximation")
	}

	return getHumanizationFromResult(approximation, currentLocale.Time.Approximation)
}

func humanizeExactDifference(duration time.Duration) (string, error) {
	results := humantime.CalculateExactDuration(duration)

	var output []string
	length := 0
	for _, result := range results {
		humanization, err := getHumanizationFromResult(result, currentLocale.Time.Approximation)

		if err != nil {
			return "", err
		}

		output = append(output, humanization)
		length++
	}

	return Slice(output), nil
}