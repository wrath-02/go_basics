// Topic 06: Structs, methods and interfaces.
//
// Run: go run ./06-structs-interfaces
package main

import (
	"fmt"
	"math"
)

func main() {
	structBasics()
	embedding()
	methods()
	interfaces()
	stringer()
	typeAssertions()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// ---------- struct basics ----------

type Owner struct {
	Name string
}

// GasEngine groups related fields into one custom type.
type GasEngine struct {
	MilesPerGallon float64
	Gallons        float64
	Owner          Owner // a struct can contain another struct
}

func structBasics() {
	section("Struct basics")
	var zero GasEngine // every field gets its zero value
	fmt.Printf("zero value: %+v\n", zero)

	// Named fields (recommended): order doesn't matter, missing fields are zero.
	car := GasEngine{MilesPerGallon: 25, Gallons: 15, Owner: Owner{Name: "Alex"}}
	fmt.Printf("named:      %+v\n", car)

	// Positional fields: must list every field in order (fragile).
	car2 := GasEngine{30, 10, Owner{"Sam"}}
	fmt.Printf("positional: %+v\n", car2)

	car.Gallons = 12 // update a field with dot notation
	fmt.Println("owner:", car.Owner.Name, "| gallons:", car.Gallons)

	// Anonymous struct: defined and used in one place (not reusable).
	point := struct{ X, Y int }{X: 1, Y: 2}
	fmt.Printf("anonymous:  %+v\n", point)

	// Structs are comparable with == when all their fields are comparable.
	fmt.Println("car == car2?", car == car2)
}

// ---------- embedding ----------

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s (%d)", p.Name, p.Age)
}

// Employee EMBEDS Person: Person's fields and methods are "promoted".
// This is composition, Go's alternative to inheritance.
type Employee struct {
	Person
	Company string
}

func embedding() {
	section("Embedding (composition over inheritance)")
	e := Employee{Person: Person{Name: "Priya", Age: 28}, Company: "Gophers Inc"}
	fmt.Println("e.Name (promoted):", e.Name)
	fmt.Println("e.Person.Name (explicit):", e.Person.Name)
	fmt.Println("e.Introduce():", e.Introduce())
}

// ---------- methods ----------

// A method is a function with a receiver. (e GasEngine) is a VALUE receiver.
func (e GasEngine) MilesLeft() float64 {
	return e.MilesPerGallon * e.Gallons
}

type ElectricEngine struct {
	MilesPerKWh float64
	KWh         float64
}

func (e ElectricEngine) MilesLeft() float64 {
	return e.MilesPerKWh * e.KWh
}

func methods() {
	section("Methods")
	gas := GasEngine{MilesPerGallon: 25, Gallons: 15}
	electric := ElectricEngine{MilesPerKWh: 4, KWh: 60}
	fmt.Println("gas.MilesLeft()      =", gas.MilesLeft())
	fmt.Println("electric.MilesLeft() =", electric.MilesLeft())
	// Methods that MODIFY the struct need a pointer receiver: see topic 07.
}

// ---------- interfaces ----------

// Engine is an interface: a set of method signatures.
// Any type with a MilesLeft() float64 method satisfies it automatically:
// there is no "implements" keyword.
type Engine interface {
	MilesLeft() float64
}

// canMakeIt works with ANY Engine.
func canMakeIt(e Engine, miles float64) string {
	if miles <= e.MilesLeft() {
		return "You can make it there!"
	}
	return "Need to fuel up first!"
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct{ W, H float64 }
type Circle struct{ R float64 }

func (r Rect) Area() float64        { return r.W * r.H }
func (r Rect) Perimeter() float64   { return 2 * (r.W + r.H) }
func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

func interfaces() {
	section("Interfaces")
	gas := GasEngine{MilesPerGallon: 25, Gallons: 15}   // 375 miles
	electric := ElectricEngine{MilesPerKWh: 4, KWh: 60} // 240 miles
	fmt.Println("gas, 300 miles:     ", canMakeIt(gas, 300))
	fmt.Println("electric, 300 miles:", canMakeIt(electric, 300))

	// A slice of interfaces can hold different concrete types.
	shapes := []Shape{Rect{W: 3, H: 4}, Circle{R: 1}}
	for _, s := range shapes {
		fmt.Printf("%-14T area=%6.2f perimeter=%6.2f\n", s, s.Area(), s.Perimeter())
	}
}

// ---------- fmt.Stringer ----------

// Implementing String() string (the fmt.Stringer interface) controls how
// fmt prints your type.
func (p Person) String() string {
	return p.Name + " (" + fmt.Sprint(p.Age) + " yrs)"
}

func stringer() {
	section("fmt.Stringer: custom printing")
	fmt.Println(Person{Name: "Lee", Age: 40})
}

// ---------- type assertions and type switches ----------

func describe(v any) string { // any == interface{}: holds a value of any type
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int doubled = %d", x*2)
	case string:
		return fmt.Sprintf("string of length %d", len(x))
	case Shape:
		return fmt.Sprintf("a Shape with area %.2f", x.Area())
	default:
		return fmt.Sprintf("something else: %T", x)
	}
}

func typeAssertions() {
	section("Type assertions and type switches")
	var e Engine = GasEngine{MilesPerGallon: 25, Gallons: 2}

	// Type assertion: get the concrete type back out. Use "comma ok" to avoid panics.
	if gas, ok := e.(GasEngine); ok {
		fmt.Println("it's a GasEngine with", gas.Gallons, "gallons")
	}
	if _, ok := e.(ElectricEngine); !ok {
		fmt.Println("it's NOT an ElectricEngine")
	}

	for _, v := range []any{21, "gopher", Rect{W: 2, H: 5}, 3.5} {
		fmt.Println(describe(v))
	}
}
