# 10 · Generics

> **Definition:** **Generics** (Go 1.18+) let you write functions and types that work with many types while staying type-safe. You declare **type parameters** in square brackets, `[T any]`, and restrict them with **constraints** (interfaces that list the allowed types or required methods).

```bash
go run ./10-generics
```

## The problem generics solve

```go
func sumInts(s []int) int { ... }
func sumFloat64s(s []float64) float64 { ... } // the same code, duplicated
```

## Generic functions

```go
func sumSlice[T int | float32 | float64](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

sumSlice[int]([]int{1, 2, 3})     // explicit type argument
sumSlice([]float32{1.5, 2.5})     // type inferred
// sumSlice([]string{"a"})        // ❌ compile error
```

```text
sumSlice[int]: 6
sumSlice float32: 4
sumSlice float64: 1
```

## Constraints

```go
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 // ~ includes custom types like `type Celsius float64`
}

func average[T Number](s []T) T { ... }
func maxOf[T cmp.Ordered](values ...T) T { ... } // standard library constraint
```

```text
average temp: 22.67 (main.Celsius)
maxOf ints: 9 | maxOf strings: pear
```

| Constraint | Allows |
|---|---|
| `any` | every type |
| `comparable` | types usable with `==` / `!=` (and as map keys) |
| `cmp.Ordered` | types usable with `<` `>` |
| `int \| string` | a union of specific types |
| `~float64` | `float64` and any type whose underlying type is `float64` |

## Multiple type parameters: Map / Filter / Reduce

```go
func mapSlice[T, U any](s []T, f func(T) U) []U
func filter[T any](s []T, keep func(T) bool) []T
func reduce[T, A any](s []T, initial A, f func(A, T) A) A
```

```text
evens: [2 4 6] | labels: [#2 #4 #6] | total: 21
product: 720
```

## Generic types

```go
type Stack[T any] struct{ items []T }
func (s *Stack[T]) Push(v T) { ... }
func (s *Stack[T]) Pop() (T, error) { ... }

type Pair[K comparable, V any] struct { Key K; Value V }

type Car[E GasEngine | ElectricEngine] struct {
	Make   string
	Engine E
}
```

```text
popped: b
pop on empty: stack is empty
pair: {Key:age Value:30}
{Make:Honda Engine:{Gallons:12}}
{Make:Tesla Engine:{KWh:57}}
```

## Practical: one JSON loader for several types

```go
func parseJSON[T ContactInfo | PurchaseInfo](data string) ([]T, error) {
	var out []T
	err := json.Unmarshal([]byte(data), &out)
	return out, err
}
contacts, _ := parseJSON[ContactInfo](`[{"name":"Alex","email":"alex@example.com"}]`)
```

```text
contacts:  [{Name:Alex Email:alex@example.com}]
purchases: [{Name:Keyboard Price:49.99 Amount:2}]
```

## Gotchas

- Methods **cannot** have their own type parameters. Only functions and types can.
- Don't over-generalize. If an interface with methods does the job, use it.
- The standard `slices`, `maps` and `cmp` packages already provide many generic helpers.

## Try it yourself

1. Write a generic `Queue[T]` with `Enqueue` and `Dequeue`.
2. Write `keys[K comparable, V any](m map[K]V) []K`.

---
[← Channels](../09-channels/) · [🏠 Index](../README.md) · [Next: Building an API →](../11-building-an-api/)
