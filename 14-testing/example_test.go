package calc_test // external test package: only sees exported names

import (
	"fmt"

	calc "github.com/wrath-02/go_basics/14-testing"
)

// Example functions are documentation AND tests: `go test` checks that
// what they print matches the "// Output:" comment exactly.

func ExampleAdd() {
	fmt.Println(calc.Add(2, 3))
	// Output: 5
}

func ExampleDivide() {
	q, err := calc.Divide(7, 0)
	fmt.Println(q, err)
	// Output: 0 divide by zero
}

func ExampleIsPalindrome() {
	fmt.Println(calc.IsPalindrome("Racecar"), calc.IsPalindrome("Go"))
	// Output: true false
}
