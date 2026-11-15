// Topic 02: Variables, constants and basic data types.
//
// Run: go run ./02-variables-constants-types
package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	integers()
	floats()
	arithmetic()
	stringsAndRunes()
	booleans()
	declarations()
	zeroValues()
	constants()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func integers() {
	section("Integers")
	// Signed: int8, int16, int32, int64, int (32 or 64 bit depending on the platform)
	// Unsigned: uint8 (byte), uint16, uint32, uint64, uint
	var small int8 = 127
	var big int64 = math.MaxInt64
	var positive uint8 = 255
	fmt.Println("int8 max:", small, "| int64 max:", big, "| uint8 max:", positive)

	// Integers wrap around (overflow) silently at runtime.
	small++
	positive++
	fmt.Println("int8 127 + 1 =", small, "| uint8 255 + 1 =", positive)
}

func floats() {
	section("Floats")
	// float32 has ~7 significant digits, float64 has ~15.
	var f32 float32 = 12345678.9
	var f64 float64 = 12345678.9
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	// Floating point is not exact at runtime...
	x, y := 0.1, 0.2
	fmt.Println("x + y (variables) =", x+y)
	// ...but untyped constants are computed exactly at compile time.
	fmt.Println("0.1 + 0.2 (constants) =", 0.1+0.2)
}

func arithmetic() {
	section("Arithmetic and type conversion")
	var a int = 7
	var b int = 2
	fmt.Println("7 / 2 =", a/b, "(integer division truncates)")
	fmt.Println("7 % 2 =", a%b)

	// Go never converts types implicitly. Mixing int and float64 is a compile error:
	//   var c float64 = 1.5
	//   a + c // invalid operation: mismatched types int and float64
	var c float64 = 1.5
	fmt.Println("float64(7) + 1.5 =", float64(a)+c)
	fmt.Println("float64(7) / 2 =", float64(a)/2)
}

func stringsAndRunes() {
	section("Strings and runes")
	greeting := "Hello" + ", " + "World"
	fmt.Println(greeting)

	raw := `Raw strings use backticks.
They can span lines and ignore escapes like \n`
	fmt.Println(raw)

	// len() counts BYTES, not characters.
	word := "γειά"
	fmt.Println("len(\"γειά\") =", len(word))
	fmt.Println("utf8.RuneCountInString(\"γειά\") =", utf8.RuneCountInString(word))

	// A rune is an alias for int32 and holds a Unicode code point.
	var r rune = 'a'
	fmt.Println("'a' as rune =", r, "| as string =", string(r))
}

func booleans() {
	section("Booleans")
	isGoFun := true
	isHard := false
	fmt.Println("true && false =", isGoFun && isHard)
	fmt.Println("true || false =", isGoFun || isHard)
	fmt.Println("!true =", !isGoFun)
}

func declarations() {
	section("Ways to declare variables")
	var explicit int = 1 // full form
	var inferred = 2     // type inferred
	short := 3           // short form, only inside functions
	var x, y = 4, "five" // multiple at once
	p, q := 6.0, true    // multiple, short form
	fmt.Println(explicit, inferred, short, x, y, p, q)

	var (
		name    = "gopher"
		version = 1.27
	)
	fmt.Printf("name=%s version=%.2f (%T)\n", name, version, version)
}

func zeroValues() {
	section("Zero values (declared but not assigned)")
	var i int
	var f float64
	var s string
	var b bool
	var sl []int
	var m map[string]int
	var p *int
	fmt.Printf("int=%d float=%v string=%q bool=%t\n", i, f, s, b)
	fmt.Println("slice nil?", sl == nil, "| map nil?", m == nil, "| pointer nil?", p == nil)
}

// Package-level constants.
const Pi = 3.14159

// iota creates incrementing constants: an easy way to build enums.
type Weekday int

const (
	Sunday  Weekday = iota // 0
	Monday                 // 1
	Tuesday                // 2
)

const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1 << 10
	MB                    // 1 << 20
	GB                    // 1 << 30
)

func constants() {
	section("Constants and iota")
	const greeting = "hi" // must be known at compile time; cannot be changed
	// greeting = "bye"   // compile error: cannot assign to greeting
	fmt.Println("Pi:", Pi, "| greeting:", greeting)
	fmt.Println("Sunday:", Sunday, "Monday:", Monday, "Tuesday:", Tuesday)
	fmt.Println("KB:", KB, "MB:", MB, "GB:", GB)
}
