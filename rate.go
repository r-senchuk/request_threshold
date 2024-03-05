package main

import (
	"sync"
	"time"
)

type RequestsData struct {
	Count   int
	Initial time.Time
}

func (rd *RequestsData) Reset() {
	rd.Count = 1
	rd.Initial = time.Now()
}

type RateLimiter struct {
	mu             sync.Mutex
	UserRequestMap map[int]*RequestsData
	MaxRequests    int
	Window         time.Duration
}

func NewRate(numberAllowed int, perTime time.Duration) RateLimiter {
	reqMap := make(map[int]*RequestsData)
	return RateLimiter{
		mu:             sync.Mutex{},
		UserRequestMap: reqMap,
		MaxRequests:    numberAllowed,
		Window:         perTime,
	}
}

func (r *RateLimiter) RateLimit(customerId int) bool {
	requestsData := r.getRequestData(customerId)
	if time.Since(requestsData.Initial) > r.Window {
		requestsData.Reset()
		return true
	}
	if requestsData.Count > r.MaxRequests {
		return false
	}

	requestsData.Count++
	return true
}

func (r *RateLimiter) getRequestData(customerId int) *RequestsData {
	rd, ok := r.UserRequestMap[customerId]

	if !ok {
		r.mu.Lock()
		r.UserRequestMap[customerId] = &RequestsData{
			Count:   1,
			Initial: time.Now(),
		}
		r.mu.Unlock()
	}

	return rd
}
