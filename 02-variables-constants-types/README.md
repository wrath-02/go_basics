# 02 · Variables, Constants & Basic Data Types

> **Definition:** A **variable** is a named storage location with a fixed **type**. Go is *statically typed*: every variable's type is known at compile time and never changes. A **constant** is a value fixed at compile time that can never be reassigned.

```bash
go run ./02-variables-constants-types
```

## Basic types

| Category | Types | Zero value |
|---|---|---|
| Signed integers | `int8` `int16` `int32` `int64` `int` | `0` |
| Unsigned integers | `uint8` (=`byte`) `uint16` `uint32` `uint64` `uint` | `0` |
| Floats | `float32` `float64` | `0` |
| Complex | `complex64` `complex128` | `0` |
| Text | `string`, `rune` (=`int32`, one Unicode code point) | `""`, `0` |
| Boolean | `bool` | `false` |

`int` and `uint` are 64-bit on 64-bit machines. When in doubt, use `int` and `float64`.

## Declaring variables

```go
var explicit int = 1 // full form
var inferred = 2     // type inferred from the value
short := 3           // short form, only inside functions
var x, y = 4, "five" // several at once
p, q := 6.0, true

var (                // grouped
	name    = "gopher"
	version = 1.27
)
```

```text
1 2 3 4 five 6 true
name=gopher version=1.27 (float64)
```

## Examples

### Integer overflow wraps silently

```go
var small int8 = 127
var positive uint8 = 255
small++
positive++
fmt.Println("int8 127 + 1 =", small, "| uint8 255 + 1 =", positive)
```

```text
int8 127 + 1 = -128 | uint8 255 + 1 = 0
```

### Float precision

```go
var f32 float32 = 12345678.9
var f64 float64 = 12345678.9
x, y := 0.1, 0.2
fmt.Println("float32:", f32)
fmt.Println("float64:", f64)
fmt.Println("x + y (variables) =", x+y)
fmt.Println("0.1 + 0.2 (constants) =", 0.1+0.2)
```

```text
float32: 1.2345679e+07
float64: 1.23456789e+07
x + y (variables) = 0.30000000000000004
0.1 + 0.2 (constants) = 0.3
```

> Untyped constants are computed with arbitrary precision by the compiler, so the constant expression is exact.

### No implicit conversion

```go
var a int = 7
var c float64 = 1.5
// a + c                       // ❌ invalid operation: mismatched types int and float64
fmt.Println(float64(a) + c)    // ✅ 8.5
fmt.Println(a / 2, a % 2)      // integer division: 3 1
```

### Strings: len() counts bytes

```go
word := "γειά"
fmt.Println(len(word))                    // 8 (bytes)
fmt.Println(utf8.RuneCountInString(word)) // 4 (characters)
var r rune = 'a'
fmt.Println(r, string(r))                 // 97 a
```

### Zero values

```go
var i int; var f float64; var s string; var b bool
var sl []int; var m map[string]int; var p *int
```

```text
int=0 float=0 string="" bool=false
slice nil? true | map nil? true | pointer nil? true
```

### Constants and `iota`

```go
const Pi = 3.14159

type Weekday int
const (
	Sunday Weekday = iota // 0
	Monday                // 1
	Tuesday               // 2
)

const (
	_  = iota             // skip 0
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)
```

```text
Sunday: 0 Monday: 1 Tuesday: 2
KB: 1024 MB: 1048576 GB: 1073741824
```

## Gotchas

- `:=` only works **inside functions**. At package level, use `var`.
- `:=` needs at least one **new** variable on the left-hand side.
- Converting a float to an int truncates: `int(3.9) == 3`.

## Try it yourself

1. Declare a `uint16` set to its maximum value (`math.MaxUint16`) and add 1. What happens?
2. Use `iota` to create `Small`, `Medium`, `Large` sizes starting from 1.

---
[← Hello World](../01-hello-world/) · [🏠 Index](../README.md) · [Next: Functions & Control Flow →](../03-functions-control-flow/)
