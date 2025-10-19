package utils

import (
	"time"
)

// SleepMilliseconds pauses the current goroutine for the specified number of milliseconds
func SleepMilliseconds(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// CurrentTimeMillis returns the current time in milliseconds since the Unix epoch
func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
