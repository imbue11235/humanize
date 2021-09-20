package humanize

import (
	"testing"
	"time"
)

type timeTest struct {
	t1, t2, expected string
}

func runTimeTest(t *testing.T, tests []timeTest, handler func(a time.Time, b time.Time) string) {
	for _, test := range tests {
		if actual := handler(parseTime(t, test.t1), parseTime(t, test.t2)); actual != test.expected {
			t.Errorf("expected '%s' but got '%s'", test.expected, actual)
		}
	}
}

func TestTimeTo(t *testing.T) {
	tests := []timeTest{
		{"2021-01-01T22:00:00+00:00", "2021-01-01T22:00:00+00:00", "just now"},
		{"2021-01-01T22:00:00+00:00", "2021-01-02T22:00:00+00:00", "in a day"},
		{"2021-01-02T22:00:00+00:00", "2021-01-01T22:00:00+00:00", "a day ago"},
		{"2019-01-01T22:00:00+00:00", "2021-01-01T22:00:00+00:00", "in 2 years"},
		{"2021-01-02T22:00:00+00:00", "2021-01-24T22:00:00+00:00", "in 3 weeks"},
		{"2021-01-02T22:00:00+00:00", "2021-05-05T22:00:00+00:00", "in 4 months"},
		{"2021-01-02T22:00:00+00:00", "2065-05-05T22:00:00+00:00", "in a long time"},
		{"2021-01-02T22:00:00+00:00", "2035-05-05T22:00:00+00:00", "in a decade"},
		{"2020-01-01T22:00:00+00:00", "2021-02-01T22:00:00+00:00", "in a year"},
	}

	runTimeTest(t, tests, func(t1 time.Time, t2 time.Time) string {
		return Time(t1).To(t2)
	})
}

func TestTimeFrom(t *testing.T) {
	tests := []timeTest{
		{"2020-01-01T22:00:00+00:00", "2021-02-01T22:00:00+00:00", "a year ago"},
		{"2021-05-05T22:00:00+00:00", "2021-05-10T22:00:00+00:00", "5 days ago"},
		{"2016-01-01T00:00:00+00:00", "2018-05-01T00:00:00+00:00", "2 years ago"},
	}

	runTimeTest(t, tests, func(t1 time.Time, t2 time.Time) string {
		return Time(t1).From(t2)
	})
}

func TestExactTimeTo(t *testing.T) {
	tests := []timeTest{
		{"2021-01-01T22:00:00+00:00", "2021-01-01T22:00:00+00:00", "just now"},
		{"2021-01-01T22:00:00+00:00", "2021-01-02T22:00:00+00:00", "in 1 day"},
		{"2021-01-01T22:00:00+00:00", "2021-01-05T23:20:21+00:00", "in 4 days, 1 hour, 20 minutes and 21 seconds"},
		{"2020-01-01T22:00:00+00:00", "2021-01-01T22:00:00+00:00", "in 1 year"},
		{"2019-01-01T22:00:00+00:00", "2020-01-01T22:00:00+00:00", "in 1 year"},
		{"2021-01-02T22:00:00+00:00", "2022-05-05T22:00:00+00:00", "in 1 year, 4 months and 3 days"},
		{"2020-05-03T15:00:00+00:00", "2021-05-08T18:30:00+00:00", "in 5 days, 3 hours and 30 minutes"},
	}

	runTimeTest(t, tests, func(t1 time.Time, t2 time.Time) string {
		return ExactTime(t1).To(t2)
	})
}

func TestExactTimeFrom(t *testing.T) {
	tests := []timeTest{
		{"2020-01-01T22:05:00+00:00", "2021-02-01T02:50:22+00:00", "1 year, 30 days, 4 hours, 45 minutes and 22 seconds ago"},
		{"2016-01-01T00:00:00+00:00", "2018-05-01T00:00:00+00:00", "2 years and 4 months ago"},
	}

	runTimeTest(t, tests, func(t1 time.Time, t2 time.Time) string {
		return ExactTime(t1).From(t2)
	})
}

func parseTime(t *testing.T, value string) time.Time {
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Errorf("could not parse time '%s'", value)
	}

	return parsed
}
