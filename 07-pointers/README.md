# 07 · Pointers

> **Definition:** A **pointer** is a variable that stores the **memory address** of another value. `&x` gives the address of `x`; `*p` reads or writes the value stored at that address (called *dereferencing*). The type `*T` means "pointer to a T". The zero value of a pointer is `nil`.

```bash
go run ./07-pointers
```

## Key concepts

| Syntax | Meaning |
|---|---|
| `p := &x` | `p` holds the address of `x` (type `*int`) |
| `*p` | the value at that address |
| `*p = 20` | write through the pointer, which changes `x` |
| `new(T)` | allocate a zero `T` and return a `*T` |
| `p.Field` | automatic dereference for struct pointers (same as `(*p).Field`) |

## Examples

### & and *

```go
x := 10
p := &x
*p = 20
q := new(int32)
```

```text
x = 10, type of p = *int
*p = 10 (read the value at the address)
after *p = 20, x = 20
*new(int32) = 0
p == p2? true
```

### nil pointers

```go
var p *int
fmt.Println(p == nil) // true
// fmt.Println(*p)    // 💥 panic: invalid memory address or nil pointer dereference
```

### Pass by value vs pass by pointer

Go **always copies** arguments. Pass a pointer to modify the caller's value, or to avoid copying large data.

```go
func squareCopy(arr [5]float64) [5]float64 { ... } // works on a copy
func squareInPlace(arr *[5]float64)          { ... } // works on the original

nums := [5]float64{1, 2, 3, 4, 5}
result := squareCopy(nums)
squareInPlace(&nums)
```

```text
squareCopy  -> result: [1 4 9 16 25] | original: [1 2 3 4 5]
squareInPlace -> original is now: [1 4 9 16 25]
```

### Slices and maps already share data

```go
func doubleAll(s []int) { for i := range s { s[i] *= 2 } }
s := []int{1, 2, 3}
doubleAll(s) // [2 4 6]: no pointer needed for element changes
```

### Value receivers vs pointer receivers

```go
func (c Counter) IncrementCopy() { c.count++ } // modifies a copy, so the change is lost
func (c *Counter) Increment()    { c.count++ } // modifies the original
```

```text
after IncrementCopy(): 0
after Increment() x2: 2
```

### Returning pointers from constructors

```go
func newUser(name string) *User {
	return &User{Name: name, Email: name + "@example.com"} // safe: escapes to the heap
}
u := newUser("gopher")
alias := u
alias.Name = "renamed" // u.Name changes too, because both point to the same User
```

## When to use pointers

- ✅ A method or function must **modify** its argument or receiver.
- ✅ The struct is **large** and copying would be expensive.
- ✅ You need to express "no value" with `nil`.
- ❌ Don't use pointers to slices or maps just to modify their elements. They already reference shared data.

## Gotchas

- Go has **no pointer arithmetic** (unlike C).
- If any method of a type uses a pointer receiver, use pointer receivers for all of its methods, for consistency.

## Try it yourself

1. Write `swap(a, b *int)` that swaps two integers.
2. Write a `BankAccount` with `Deposit` and `Withdraw` methods using pointer receivers.

---
[← Structs & Interfaces](../06-structs-interfaces/) · [🏠 Index](../README.md) · [Next: Goroutines →](../08-goroutines/)
