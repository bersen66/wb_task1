package questions

import (
	"runtime"
	"time"
)

// Реализовать собственную функцию sleep.

func Sleep(t time.Duration) {
	start := time.Now()

	for time.Since(start) < t {
		runtime.Gosched()
	}
}
