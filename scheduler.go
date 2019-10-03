package scheduler

import (
	"sync"
	"time"
)

type SchedulerInterface interface {
	Execute()
}

type Scheduler struct {
	si        SchedulerInterface
	startTime time.Time
	location  *time.Location
	config    *Config
	mu        sync.Mutex
	isLock    bool
}

func NewScheduler(si SchedulerInterface, opts ...Option) (*Scheduler, error) {
	//default
	config := &Config{
		start:      defaultStart,
		interval:   defaultInterval,
		concurrent: defaultConcurrent,
	}
	for _, opt := range opts {
		opt(config)
	}
	//location, err := time.LoadLocation("Asia/Shanghai")
	location, err := time.LoadLocation("Local")
	if err != nil {
		return nil, err
	}
	t, err := time.ParseInLocation("15:04:05", config.start, location)
	if err != nil {
		return nil, err
	}
	return &Scheduler{
		si:        si,
		startTime: t,
		location:  location,
		config:    config,
	}, nil
}

func (s *Scheduler) Run() {
	go s.start()
}

func (s *Scheduler) checkGoroutineEnd() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.config.concurrent && s.isLock {
		return false
	}
	return true
}

func (s *Scheduler) lock() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isLock = true
}

func (s *Scheduler) unlock() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.isLock = false
}

func (s *Scheduler) start() {
	now := time.Now()
	// Start time.
	t := time.Date(now.Year(), now.Month(), now.Day(), s.startTime.Hour(), s.startTime.Minute(), s.startTime.Second(), 0, s.location)
	if now.After(t) {
		t = t.Add((now.Sub(t)/s.config.interval + 1) * s.config.interval)
	}
	time.Sleep(t.Sub(now)) //Sleep until the start of the next execution
	go func() {
		s.lock()
		s.si.Execute() //first execute
		s.unlock()
	}()
	ticker := time.NewTicker(s.config.interval)
	for _ = range ticker.C {
		go func() {
			if s.checkGoroutineEnd() {
				s.lock()
				s.si.Execute()
				s.unlock()
			}
		}()
	}
}
