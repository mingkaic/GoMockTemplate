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

// NewServerIntervals ... creates a NewServer with input workInterval and lifeTime
func NewServerIntervals(work func(), workInterval, lifeTime time.Duration) *MockServer {
	return &MockServer{
		Work:         work,
		WorkInterval: workInterval,
		LifeTime:     lifeTime,
	}
}

// Serve ... serves mock server
func (server *MockServer) Serve() {
	exit := make(chan struct{})
	go func() {
		select {
		case <-time.After(server.LifeTime):
			close(exit)
		}
	}()
	WasteTime(server.Work, server.WorkInterval, exit)
}
