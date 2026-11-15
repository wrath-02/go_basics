# 14 · Testing *(bonus)*

> **Definition:** Go has testing built in. Tests live in `*_test.go` files next to the code, and `go test` finds and runs every function named `TestXxx(t *testing.T)`. The same tool runs **benchmarks** (`BenchmarkXxx`) and **examples** (`ExampleXxx`, whose printed output is checked against an `// Output:` comment).

This folder is a **library** package (`calc`), not a program, so you test it instead of running it:

```bash
go test -v ./14-testing             # verbose
go test -cover ./14-testing         # coverage %
go test -run Palindrome ./14-testing # only tests matching a regex
go test -bench=. ./14-testing       # benchmarks
```

## Files

| File | Contains |
|---|---|
| [calc.go](calc.go) | the code under test: `Add`, `Divide`, `IsPalindrome`, `Fib` |
| [calc_test.go](calc_test.go) | unit tests, a table-driven test, subtests, a benchmark |
| [example_test.go](example_test.go) | runnable examples that also appear in the docs |

## Examples

### Simple test

```go
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", got)
	}
}
```

### Table-driven test with subtests (the idiomatic style)

```go
tests := []struct {
	name  string
	input string
	want  bool
}{
	{"empty", "", true},
	{"mixed case", "Level", true},
	{"unicode", "été", true},
	{"not a palindrome", "gopher", false},
}
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		if got := IsPalindrome(tt.input); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	})
}
```

### Example test

```go
func ExampleAdd() {
	fmt.Println(calc.Add(2, 3))
	// Output: 5
}
```

### Output

```text
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestIsPalindrome
--- PASS: TestIsPalindrome (0.00s)
    --- PASS: TestIsPalindrome/empty (0.00s)
    --- PASS: TestIsPalindrome/single_letter (0.00s)
    --- PASS: TestIsPalindrome/simple (0.00s)
    --- PASS: TestIsPalindrome/mixed_case (0.00s)
    --- PASS: TestIsPalindrome/unicode (0.00s)
    --- PASS: TestIsPalindrome/not_a_palindrome (0.00s)
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
    --- PASS: TestDivide/ok (0.00s)
    --- PASS: TestDivide/by_zero (0.00s)
=== RUN   TestFib
--- PASS: TestFib (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
=== RUN   ExampleDivide
--- PASS: ExampleDivide (0.00s)
=== RUN   ExampleIsPalindrome
--- PASS: ExampleIsPalindrome (0.00s)
PASS
```

## t.Error vs t.Fatal

| Method | Behaviour |
|---|---|
| `t.Error` / `t.Errorf` | mark the test as failed and **keep going** |
| `t.Fatal` / `t.Fatalf` | mark it as failed and **stop this test now** |
| `t.Run(name, fn)` | run a named subtest |
| `t.Helper()` | report failures at the caller's line (for helper functions) |

> HTTP handlers are tested with `net/http/httptest`. See [11-building-an-api/internal/handlers/handlers_test.go](../11-building-an-api/internal/handlers/handlers_test.go).

---
[← Defer, Panic & Recover](../13-defer-panic-recover/) · [🏠 Index](../README.md)
