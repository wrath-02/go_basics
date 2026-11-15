# 04 · Arrays, Slices, Maps & Loops

> **Definitions:**
> - **Array:** a fixed-length sequence of elements of one type, stored contiguously in memory. The length is part of the type (`[3]int` ≠ `[4]int`).
> - **Slice:** a dynamically-sized *view* into an underlying array. It has a pointer, a length (`len`) and a capacity (`cap`). This is what you use 99% of the time.
> - **Map:** an unordered collection of key → value pairs (a hash table).
> - **Loop:** Go has only one loop keyword, **`for`**, which covers classic, while-style, infinite and `range` loops.

```bash
go run ./04-arrays-slices-maps-loops
go test -bench=. -benchmem ./04-arrays-slices-maps-loops   # preallocation benchmark
```

## Arrays

```go
var intArr [3]int32          // [0 0 0]
intArr[1] = 123
letters := [...]string{"a", "b", "c"} // compiler counts the length

copied := letters            // arrays are VALUES: this copies all elements
copied[0] = "z"
```

```text
intArr: [0 123 0] | intArr[1]: 123 | intArr[1:3]: [123 0]
letters: [a b c] len: 3
original: [a b c] | copy: [z b c]
```

## Slices

```go
nums := []int{4, 5, 6}
nums = append(nums, 7)        // grows capacity when full
nums = append(nums, more...)  // append another slice
buf := make([]int, 3, 10)     // make(type, len, cap)
nums[1:3]                     // low is inclusive, high is exclusive
```

```text
nums=[4 5 6] len=3 cap=3
after append(7): nums=[4 5 6 7] len=4 cap=6
after append(more...): [4 5 6 7 8 9]
make([]int, 3, 10): [0 0 0] len=3 cap=10
nums[1:3] = [5 6] | nums[:2] = [4 5] | nums[4:] = [8 9]
sorted: [1 2 3] | contains 2? true | index of 3: 2
```

### ⚠️ Slices share their backing array

```go
original := []int{1, 2, 3, 4}
view := original[0:2]
view[0] = 100                 // also changes original!

independent := make([]int, len(original))
copy(independent, original)   // or: slices.Clone(original)
```

```text
original after view[0] = 100: [100 2 3 4]
original: [100 2 3 4] | independent: [-1 2 3 4]
```

## Maps

```go
ages := map[string]int{"Adam": 23, "Sarah": 45}
ages["Jason"] = 31                         // add / update

if age, ok := ages["Sarah"]; ok { ... }    // "comma ok": does the key exist?
delete(ages, "Adam")

scores := make(map[string][]int)           // empty, ready to use
```

```text
ages: map[Adam:23 Jason:31 Sarah:45] len: 3
ages["Nobody"] = 0
Sarah is 45
Nobody is not in the map
after delete(Adam): map[Jason:31 Sarah:45]
```

## Loops

```go
for i := 0; i < 3; i++ {}         // classic
for i := range 3 {}               // range over int (Go 1.22+)
for n < 50 { n *= 3 }             // while-style
for {}                            // infinite (exit with break/return)
for i, v := range slice {}        // index, value
for k, v := range someMap {}      // key, value (RANDOM order!)
for i, r := range "héllo" {}      // byte index, rune
```

```text
0 1 2 <- classic for
0 1 2 <- range over int
while-style loop ended with n = 81
0:apple 1:banana 2:cherry <- range over slice
books=3 cups=7 pens=10 <- range over map (sorted keys)
0:h 1:é 3:l 4:l 5:o <- range over string
1 3 5 7 <- continue/break
(0,0) (0,1) (0,2) (1,0) <- labeled break
```

## Performance: preallocate when you know the size

```go
s := make([]int, 0, n) // one allocation instead of many grow-and-copy steps
```

Example run (timings vary by machine):

```text
without preallocation: 10.23ms
with preallocation:    901µs
```

## Gotchas

- Writing to a **nil map** panics (`var m map[string]int; m["a"] = 1`). Create it with `make` or a literal first.
- **Map iteration order is random** by design. Sort the keys if you need stable output: `slices.Sorted(maps.Keys(m))`.
- `append` may or may not allocate a new array, so always use the result: `s = append(s, x)`.

## Try it yourself

1. Count how many times each word appears in a sentence, using `map[string]int` and `strings.Fields`.
2. Remove the element at index `i` from a slice (hint: `slices.Delete`).

---
[← Functions & Control Flow](../03-functions-control-flow/) · [🏠 Index](../README.md) · [Next: Strings, Runes & Bytes →](../05-strings-runes-bytes/)
