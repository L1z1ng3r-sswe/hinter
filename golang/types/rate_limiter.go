package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu         sync.Mutex
	lastTime   time.Time
	limit      time.Duration
	reqsPerSec int // maximum requests per second
	allows     int
}

// one second / reqsLimit
func NewRateLimiter(reqsLimit int, limit time.Duration) *RateLimiter {
	return nil
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.lastTime) >= r.limit {
		r.lastTime = now
		r.allows = r.reqsPerSec
	}

	if r.allows > 0 {
		r.allows--
		return true
	}

	return false
}

type RateLimiterManager struct {
	mu         sync.Mutex
	users      map[string]*RateLimiter
	limit      time.Duration
	reqsPerSec int // equals to RateLimiter.reqsPerSec
}

func NewRateLimiterManager(reqsPerSec int, limit time.Duration) *RateLimiterManager {
	return &RateLimiterManager{
		users:      make(map[string]*RateLimiter),
		limit:      limit,
		reqsPerSec: reqsPerSec,
	}
}

func (rm *RateLimiterManager) GetLimiter(user string) *RateLimiter {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if limiter, ok := rm.users[user]; ok {
		return limiter
	}

	rm.users[user] = NewRateLimiter(rm.reqsPerSec, rm.limit)
	return rm.users[user]
}
