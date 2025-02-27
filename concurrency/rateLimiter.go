package concurrency

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests    map[string]int
	maxRequests int
	window      time.Duration
	mu          sync.Mutex
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests:    make(map[string]int),
		maxRequests: maxRequests,
		window:      window,
	}
}

func (rl *RateLimiter) reset(key string, currentTime int64) {
	time.Sleep(rl.window)
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.requests, key)
}
