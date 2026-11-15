// Topic 13 (bonus): defer, panic and recover.
//
// Run: go run ./13-defer-panic-recover
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	deferBasics()
	deferLIFO()
	deferArgumentsEvaluatedEarly()
	deferForCleanup()
	panicAndRecover()
	fmt.Println("\nmain finished normally")
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func deferBasics() {
	section("defer runs when the surrounding function returns")
	defer fmt.Println("3. deferred: printed last")
	fmt.Println("1. first")
	fmt.Println("2. second")
}

func deferLIFO() {
	section("Multiple defers run in reverse order (LIFO)")
	for i := range 3 {
		defer fmt.Println("deferred", i)
	}
	fmt.Println("loop done")
}

func deferArgumentsEvaluatedEarly() {
	section("Arguments are evaluated when defer is declared")
	x := 1
	defer fmt.Println("deferred sees x =", x) // x is captured as 1 right now
	x = 100
	fmt.Println("x is now", x)
}

// The most common use: release a resource no matter how the function exits.
func writeTempFile() (path string, err error) {
	f, err := os.CreateTemp("", "go-basics-*.txt")
	if err != nil {
		return "", err
	}
	defer f.Close() // always closed, even if a later step fails

	if _, err := f.WriteString("hello from defer\n"); err != nil {
		return "", err
	}
	return f.Name(), nil
}

func deferForCleanup() {
	section("defer for cleanup (files, locks, connections)")
	path, err := writeTempFile()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer os.Remove(path)
	data, _ := os.ReadFile(path)
	fmt.Printf("wrote and read back a temp file: %q\n", string(data))
}

// safeDivide turns a panic into a normal error using recover.
// recover only works inside a deferred function.
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v", r)
		}
	}()
	return a / b, nil // b == 0 causes a runtime panic
}

func mustPositive(n int) int {
	if n < 0 {
		panic(errors.New("mustPositive: got a negative number"))
	}
	return n
}

func panicAndRecover() {
	section("panic and recover")
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(1, 0))

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered:", r)
			}
		}()
		mustPositive(-1)
		fmt.Println("never printed")
	}()
	// Use panic only for truly unrecoverable programmer errors.
	// For expected failures (bad input, missing file, ...) return an error.
}
