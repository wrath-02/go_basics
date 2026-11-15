// Topic 07: Pointers.
//
// Run: go run ./07-pointers
package main

import "fmt"

func main() {
	pointerBasics()
	nilPointers()
	passByValueVsPointer()
	slicesAlreadyShare()
	pointerReceivers()
	pointersToStructs()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func pointerBasics() {
	section("& (address of) and * (dereference)")
	x := 10
	p := &x // p has type *int and stores the memory address of x

	fmt.Printf("x = %d, type of p = %T\n", x, p)
	fmt.Println("*p =", *p, "(read the value at the address)")

	*p = 20 // write through the pointer: changes x
	fmt.Println("after *p = 20, x =", x)

	// new(T) allocates a zero-valued T and returns a *T.
	q := new(int32)
	fmt.Println("*new(int32) =", *q)
	*q = 5
	fmt.Println("after *q = 5:", *q)

	// Two pointers to the same variable are equal.
	p2 := &x
	fmt.Println("p == p2?", p == p2)
}

func nilPointers() {
	section("nil pointers")
	var p *int // zero value of a pointer is nil
	fmt.Println("p == nil?", p == nil)
	// fmt.Println(*p) // runtime panic: invalid memory address or nil pointer dereference

	// Always check before dereferencing a pointer that might be nil.
	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("p is nil, not dereferencing")
	}
}

// Arguments are always COPIED. squareCopy gets its own copy of the array.
func squareCopy(arr [5]float64) [5]float64 {
	for i := range arr {
		arr[i] *= arr[i]
	}
	return arr
}

// squareInPlace gets a pointer, so no copy is made and the caller's array changes.
func squareInPlace(arr *[5]float64) {
	for i := range arr { // Go auto-dereferences: arr[i] == (*arr)[i]
		arr[i] *= arr[i]
	}
}

func passByValueVsPointer() {
	section("Pass by value vs pass by pointer")
	nums := [5]float64{1, 2, 3, 4, 5}
	result := squareCopy(nums)
	fmt.Println("squareCopy  -> result:", result, "| original:", nums)

	squareInPlace(&nums)
	fmt.Println("squareInPlace -> original is now:", nums)
	// Passing a pointer also avoids copying large values (saves memory and time).
}

func doubleAll(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func slicesAlreadyShare() {
	section("Slices and maps already reference shared data")
	// A slice header contains a pointer to its backing array,
	// so modifying elements inside a function affects the caller.
	s := []int{1, 2, 3}
	doubleAll(s)
	fmt.Println("after doubleAll(slice):", s)

	// The same is true for maps.
	m := map[string]int{"a": 1}
	func(mm map[string]int) { mm["a"] = 100 }(m)
	fmt.Println("map after modification in function:", m)
}

type Counter struct {
	count int
}

// Value receiver: works on a COPY, so the change is lost.
func (c Counter) IncrementCopy() {
	c.count++
}

// Pointer receiver: works on the original.
func (c *Counter) Increment() {
	c.count++
}

func pointerReceivers() {
	section("Value receivers vs pointer receivers")
	c := Counter{}
	c.IncrementCopy()
	fmt.Println("after IncrementCopy():", c.count)

	c.Increment() // Go takes &c automatically
	c.Increment()
	fmt.Println("after Increment() x2:", c.count)
	// Rule of thumb: if any method needs a pointer receiver, use pointer
	// receivers for all methods of that type.
}

type User struct {
	Name  string
	Email string
}

// newUser returns a pointer. It's safe to return the address of a local
// variable: Go moves it to the heap automatically ("escape analysis").
func newUser(name string) *User {
	return &User{Name: name, Email: name + "@example.com"}
}

func pointersToStructs() {
	section("Pointers to structs")
	u := newUser("gopher")
	fmt.Printf("u is %T: %+v\n", u, *u)

	u.Email = "new@example.com" // shorthand for (*u).Email
	fmt.Println("updated email:", u.Email)

	alias := u // copies the POINTER, both point to the same User
	alias.Name = "renamed"
	fmt.Println("u.Name after alias.Name = \"renamed\":", u.Name)
}
