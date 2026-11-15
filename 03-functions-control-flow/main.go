// Topic 03: Functions and control structures (if/else, switch).
//
// Run: go run ./03-functions-control-flow
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	basicFunctions()
	multipleReturnsAndErrors()
	ifElse()
	switches()
	variadicFunctions()
	functionValuesAndClosures()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// ---------- functions ----------

// Parameters of the same type can share the type: (a, b int).
func add(a, b int) int {
	return a + b
}

func greet(name string) string {
	return "Hello, " + name + "!"
}

func basicFunctions() {
	section("Basic functions")
	fmt.Println("add(2, 3) =", add(2, 3))
	fmt.Println(greet("Gopher"))
}

// ---------- multiple returns and errors ----------

// ErrDivideByZero is a "sentinel" error callers can compare against.
var ErrDivideByZero = errors.New("cannot divide by zero")

// intDivision returns the quotient, the remainder and an error.
// By convention the error is the LAST return value.
func intDivision(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, ErrDivideByZero
	}
	return numerator / denominator, numerator % denominator, nil
}

// Named return values: they are declared up front and returned by a bare "return".
func rectangle(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = 2 * (w + h)
	return
}

func multipleReturnsAndErrors() {
	section("Multiple return values and errors")
	for _, d := range []int{3, 0} {
		q, r, err := intDivision(11, d)
		if err != nil {
			fmt.Printf("11 / %d -> error: %v\n", d, err)
			continue
		}
		fmt.Printf("11 / %d -> quotient %d, remainder %d\n", d, q, r)
	}

	area, perimeter := rectangle(3, 4)
	fmt.Println("rectangle(3, 4) -> area:", area, "perimeter:", perimeter)

	// Use _ (blank identifier) to ignore a value you don't need.
	_, perimeterOnly := rectangle(1, 1)
	fmt.Println("perimeter only:", perimeterOnly)
}

// ---------- if / else ----------

func ifElse() {
	section("if / else if / else")
	for _, score := range []int{95, 72, 40} {
		if score >= 90 {
			fmt.Println(score, "-> A")
		} else if score >= 70 {
			fmt.Println(score, "-> B")
		} else {
			fmt.Println(score, "-> F")
		}
	}

	// if with a short statement: the variable is scoped to the if/else block.
	if q, _, err := intDivision(10, 5); err == nil && q == 2 {
		fmt.Println("10 / 5 is 2 (err == nil && q == 2)")
	}
}

// ---------- switch ----------

func switches() {
	section("switch")

	// switch on a value. No "break" needed: cases don't fall through by default.
	day := "sat"
	switch day {
	case "sat", "sun":
		fmt.Println(day, "-> weekend")
	default:
		fmt.Println(day, "-> weekday")
	}

	// switch with no condition is a cleaner if/else chain.
	temp := 31
	switch {
	case temp > 30:
		fmt.Println(temp, "°C -> hot")
	case temp > 15:
		fmt.Println(temp, "°C -> pleasant")
	default:
		fmt.Println(temp, "°C -> cold")
	}

	// fallthrough explicitly continues into the next case.
	switch n := 1; n {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2 (reached via fallthrough)")
	case 3:
		fmt.Println("case 3 (not reached)")
	}
}

// ---------- variadic ----------

// sum accepts any number of ints. Inside the function, nums is a []int.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func variadicFunctions() {
	section("Variadic functions")
	fmt.Println("sum() =", sum())
	fmt.Println("sum(1, 2, 3) =", sum(1, 2, 3))
	values := []int{10, 20, 30}
	fmt.Println("sum(values...) =", sum(values...)) // spread a slice
}

// ---------- function values and closures ----------

// apply takes a function as a parameter.
func apply(words []string, transform func(string) string) []string {
	out := make([]string, 0, len(words))
	for _, w := range words {
		out = append(out, transform(w))
	}
	return out
}

// counter returns a closure that remembers "count" between calls.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func functionValuesAndClosures() {
	section("Functions as values and closures")
	words := []string{"go", "is", "fun"}
	fmt.Println(apply(words, strings.ToUpper))

	// anonymous function passed inline
	fmt.Println(apply(words, func(s string) string { return s + "!" }))

	next := counter()
	fmt.Println("counter:", next(), next(), next())

	// immediately invoked anonymous function
	result := func(x int) int { return x * x }(9)
	fmt.Println("square of 9:", result)
}
