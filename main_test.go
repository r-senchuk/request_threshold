package main

import (
	"testing"
	"time"
)

//Imagine we are building an application that is used by many different customers.
//We want to avoid one customer being able to overload the system by sending too many requests,
//so we enforce a per-customer rate limit. The rate limit is defined as:

//“Each customer can make X requests per Y seconds”

//Assuming that customer ID is extracted somehow from the request, implement the following function.

// Perform rate limiting logic for provided customer ID. Return true if the
// request is allowed, and false if it is not.
func TestRateLimit(t *testing.T) {

	// 2 requests per 2 secs
	tests := []struct {
		Name string
		Id   int
		Wait int
		Want bool
	}{
		{Name: "Example 1", Id: 2, Wait: 0, Want: true},
		{Name: "Example 2", Id: 2, Wait: 0, Want: true},
		{Name: "Example 3", Id: 2, Wait: 0, Want: true},
	}

	for _, tt := range tests {
		nr := NewRate(2, time.Second*20)
		t.Run(tt.Name, func(t *testing.T) {
			time.Sleep(time.Duration(tt.Wait))
			got := nr.RateLimit(tt.Id)
			if got != tt.Want {
				t.Errorf("rateLimit(%v) responded with %v but should %v", tt.Id, got, tt.Want)
			}
		})
	}
}
