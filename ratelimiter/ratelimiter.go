package ratelimiter

import (
	"container/list"
	"sync"
	"time"
)

// RateLimiter is a struct that represents a rate limiter.
type RateLimiter struct {
	attempts  int // Number of attempts allowed
	timeLimit int // Time limit in milliseconds
	clientMap map[string]*list.List
	mu        sync.Mutex // Mutex to synchronize access to the clientMap
}

// NewRateLimiter creates a new RateLimiter instance with the specified number of attempts and time limit.
func NewRateLimiter(attempts, timeLimit int) *RateLimiter {
	return &RateLimiter{
		attempts:  attempts,
		timeLimit: timeLimit,
		clientMap: make(map[string]*list.List),
	}
}

// IsAllowed checks if a client with the given ID is allowed to perform an action based on the rate limiting rules.
// It returns true if the client is allowed, false otherwise.
func (r *RateLimiter) IsAllowed(clientID string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if val, ok := r.clientMap[clientID]; ok {
		r.removeExpiredTime(val)
		if val.Len() > r.attempts {
			return false
		}
		r.addTimestamp(val)
	} else {
		list := list.New()
		r.addTimestamp(list)
		r.clientMap[clientID] = list
	}
	return true
}

// removeExpiredTime removes expired timestamps from the given list and returns the updated list.
func (r *RateLimiter) removeExpiredTime(val *list.List) {
	curTime := time.Now().UnixMilli()
	// Remove expired time
	for val.Len() > 0 && (curTime-val.Front().Value.(int64)) >= int64(r.timeLimit) {
		val.Remove(val.Front())
	}
}

// addTimestamp adds the current timestamp to the given list.
func (r *RateLimiter) addTimestamp(val *list.List) {
	val.PushBack(time.Now().UnixMilli())
}
