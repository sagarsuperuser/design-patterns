package main

import (
	"fmt"
	"sync"
	"time"
)

// ApiService defines the interface for API requests.
type ApiService interface {
	FetchData(userID int) string
}

// RealAPIService represents the actual API performing expensive operations.
type RealAPIService struct{}

func (r *RealAPIService) FetchData(userID int) string {
	fmt.Println("Fetching data from real API service....")
	time.Sleep(1 * time.Second) // Simulating API response time
	return fmt.Sprintf("UserData for user %d", userID)
}

// RateLimitingProxy acts as a proxy to limit API requests.
type RateLimitingProxy struct {
	apiService    ApiService
	requests      map[int]int // Tracks requests per user
	mutex         sync.Mutex  // Ensures thread safety
	rateLimit     int         // Max requests per user
	resetInterval time.Duration
}

// NewRateLimitingProxy initializes the proxy
func NewRateLimitingProxy(api ApiService, rateLimit int, resetInterval time.Duration) *RateLimitingProxy {
	proxy := &RateLimitingProxy{
		apiService:    api,
		requests:      make(map[int]int),
		rateLimit:     rateLimit,
		resetInterval: resetInterval,
	}
	// Start a goroutine to reset request counts periodically.
	go func() {
		for {
			time.Sleep(resetInterval)
			proxy.mutex.Lock()
			proxy.requests = make(map[int]int)
			proxy.mutex.Unlock()
			fmt.Println("Rate limit reset.")
		}
	}()
	return proxy
}

// FetchData enforces rate limiting before calling the real API.
func (p *RateLimitingProxy) FetchData(userID int) string {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Check if the user exceeded the rate limit
	if p.requests[userID] >= p.rateLimit {
		return fmt.Sprintf("Rate Limit exceeded for user: %d. Try again later.", userID)
	}

	// Increment request count and forward request to the real service.
	p.requests[userID]++
	return p.apiService.FetchData(userID)
}

func main() {
	realAPI := &RealAPIService{}
	proxy := NewRateLimitingProxy(realAPI, 3, 8*time.Second) // Allow 3 requests per user every 8 seconds
	userID := 1

	// Simulate multiple API requests
	for i := 0; i < 5; i++ {
		fmt.Println(proxy.FetchData(userID))
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\nWaiting for rate limit reset... \n")
	time.Sleep(6 * time.Second)

	// After reset, requests should work again
	fmt.Println(proxy.FetchData(userID))

}
