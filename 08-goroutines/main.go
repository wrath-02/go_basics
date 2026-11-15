// Topic 08: Goroutines, WaitGroup and Mutex.
//
// Run:            go run ./08-goroutines
// Detect races:   go run -race ./08-goroutines
//
// NOTE: concurrent output order and timings vary between runs.
package main

import (
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	sequential()
	concurrentWithWaitGroup()
	raceAndMutex()
	readWriteMutex()
	atomicCounter()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// dbCall simulates a slow database query.
func dbCall(i int) string {
	time.Sleep(100 * time.Millisecond)
	return dbData[i]
}

func sequential() {
	section("Sequential: one call after another")
	start := time.Now()
	for i := range dbData {
		dbCall(i)
	}
	fmt.Printf("5 calls took ~%v\n", time.Since(start).Round(100*time.Millisecond))
}

func concurrentWithWaitGroup() {
	section("Concurrent: goroutines + sync.WaitGroup")
	// "go f()" starts f in a new goroutine (a lightweight thread managed by Go).
	// main does NOT wait for goroutines, so we use a WaitGroup to wait.
	start := time.Now()
	var wg sync.WaitGroup
	for i := range dbData {
		wg.Add(1) // one more goroutine to wait for
		go func() {
			defer wg.Done() // mark this goroutine as finished
			fmt.Println("  got", dbCall(i))
		}()
	}
	wg.Wait() // block until the counter reaches 0
	fmt.Printf("5 calls took ~%v (order above is random)\n", time.Since(start).Round(100*time.Millisecond))

	// Go 1.25+: wg.Go does Add(1) + go + Done() for you.
	var wg2 sync.WaitGroup
	for i := range 3 {
		wg2.Go(func() { _ = dbCall(i) })
	}
	wg2.Wait()
	fmt.Println("wg.Go version finished too")
}

func raceAndMutex() {
	section("Race condition and sync.Mutex")
	// Many goroutines appending to the same slice at once is a DATA RACE:
	// appends can overwrite each other and results get lost.
	// A Mutex lets only one goroutine into the "critical section" at a time.
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		results []string
	)
	for i := range dbData {
		wg.Go(func() {
			value := dbCall(i) // slow work happens OUTSIDE the lock
			mu.Lock()
			results = append(results, value)
			mu.Unlock()
		})
	}
	wg.Wait()
	slices.Sort(results)
	fmt.Println("results (sorted):", results, "count:", len(results))
}

// Cache uses an RWMutex: many readers OR one writer at a time.
type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock() // multiple goroutines may hold a read lock together
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) Set(key string, v int) {
	c.mu.Lock() // exclusive: waits for all readers to finish
	defer c.mu.Unlock()
	c.data[key] = v
}

func readWriteMutex() {
	section("sync.RWMutex: many readers, one writer")
	cache := &Cache{data: map[string]int{}}
	var wg sync.WaitGroup
	for i := range 3 {
		wg.Go(func() { cache.Set(fmt.Sprintf("key%d", i), i*10) })
	}
	wg.Wait()
	for i := range 3 {
		wg.Go(func() { cache.Get(fmt.Sprintf("key%d", i)) })
	}
	wg.Wait()
	v, _ := cache.Get("key2")
	fmt.Println("cache key2 =", v, "| entries:", len(cache.data))
}

func atomicCounter() {
	section("sync/atomic for simple counters")
	var counter atomic.Int64
	var wg sync.WaitGroup
	for range 1000 {
		wg.Go(func() { counter.Add(1) })
	}
	wg.Wait()
	fmt.Println("counter =", counter.Load(), "(always 1000, no lock needed)")
}
