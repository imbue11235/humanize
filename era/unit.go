package era

import "time"

const (
	second = time.Second
	minute = time.Minute
	hour   = time.Hour
	day    = 24 * hour
	week   = 7 * day
	month  = 365 * day / 12
	year   = 12 * month
	decade = 10 * year
	long   = 5 * decade
)
