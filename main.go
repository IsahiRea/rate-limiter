package main

import (
	"time"

	"github.com/IsahiRea/rate-limiter/concurrency"
)

func main() {
	rl := concurrency.NewRateLimiter(10, 1*time.Second)
	_ = rl
}
