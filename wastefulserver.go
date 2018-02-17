package gomock

import "time"

type MockServer struct {
	Work         func()
	WorkInterval time.Duration
	LifeTime     time.Duration
}

// NewServer ... creates a mock server that:
// run work every 10 seconds
// dies after 10 minutes
func NewServer(work func()) *MockServer {
	return &MockServer{
		Work:         work,
		WorkInterval: time.Duration(10) * time.Second,
		LifeTime:     time.Duration(10) * time.Minute,
	}
}

// Serve ... serves mock server
func (server *MockServer) Serve() {
	exit := make(chan struct{})
	select {
	case <-time.Tick(server.LifeTime):
		close(exit)
	}
	WasteTime(server.Work, server.WorkInterval, exit)
}
