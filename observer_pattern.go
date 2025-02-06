package main

import (
	"fmt"
	"sync"
)

// All observers must implement this interface
type Observer interface {
	Update(stock string, price float64)
}

// Subject being observed
type Stock struct {
	name      string
	price     float64
	observers map[Observer]struct{}
	mu        sync.Mutex // Ensures thread safety
}

// Constructor
func NewStock(name string, price float64) *Stock {
	return &Stock{
		name:      name,
		price:     price,
		observers: make(map[Observer]struct{}),
	}
}

func (s *Stock) AddObserver(ob Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers[ob] = struct{}{}

}

func (s *Stock) RemoveObserver(ob Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.observers, ob)
}

func (s *Stock) NotifyObservers() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for ob := range s.observers {
		ob.Update(s.name, s.price)
	}
}

func (s *Stock) SetPrice(price float64) {
	fmt.Printf("\n[Market Update] %s price updated to $%.2f\n", s.name, price)
	s.price = price
	s.NotifyObservers()
}

// Investor Observer
type Investor struct {
	name string
}

func (i *Investor) Update(stock string, price float64) {
	fmt.Printf("[Investor %s] Alert! %s is now $%.2f\n", i.name, stock, price)
}

func main() {
	tesla := NewStock("Tesla", 100.00)

	alice := &Investor{name: "alice"}
	bob := &Investor{name: "bob"}

	tesla.AddObserver(alice)
	tesla.AddObserver(bob)

	// bob and alice recieves the updates
	tesla.SetPrice(920.50)
	tesla.SetPrice(550.60)

	tesla.RemoveObserver(alice)

	// bob only recieves updates
	tesla.SetPrice(1120.00)

}
