package main

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

func main() {

}
