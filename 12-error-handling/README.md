# 12 · Error Handling *(bonus)*

> **Definition:** In Go, errors are **ordinary values** returned from functions, not exceptions. The built-in `error` type is an interface with a single method: `Error() string`. By convention, a function returns its error as the **last** value, and `nil` means success.

```bash
go run ./12-error-handling
```

## Key tools

| Tool | Use |
|---|---|
| `errors.New("msg")` | create a simple error |
| `fmt.Errorf("ctx: %w", err)` | **wrap** an error with more context |
| `errors.Is(err, target)` | is `target` anywhere in the wrap chain? |
| `errors.As(err, &target)` | find an error of a specific **type** in the chain |
| `errors.Unwrap(err)` | get the wrapped error |
| `errors.Join(e1, e2)` | combine several errors into one |

## Examples

### The basic pattern

```go
age, err := parseAge(in)
if err != nil {
	fmt.Println("error:", err)
	return
}
```

```text
age: 30
error: parseAge("abc"): not a number
error: age cannot be negative
```

### Sentinel errors

```go
var (
	ErrNotFound     = errors.New("not found")
	ErrInsufficient = errors.New("insufficient funds")
)
```

```text
withdraw(bob, 10): not found
withdraw(alex, 100): insufficient funds
withdraw(alex, 20): <nil> -> balance 30
```

### Wrapping with `%w` + `errors.Is`

```go
return fmt.Errorf("charge order for %s: %w", user, err)
```

```text
error: charge order for alex: insufficient funds
err == ErrInsufficient?          false
errors.Is(err, ErrInsufficient)? true
errors.Unwrap(err): insufficient funds
os.Open missing file is fs.ErrNotExist? true
```

### Custom error types + `errors.As`

```go
type ValidationError struct{ Field, Issue string }
func (e *ValidationError) Error() string { return e.Field + ": " + e.Issue }

var ve *ValidationError
if errors.As(err, &ve) {
	fmt.Printf("field=%q issue=%q\n", ve.Field, ve.Issue)
}
```

```text
error: signup failed: email: is required
field="email" issue="is required"
```

### `errors.Join`

```text
valid user: <nil>
invalid user:
name is required
must be 18 or older
```

## Best practices

- Handle each error **once**: either log it or return it, not both.
- Add context when returning: `fmt.Errorf("load config: %w", err)`.
- Error strings are lowercase with no trailing punctuation.
- Compare with `errors.Is`, never with `==` or string matching.

---
[← Building an API](../11-building-an-api/) · [🏠 Index](../README.md) · [Next: Defer, Panic & Recover →](../13-defer-panic-recover/)
