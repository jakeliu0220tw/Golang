package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc() increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // lock when entering ...
	c.v[key]++
	c.mu.Unlock() // unlock when leaving ...
}

// Value() return the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func mutexDemo() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("mykey")
	}

	time.Sleep(time.Second)
	fmt.Println("value for the given key is ...", c.Value("mykey"))
}
