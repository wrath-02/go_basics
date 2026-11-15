// Topic 12 (bonus): Error handling.
//
// Run: go run ./12-error-handling
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func main() {
	basics()
	sentinelErrors()
	wrapping()
	customErrorTypes()
	joiningErrors()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// ---------- basics ----------

// In Go, errors are VALUES returned as the last result, not exceptions.
// The error type is a built-in interface: interface { Error() string }
func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge(%q): not a number", s)
	}
	if age < 0 {
		return 0, errors.New("age cannot be negative")
	}
	return age, nil
}

func basics() {
	section("Errors are values: if err != nil")
	for _, in := range []string{"30", "abc", "-5"} {
		age, err := parseAge(in)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("age:", age)
	}
}

// ---------- sentinel errors ----------

// Sentinel errors are package-level values callers can check for.
var (
	ErrNotFound     = errors.New("not found")
	ErrInsufficient = errors.New("insufficient funds")
)

var balances = map[string]float64{"alex": 50}

func withdraw(user string, amount float64) error {
	bal, ok := balances[user]
	if !ok {
		return ErrNotFound
	}
	if amount > bal {
		return ErrInsufficient
	}
	balances[user] = bal - amount
	return nil
}

func sentinelErrors() {
	section("Sentinel errors")
	fmt.Println("withdraw(bob, 10):", withdraw("bob", 10))
	fmt.Println("withdraw(alex, 100):", withdraw("alex", 100))
	fmt.Println("withdraw(alex, 20):", withdraw("alex", 20), "-> balance", balances["alex"])
}

// ---------- wrapping ----------

// %w wraps an error, adding context while keeping the original inside.
func chargeOrder(user string, amount float64) error {
	if err := withdraw(user, amount); err != nil {
		return fmt.Errorf("charge order for %s: %w", user, err)
	}
	return nil
}

func wrapping() {
	section("Wrapping with %w, checking with errors.Is")
	err := chargeOrder("alex", 1000)
	fmt.Println("error:", err)

	// errors.Is looks through the whole wrap chain. Plain == would fail here.
	fmt.Println("err == ErrInsufficient?         ", err == ErrInsufficient)
	fmt.Println("errors.Is(err, ErrInsufficient)?", errors.Is(err, ErrInsufficient))
	fmt.Println("errors.Unwrap(err):", errors.Unwrap(err))

	// Standard library errors work the same way.
	_, err = os.Open("does-not-exist.txt")
	fmt.Println("os.Open missing file is fs.ErrNotExist?", errors.Is(err, fs.ErrNotExist))
}

// ---------- custom error types ----------

// ValidationError carries extra data about what went wrong.
type ValidationError struct {
	Field string
	Issue string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Issue
}

func validateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("signup failed: %w", &ValidationError{Field: "email", Issue: "is required"})
	}
	return nil
}

func customErrorTypes() {
	section("Custom error types and errors.As")
	err := validateEmail("")
	fmt.Println("error:", err)

	// errors.As finds the first error in the chain of the given type
	// and assigns it, so you can read its fields.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("field=%q issue=%q\n", ve.Field, ve.Issue)
	}
}

// ---------- joining ----------

func validateUser(name string, age int) error {
	var errs []error
	if name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if age < 18 {
		errs = append(errs, errors.New("must be 18 or older"))
	}
	return errors.Join(errs...) // nil if errs is empty
}

func joiningErrors() {
	section("errors.Join: report several errors at once")
	fmt.Println("valid user:", validateUser("Sam", 30))
	fmt.Println("invalid user:")
	fmt.Println(validateUser("", 12))
}
