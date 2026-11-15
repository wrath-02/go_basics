package calc

import (
	"errors"
	"testing"
)

// A test is a function named TestXxx(t *testing.T) in a _test.go file.
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		t.Errorf("Add(2, 3) = %d, want %d", got, want)
	}
}

// Table-driven tests: the idiomatic way to test many cases.
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"single letter", "a", true},
		{"simple", "racecar", true},
		{"mixed case", "Level", true},
		{"unicode", "été", true},
		{"not a palindrome", "gopher", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // subtest, shows up as TestIsPalindrome/<name>
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

// Testing errors.
func TestDivide(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err) // Fatal stops this test immediately
		}
		if got != 5 {
			t.Errorf("got %d, want 5", got)
		}
	})
	t.Run("by zero", func(t *testing.T) {
		_, err := Divide(1, 0)
		if !errors.Is(err, ErrDivideByZero) {
			t.Errorf("err = %v, want ErrDivideByZero", err)
		}
	})
}

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13}
	for n, w := range want {
		if got := Fib(n); got != w {
			t.Errorf("Fib(%d) = %d, want %d", n, got, w)
		}
	}
}

// Benchmarks: BenchmarkXxx(b *testing.B). Run with -bench=.
func BenchmarkFib(b *testing.B) {
	for b.Loop() {
		Fib(50)
	}
}
