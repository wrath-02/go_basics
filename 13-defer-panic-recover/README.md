# 13 · Defer, Panic & Recover *(bonus)*

> **Definitions:**
> - **defer** schedules a function call to run when the surrounding function **returns**, whether it returns normally or panics. It's mostly used for cleanup.
> - **panic** stops normal execution and starts unwinding the stack, running deferred calls on the way. If nothing recovers, the program crashes.
> - **recover** (only useful inside a deferred function) stops a panic and returns the panic value.

```bash
go run ./13-defer-panic-recover
```

## Examples

### defer runs last

```go
defer fmt.Println("3. deferred: printed last")
fmt.Println("1. first")
fmt.Println("2. second")
```

```text
1. first
2. second
3. deferred: printed last
```

### Multiple defers run LIFO (like a stack)

```go
for i := range 3 {
	defer fmt.Println("deferred", i)
}
fmt.Println("loop done")
```

```text
loop done
deferred 2
deferred 1
deferred 0
```

### Arguments are evaluated immediately

```go
x := 1
defer fmt.Println("deferred sees x =", x) // x is captured as 1 now
x = 100
```

```text
x is now 100
deferred sees x = 1
```

### Cleanup (the #1 use case)

```go
f, err := os.Create(path)
if err != nil { return err }
defer f.Close() // also used for mu.Unlock(), rows.Close(), resp.Body.Close()...
```

### Turning a panic into an error

```go
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	return a / b, nil
}
```

```text
5 <nil>
0 recovered from panic: runtime error: integer divide by zero
recovered: mustPositive: got a negative number

main finished normally
```

## When to panic?

- ❌ Not for expected failures (bad input, missing file, network error). **Return an error** instead.
- ✅ For truly impossible states or programmer bugs, and in `Must...` helpers used at startup, such as `regexp.MustCompile`.
- ✅ `net/http` recovers panics in handlers, so one bad request doesn't crash the server.

---
[← Error Handling](../12-error-handling/) · [🏠 Index](../README.md) · [Next: Testing →](../14-testing/)
