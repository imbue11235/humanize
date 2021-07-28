package humantime

import "time"

func absDuration(duration time.Duration) time.Duration {
	if duration < 0 {
		return -duration
	}

	return duration
}