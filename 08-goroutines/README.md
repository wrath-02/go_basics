# 08 · Goroutines

> **Definition:** A **goroutine** is a lightweight thread managed by the Go runtime. Starting one is as simple as putting `go` before a function call. Goroutines run **concurrently** (their execution overlaps), and on a multi-core CPU they also run in **parallel**. They are cheap: thousands or even millions can run at once.

```bash
go run ./08-goroutines
go run -race ./08-goroutines   # the race detector catches unsafe shared memory access
```

> Output order and timings of concurrent code **vary between runs**.

## Key tools

| Tool | Purpose |
|---|---|
| `go f()` | run `f` in a new goroutine |
| `sync.WaitGroup` | wait for a group of goroutines to finish (`Add`, `Done`, `Wait`, or `Go` since Go 1.25) |
| `sync.Mutex` | only one goroutine at a time in a critical section (`Lock` / `Unlock`) |
| `sync.RWMutex` | many readers **or** one writer (`RLock` / `RUnlock`, `Lock` / `Unlock`) |
| `sync/atomic` | lock-free counters and flags |

## Examples

### Sequential vs concurrent

```go
func dbCall(i int) string {
	time.Sleep(100 * time.Millisecond) // simulate a slow query
	return dbData[i]
}

// Sequential: 5 × 100ms
for i := range dbData { dbCall(i) }

// Concurrent: all 5 calls overlap
var wg sync.WaitGroup
for i := range dbData {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  got", dbCall(i))
	}()
}
wg.Wait()
```

```text
5 calls took ~500ms
  got id4
  got id1
  got id3
  got id5
  got id2
5 calls took ~100ms (order above is random)
```

Since Go 1.25, `wg.Go(func(){...})` does `Add(1)`, `go` and `Done()` for you.

### Race condition → Mutex

Appending to a shared slice from many goroutines is a **data race**: writes can overwrite each other.

```go
var (
	mu      sync.Mutex
	results []string
)
wg.Go(func() {
	value := dbCall(i) // slow work happens OUTSIDE the lock
	mu.Lock()
	results = append(results, value)
	mu.Unlock()
})
```

```text
results (sorted): [id1 id2 id3 id4 id5] count: 5
```

### RWMutex for read-heavy data

```go
func (c *Cache) Get(key string) (int, bool) {
	c.mu.RLock()            // many readers at once
	defer c.mu.RUnlock()
	...
}
func (c *Cache) Set(key string, v int) {
	c.mu.Lock()             // writer gets exclusive access
	defer c.mu.Unlock()
	...
}
```

### Atomic counter

```go
var counter atomic.Int64
for range 1000 {
	wg.Go(func() { counter.Add(1) })
}
```

```text
counter = 1000 (always 1000, no lock needed)
```

## Gotchas

- `main` **doesn't wait** for goroutines. When `main` returns, they're killed. Use a `WaitGroup` or channels.
- Never copy a `sync.Mutex` or `WaitGroup` after first use. Pass a pointer instead.
- Concurrency ≠ automatic speed-up. CPU-bound work only scales up to the number of CPU cores.
- Since Go 1.22, every loop iteration gets its own variable, so capturing `i` in a goroutine is safe.

## Try it yourself

1. Remove the mutex from `raceAndMutex` and run with `-race`. Read the report.
2. Fetch 10 URLs concurrently with `net/http` and print each status code.

---
[← Pointers](../07-pointers/) · [🏠 Index](../README.md) · [Next: Channels →](../09-channels/)
