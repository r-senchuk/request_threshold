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

func NewRate(numberAllowed int, perTime time.Duration) *RateLimiter {
	return &RateLimiter{
		mu:             sync.Mutex{},
		UserRequestMap: map[int]*RequestsData{},
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
	var rd *RequestsData
	var ok bool
	r.mu.Lock()
	defer r.mu.Unlock()
	if rd, ok = r.UserRequestMap[customerId]; !ok {
		rd = &RequestsData{
			Count:   1,
			Initial: time.Now(),
		}
		r.UserRequestMap[customerId] = rd
	}

	return rd
}
