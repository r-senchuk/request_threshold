package main

import (
	"testing"
	"time"
)

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
