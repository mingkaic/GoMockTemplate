package gomock

import (
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	var counter = 0
	server := NewServerIntervals(func() {
		counter++
	}, time.Duration(1)*time.Second, time.Duration(6)*time.Second)
	server.Serve()
	if counter != 5 {
		t.Errorf("expecting work to be called 5 times, it ran %d times", counter)
	}
}
