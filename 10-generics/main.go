// Topic 10: Generics (type parameters), Go 1.18+.
//
// Run: go run ./10-generics
package main

import (
	"cmp"
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	withoutGenerics()
	genericFunctions()
	customConstraints()
	anyAndComparable()
	mapFilterReduce()
	genericTypes()
	genericJSON()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// ---------- the problem ----------

// Without generics you'd need one function per type...
func sumInts(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

func sumFloat64s(s []float64) float64 {
	total := 0.0
	for _, v := range s {
		total += v
	}
	return total
}

func withoutGenerics() {
	section("The problem: duplicated code per type")
	fmt.Println("sumInts:", sumInts([]int{1, 2, 3}), "| sumFloat64s:", sumFloat64s([]float64{1.5, 2.5}))
}

// ---------- generic functions ----------

// T is a TYPE PARAMETER; "int | float32 | float64" is its CONSTRAINT.
func sumSlice[T int | float32 | float64](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

func genericFunctions() {
	section("One generic function for many types")
	fmt.Println("sumSlice[int]:", sumSlice[int]([]int{1, 2, 3}))    // explicit type argument
	fmt.Println("sumSlice float32:", sumSlice([]float32{1.5, 2.5})) // inferred
	fmt.Println("sumSlice float64:", sumSlice([]float64{0.25, 0.75}))
	// sumSlice([]string{"a"}) // compile error: string does not satisfy int | float32 | float64
}

// ---------- reusable constraints ----------

// Number is a constraint interface. The ~ means "this type OR any type
// whose underlying type is this", e.g. type Celsius float64.
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type Celsius float64

func average[T Number](s []T) T {
	if len(s) == 0 {
		return 0
	}
	var total T
	for _, v := range s {
		total += v
	}
	return total / T(len(s))
}

// cmp.Ordered (standard library) covers every type that supports < <= > >=.
func maxOf[T cmp.Ordered](values ...T) T {
	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func customConstraints() {
	section("Constraint interfaces, ~ and cmp.Ordered")
	temps := []Celsius{20.5, 22, 25.5}
	fmt.Printf("average temp: %.2f (%T)\n", average(temps), average(temps))
	fmt.Println("maxOf ints:", maxOf(3, 9, 4), "| maxOf strings:", maxOf("apple", "pear", "fig"))
}

// ---------- any and comparable ----------

// any allows every type.
func isEmpty[T any](s []T) bool {
	return len(s) == 0
}

// comparable allows types that support == and != (needed for map keys too).
func indexOf[T comparable](s []T, target T) int {
	for i, v := range s {
		if v == target {
			return i
		}
	}
	return -1
}

func anyAndComparable() {
	section("any and comparable")
	fmt.Println("isEmpty([]string{}):", isEmpty([]string{}), "| isEmpty([]int{1}):", isEmpty([]int{1}))
	fmt.Println("indexOf b:", indexOf([]string{"a", "b", "c"}, "b"), "| indexOf 9:", indexOf([]int{1, 2}, 9))
}

// ---------- multiple type parameters ----------

func mapSlice[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func reduce[T, A any](s []T, initial A, f func(A, T) A) A {
	acc := initial
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func mapFilterReduce() {
	section("Map / Filter / Reduce with two type parameters")
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	labels := mapSlice(evens, func(n int) string { return fmt.Sprintf("#%d", n) })
	total := reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("evens:", evens, "| labels:", labels, "| total:", total)
	product := reduce(nums, 1, func(acc, n int) int { return acc * n })
	fmt.Println("product:", product)
}

// ---------- generic types ----------

// Stack works for any element type.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, nil
}

// Pair holds two values of possibly different types.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// A generic struct can restrict which types it accepts.
type GasEngine struct{ Gallons float64 }
type ElectricEngine struct{ KWh float64 }

type Car[E GasEngine | ElectricEngine] struct {
	Make   string
	Engine E
}

func genericTypes() {
	section("Generic types: Stack[T], Pair[K, V], Car[E]")
	var s Stack[string]
	s.Push("a")
	s.Push("b")
	top, _ := s.Pop()
	fmt.Println("popped:", top)
	s.Pop()
	_, err := s.Pop()
	fmt.Println("pop on empty:", err)

	p := Pair[string, int]{Key: "age", Value: 30}
	fmt.Printf("pair: %+v\n", p)

	gasCar := Car[GasEngine]{Make: "Honda", Engine: GasEngine{Gallons: 12}}
	eCar := Car[ElectricEngine]{Make: "Tesla", Engine: ElectricEngine{KWh: 57}}
	fmt.Printf("%+v\n%+v\n", gasCar, eCar)
}

// ---------- practical example ----------

type ContactInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PurchaseInfo struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

// parseJSON decodes a JSON array into a slice of whichever type you ask for.
func parseJSON[T ContactInfo | PurchaseInfo](data string) ([]T, error) {
	var out []T
	err := json.Unmarshal([]byte(data), &out)
	return out, err
}

func genericJSON() {
	section("Practical: one JSON loader for several struct types")
	contacts, _ := parseJSON[ContactInfo](`[{"name":"Alex","email":"alex@example.com"}]`)
	purchases, _ := parseJSON[PurchaseInfo](`[{"name":"Keyboard","price":49.99,"amount":2}]`)
	fmt.Printf("contacts:  %+v\n", contacts)
	fmt.Printf("purchases: %+v\n", purchases)
}
