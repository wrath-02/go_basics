// Topic 04: Arrays, slices, maps and loops.
//
// Run:       go run ./04-arrays-slices-maps-loops
// Benchmark: go test -bench=. ./04-arrays-slices-maps-loops
package main

import (
	"fmt"
	"maps"
	"slices"
	"time"
)

func main() {
	arrays()
	slicesBasics()
	slicesSharing()
	mapsBasics()
	loops()
	preallocation()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func arrays() {
	section("Arrays (fixed length, stored contiguously)")
	var intArr [3]int32 // zero-valued: [0 0 0]
	intArr[1] = 123
	fmt.Println("intArr:", intArr, "| intArr[1]:", intArr[1], "| intArr[1:3]:", intArr[1:3])

	letters := [...]string{"a", "b", "c"} // [...] lets the compiler count
	fmt.Println("letters:", letters, "len:", len(letters))

	// Arrays are VALUES: assigning copies every element.
	copied := letters
	copied[0] = "z"
	fmt.Println("original:", letters, "| copy:", copied)
}

func slicesBasics() {
	section("Slices (dynamic view over an array)")
	nums := []int{4, 5, 6}
	fmt.Printf("nums=%v len=%d cap=%d\n", nums, len(nums), cap(nums))

	nums = append(nums, 7) // capacity grows (roughly doubles) when full
	fmt.Printf("after append(7): nums=%v len=%d cap=%d\n", nums, len(nums), cap(nums))

	more := []int{8, 9}
	nums = append(nums, more...) // append another slice
	fmt.Println("after append(more...):", nums)

	// make(type, length, capacity)
	buf := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10): %v len=%d cap=%d\n", buf, len(buf), cap(buf))

	// Slicing: s[low:high] includes low, excludes high.
	fmt.Println("nums[1:3] =", nums[1:3], "| nums[:2] =", nums[:2], "| nums[4:] =", nums[4:])

	// The slices package (Go 1.21+) has many helpers.
	unsorted := []int{3, 1, 2}
	slices.Sort(unsorted)
	fmt.Println("sorted:", unsorted, "| contains 2?", slices.Contains(unsorted, 2), "| index of 3:", slices.Index(unsorted, 3))
}

func slicesSharing() {
	section("Slices share their underlying array (gotcha!)")
	original := []int{1, 2, 3, 4}
	view := original[0:2]
	view[0] = 100
	fmt.Println("original after view[0] = 100:", original)

	// copy() makes an independent slice.
	independent := make([]int, len(original))
	copy(independent, original)
	independent[0] = -1
	fmt.Println("original:", original, "| independent:", independent)

	// slices.Clone does the same in one call.
	clone := slices.Clone(original)
	clone[1] = -2
	fmt.Println("original:", original, "| clone:", clone)
}

func mapsBasics() {
	section("Maps (key -> value)")
	ages := map[string]int{"Adam": 23, "Sarah": 45}
	ages["Jason"] = 31 // add / update
	fmt.Println("ages:", ages, "len:", len(ages))

	// Missing keys return the zero value, so use the "comma ok" idiom.
	fmt.Println(`ages["Nobody"] =`, ages["Nobody"])
	if age, ok := ages["Sarah"]; ok {
		fmt.Println("Sarah is", age)
	}
	if _, ok := ages["Nobody"]; !ok {
		fmt.Println("Nobody is not in the map")
	}

	delete(ages, "Adam")
	fmt.Println("after delete(Adam):", ages)

	// make() creates an empty map. A nil map (var m map[string]int) panics on write.
	scores := make(map[string][]int)
	scores["math"] = append(scores["math"], 90, 85)
	fmt.Println("scores:", scores)
}

func loops() {
	section("Loops (Go only has 'for')")

	// classic three-part loop
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("<- classic for")

	// range over an integer (Go 1.22+)
	for i := range 3 {
		fmt.Print(i, " ")
	}
	fmt.Println("<- range over int")

	// "while" loop
	n := 1
	for n < 50 {
		n *= 3
	}
	fmt.Println("while-style loop ended with n =", n)

	// range over a slice: index and value
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("%d:%s ", i, fruit)
	}
	fmt.Println("<- range over slice")

	// range over a map: iteration order is RANDOM, so sort keys for stable output.
	stock := map[string]int{"pens": 10, "books": 3, "cups": 7}
	for _, key := range slices.Sorted(maps.Keys(stock)) {
		fmt.Printf("%s=%d ", key, stock[key])
	}
	fmt.Println("<- range over map (sorted keys)")

	// range over a string gives runes (see topic 05)
	for i, ch := range "héllo" {
		fmt.Printf("%d:%c ", i, ch)
	}
	fmt.Println("<- range over string")

	// break and continue
	for i := range 10 {
		if i%2 == 0 {
			continue // skip even numbers
		}
		if i > 7 {
			break // stop the loop
		}
		fmt.Print(i, " ")
	}
	fmt.Println("<- continue/break")

	// labeled break exits an outer loop
outer:
	for i := range 3 {
		for j := range 3 {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println("<- labeled break")
}

// Preallocating capacity avoids repeated grow-and-copy work.
func fillWithoutPrealloc(n int) []int {
	var s []int
	for i := range n {
		s = append(s, i)
	}
	return s
}

func fillWithPrealloc(n int) []int {
	s := make([]int, 0, n)
	for i := range n {
		s = append(s, i)
	}
	return s
}

func preallocation() {
	section("Why preallocate? (timings vary per machine)")
	const n = 1_000_000
	start := time.Now()
	fillWithoutPrealloc(n)
	without := time.Since(start)

	start = time.Now()
	fillWithPrealloc(n)
	with := time.Since(start)

	fmt.Printf("without preallocation: %v\n", without.Round(time.Microsecond))
	fmt.Printf("with preallocation:    %v\n", with.Round(time.Microsecond))
}
