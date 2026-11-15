# 09 · Channels

> **Definition:** A **channel** is a typed pipe that goroutines use to **send and receive values** safely. It both transfers data and synchronizes the goroutines.
> *"Don't communicate by sharing memory; share memory by communicating."*

```bash
go run ./09-channels
```

## Key concepts

| Syntax | Meaning |
|---|---|
| `ch := make(chan int)` | unbuffered: a send blocks until a receiver is ready |
| `ch := make(chan int, 3)` | buffered: a send blocks only when the buffer is full |
| `ch <- v` | send |
| `v := <-ch` | receive |
| `v, ok := <-ch` | `ok == false` if the channel is closed and empty |
| `close(ch)` | no more values will be sent (only the sender should close) |
| `for v := range ch` | receive until the channel is closed |
| `chan<- T` / `<-chan T` | send-only / receive-only channel types |
| `select` | wait on several channel operations at once |

## Examples

### Unbuffered channel

```go
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch
```

```text
received: 42
```

> Sending and receiving in the **same** goroutine on an unbuffered channel deadlocks:
> `fatal error: all goroutines are asleep - deadlock!`

### Buffered channel

```go
ch := make(chan string, 2)
ch <- "first"   // doesn't block
ch <- "second"
fmt.Println(<-ch, <-ch)
```

```text
len: 2 cap: 2
first second
```

### close() and range

```go
func produce(n int, ch chan int) {
	defer close(ch)
	for i := range n { ch <- i * i }
}

go produce(5, ch)
for v := range ch { fmt.Print(v, " ") }
v, ok := <-ch // after close
```

```text
0 1 4 9 16
after close: 0 ok = false
```

### Directional channels

```go
func sender(out chan<- int) { ... }              // can only send
func receiver(in <-chan int, done chan<- int) { ... } // can only receive from in
```

### select, default and timeouts

```go
select {
case msg := <-fast:
	fmt.Println("received", msg)
case msg := <-slow:
	fmt.Println("received", msg)
}

select {
case r := <-result:
	fmt.Println(r)
case <-time.After(50 * time.Millisecond):
	fmt.Println("timed out after 50ms")
}
```

```text
received fast
received slow
nothing ready, moving on (default case)
timed out after 50ms
```

### Worker pool

```go
jobs := make(chan int, 10)
results := make(chan int, 10)
for range 3 {
	wg.Go(func() {
		for j := range jobs { results <- j * 10 }
	})
}
for j := 1; j <= 6; j++ { jobs <- j }
close(jobs)
wg.Wait()
close(results)
```

```text
sum of results: 210 (10+20+...+60 = 210)
```

### Real-world style: the first deal wins

Three goroutines watch shop prices. `select` reacts to whichever one finds a deal first, with a timeout as a fallback.

```text
Deal found: ShopB has a laptop for $790.00
```

## Gotchas

- Sending to a **closed** channel panics. Receiving from one returns the zero value.
- A `nil` channel blocks forever.
- Use channels to pass ownership of data. Use a mutex to protect shared state. Both approaches are fine.

## Try it yourself

1. Build a pipeline: `generate → square → print`, connecting each stage with a channel.
2. Add a `quit` channel to `dealFinder` to cancel the remaining shops early.

---
[← Goroutines](../08-goroutines/) · [🏠 Index](../README.md) · [Next: Generics →](../10-generics/)
