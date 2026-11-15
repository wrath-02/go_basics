# 06 · Structs & Interfaces

> **Definitions:**
> - **Struct:** a custom type that groups named fields together (like a lightweight class without inheritance).
> - **Method:** a function attached to a type through a **receiver**: `func (e GasEngine) MilesLeft() float64`.
> - **Interface:** a set of method signatures. A type satisfies an interface **implicitly**, simply by having those methods. There is no `implements` keyword.

```bash
go run ./06-structs-interfaces
```

## Structs

```go
type Owner struct{ Name string }

type GasEngine struct {
	MilesPerGallon float64
	Gallons        float64
	Owner          Owner
}

var zero GasEngine                                                       // all fields zero
car := GasEngine{MilesPerGallon: 25, Gallons: 15, Owner: Owner{"Alex"}} // named (recommended)
car2 := GasEngine{30, 10, Owner{"Sam"}}                                 // positional
point := struct{ X, Y int }{1, 2}                                       // anonymous struct
```

```text
zero value: {MilesPerGallon:0 Gallons:0 Owner:{Name:}}
named:      {MilesPerGallon:25 Gallons:15 Owner:{Name:Alex}}
positional: {MilesPerGallon:30 Gallons:10 Owner:{Name:Sam}}
anonymous:  {X:1 Y:2}
```

## Embedding: composition instead of inheritance

```go
type Person struct{ Name string; Age int }
func (p Person) Introduce() string { ... }

type Employee struct {
	Person          // embedded: its fields and methods are promoted
	Company string
}

e := Employee{Person: Person{"Priya", 28}, Company: "Gophers Inc"}
e.Name        // promoted field
e.Introduce() // promoted method
```

```text
e.Name (promoted): Priya
e.Introduce(): Hi, I'm Priya (28)
```

## Methods and interfaces

```go
type Engine interface {
	MilesLeft() float64
}

func (e GasEngine) MilesLeft() float64      { return e.MilesPerGallon * e.Gallons }
func (e ElectricEngine) MilesLeft() float64 { return e.MilesPerKWh * e.KWh }

// Works with ANY type that has MilesLeft()
func canMakeIt(e Engine, miles float64) string {
	if miles <= e.MilesLeft() {
		return "You can make it there!"
	}
	return "Need to fuel up first!"
}
```

```text
gas, 300 miles:      You can make it there!
electric, 300 miles: Need to fuel up first!
```

A slice of interface values can hold different concrete types:

```go
shapes := []Shape{Rect{W: 3, H: 4}, Circle{R: 1}}
for _, s := range shapes {
	fmt.Printf("%-14T area=%6.2f perimeter=%6.2f\n", s, s.Area(), s.Perimeter())
}
```

```text
main.Rect      area= 12.00 perimeter= 14.00
main.Circle    area=  3.14 perimeter=  6.28
```

## fmt.Stringer: control how your type prints

```go
func (p Person) String() string { return p.Name + " (" + fmt.Sprint(p.Age) + " yrs)" }
fmt.Println(Person{Name: "Lee", Age: 40}) // Lee (40 yrs)
```

## Type assertions and type switches

```go
var e Engine = GasEngine{Gallons: 2}
if gas, ok := e.(GasEngine); ok { ... }  // safe assertion

switch x := v.(type) {                  // v is `any`
case int:    ...
case string: ...
case Shape:  ...
default:     ...
}
```

```text
it's a GasEngine with 2 gallons
it's NOT an ElectricEngine
int doubled = 42
string of length 6
a Shape with area 10.00
something else: float64
```

## Gotchas

- `any` is just an alias for `interface{}`, the empty interface that every type satisfies.
- `v.(T)` **panics** if `v` isn't a `T`. Use the two-value form `t, ok := v.(T)`.
- Keep interfaces small (1–3 methods). "Accept interfaces, return structs."

## Try it yourself

1. Add a `Triangle` type that satisfies `Shape`.
2. Give `GasEngine` a `String()` method and print one.

---
[← Strings, Runes & Bytes](../05-strings-runes-bytes/) · [🏠 Index](../README.md) · [Next: Pointers →](../07-pointers/)
