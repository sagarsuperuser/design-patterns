package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct{}

var instance *Singleton
var once sync.Once

// GetInstance ensures only one instance exists
// Thread safe, prevents multiple instances
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating Intance...")
		instance = &Singleton{}
	})
	return instance
}

// func main() {
// 	s1 := GetInstance()
// 	s2 := GetInstance()
// 	fmt.Println(s1 == s2)
// }

// Test for single instance
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = GetInstance()
		}()
	}

	wg.Wait()
}
