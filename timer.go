package main

import "time"

type timer struct {
	number int64
	unit   string
}

func (t *timer) Duration() time.Duration {
	switch t.unit {
	case "min", "minutes":
		return time.Duration(t.number) * time.Minute
	case "nsec", "nanoseconds", "nanosec":
		return time.Duration(t.number) * time.Nanosecond
	case "usec", "microseconds", "microsec": // micro seconds
		return time.Duration(t.number) * time.Nanosecond * time.Duration(1000)
	case "msec", "milliseconds", "millisec":
		return time.Duration(t.number) * time.Nanosecond * time.Duration(1000000)
	case "hour", "hours":
		return time.Duration(t.number) * time.Hour
	case "sec", "seconds":
		return time.Duration(t.number) * time.Second
	default:
		return time.Duration(t.number) * time.Second
	}
}
