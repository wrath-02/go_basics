# 03 · Functions & Control Structures

> **Definition:** A **function** is a named, reusable block of code that takes typed parameters and can return **multiple** values. **Control structures** (`if`, `switch`, `for`) decide which code runs and how often.

```bash
go run ./03-functions-control-flow
```

## Key concepts

| Concept | Syntax |
|---|---|
| Function | `func add(a, b int) int { return a + b }` |
| Multiple returns | `func f() (int, error)` |
| Named returns | `func f() (area, perim float64) { ...; return }` |
| Variadic | `func sum(nums ...int) int`, called as `sum(1, 2)` or `sum(slice...)` |
| Function value | `var f func(string) string = strings.ToUpper` |
| Closure | an anonymous function that captures outer variables |
| `if` with init | `if v, err := f(); err != nil { ... }` |
| `switch` | no `break` needed; use `fallthrough` to continue to the next case |

## Examples

### Multiple return values and errors

```go
var ErrDivideByZero = errors.New("cannot divide by zero")

func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, ErrDivideByZero
	}
	return numerator / denominator, numerator % denominator, nil
}

for _, d := range []int{3, 0} {
	q, r, err := intDivision(11, d)
	if err != nil {
		fmt.Printf("11 / %d -> error: %v\n", d, err)
		continue
	}
	fmt.Printf("11 / %d -> quotient %d, remainder %d\n", d, q, r)
}
```

```text
11 / 3 -> quotient 3, remainder 2
11 / 0 -> error: cannot divide by zero
```

### Named return values

```go
func rectangle(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = 2 * (w + h)
	return // "naked" return sends back area and perimeter
}
```

```text
rectangle(3, 4) -> area: 12 perimeter: 14
```

### if / else if / else

```go
if score >= 90 {
	fmt.Println(score, "-> A")
} else if score >= 70 {
	fmt.Println(score, "-> B")
} else {
	fmt.Println(score, "-> F")
}
```

```text
95 -> A
72 -> B
40 -> F
```

### switch

```go
switch day {               // value switch, several values per case
case "sat", "sun":
	fmt.Println(day, "-> weekend")
default:
	fmt.Println(day, "-> weekday")
}

switch {                   // no condition = cleaner if/else chain
case temp > 30:
	fmt.Println(temp, "°C -> hot")
case temp > 15:
	fmt.Println(temp, "°C -> pleasant")
}

switch n := 1; n {
case 1:
	fmt.Println("case 1")
	fallthrough            // explicitly run the next case
case 2:
	fmt.Println("case 2 (reached via fallthrough)")
}
```

```text
sat -> weekend
31 °C -> hot
case 1
case 2 (reached via fallthrough)
```

### Variadic functions

```go
func sum(nums ...int) int { /* nums is a []int */ }

sum()                  // 0
sum(1, 2, 3)           // 6
sum([]int{10, 20, 30}...) // 60
```

### Functions as values and closures

```go
func counter() func() int {
	count := 0
	return func() int { count++; return count }
}

fmt.Println(apply(words, strings.ToUpper))
fmt.Println(apply(words, func(s string) string { return s + "!" }))
next := counter()
fmt.Println("counter:", next(), next(), next())
```

```text
[GO IS FUN]
[go! is! fun!]
counter: 1 2 3
```

## Gotchas

- Go has **no default parameters and no function overloading**.
- There is no ternary operator (`a ? b : c`). Use `if`/`else`.
- Always handle the `error`. Ignoring it with `_` should be a deliberate choice.

## Try it yourself

1. Write `minMax(nums ...int) (min, max int)`.
2. Write a `switch` that prints the season for a month number from 1 to 12.

---
[← Variables & Types](../02-variables-constants-types/) · [🏠 Index](../README.md) · [Next: Arrays, Slices, Maps & Loops →](../04-arrays-slices-maps-loops/)
