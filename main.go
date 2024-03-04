package main

import "time"

type RequestRate struct {
	ReqCount  int
	FirstTime time.Time
}

type Rate struct {
	UserRequests  map[int]*RequestRate
	NumberAllowed int
	PerTime       time.Duration
}

func NewRate(numberAllowed int, perTime time.Duration) Rate {
	reqMap := make(map[int]*RequestRate)
	return Rate{
		UserRequests:  reqMap,
		NumberAllowed: numberAllowed,
		PerTime:       perTime,
	}
}

func (r *Rate) RateLimit(customerId int) bool {

	if _, ok := r.UserRequests[customerId]; !ok {
		r.UserRequests[customerId] = &RequestRate{
			ReqCount:  1,
			FirstTime: time.Now(),
		}
	}

	timefromLastReq := time.Since(r.UserRequests[customerId].FirstTime)

	if timefromLastReq < r.PerTime {
		if rate, ok := r.UserRequests[customerId]; ok {
			if rate.ReqCount < r.NumberAllowed {
				rate.ReqCount++
				r.UserRequests[customerId] = rate
				return true
			}
		}

		return false
	}

	r.UserRequests[customerId].ReqCount = 1
	r.UserRequests[customerId].FirstTime = time.Now()

	return true
}
