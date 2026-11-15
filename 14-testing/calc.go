// Package calc is a tiny library used to demonstrate Go testing.
//
// Run tests:      go test -v ./14-testing
// With coverage:  go test -cover ./14-testing
// Benchmarks:     go test -bench=. ./14-testing
package calc

import (
	"errors"
	"strings"
)

var ErrDivideByZero = errors.New("divide by zero")

func Add(a, b int) int { return a + b }

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case. It works on runes so "été" is handled correctly.
func IsPalindrome(s string) bool {
	r := []rune(strings.ToLower(s))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// Fib returns the n-th Fibonacci number.
func Fib(n int) int {
	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}
	return a
}
