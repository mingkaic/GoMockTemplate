package gomock

import (
	"time"
)

// WasteTime ... do something every time interval until exit is closed
func WasteTime(closure func(), interval time.Duration, exit chan struct{}) {
	var timeout <-chan time.Time
	for {
		timeout = time.After(interval)
		select {
		case <-exit:
			return
		case <-timeout:
			closure()
		}
	}
}
