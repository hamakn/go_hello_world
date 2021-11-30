package helper

import "time"

func IsBetweenFunc(startedAt time.Time, endedAt time.Time) func(now time.Time) bool {
	return func(now time.Time) bool {
		return now == startedAt || now == endedAt || (now.After(startedAt) && now.Before(endedAt))
	}
}
