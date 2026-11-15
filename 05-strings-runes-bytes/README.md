# 05 · Strings, Runes & Bytes

> **Definitions:**
> - **string:** an immutable, read-only sequence of **bytes**, usually holding UTF-8 encoded text.
> - **byte:** an alias for `uint8`, one raw byte.
> - **rune:** an alias for `int32`, one Unicode **code point** (a "character"). A character can take 1 to 4 bytes in UTF-8.

```bash
go run ./05-strings-runes-bytes
```

## Indexing vs ranging

```go
word := "résumé"
len(word)                     // 8 (bytes), but 6 characters
word[0]                       // byte 114 -> 'r'
word[1]                       // byte 195 -> only HALF of 'é'
for i, r := range word { ... } // range decodes runes; the index jumps
```

```text
len(word) = 8 bytes, but 6 characters
word[0] = 114 (uint8) -> r
word[1] = 195 (uint8) -> first byte of 'é' only
range: [0]r [1]é [3]s [4]u [5]m [6]é
```

## Converting to []rune and []byte

```go
runes := []rune(word) // one element per character -> safe to index
bytes := []byte(word) // raw UTF-8 bytes
```

```text
[]rune: [114 233 115 117 109 233] len: 6
runes[1] = é (code point U+00E9)
[]byte: [114 195 169 115 117 109 195 169] len: 8
byte 'A' = 65, rune '世' = 19990, back to string: A世
```

## Strings are immutable → use strings.Builder

```go
s := "hello"
// s[0] = 'H'          // ❌ compile error
bs := []byte(s); bs[0] = 'H' // ✅ works on a copy

var sb strings.Builder // efficient concatenation in loops
for _, l := range letters {
	sb.WriteString(l)
}
sb.String()            // "gopher"
```

> `s += x` in a loop allocates a new string on every iteration. `strings.Builder` writes into one growing buffer.

## The `strings` package

```text
TrimSpace:   "Go is simple, Go is fast"
ToUpper:     GO IS SIMPLE, GO IS FAST
Contains:    true
HasPrefix:   true
Count 'Go':  2
Index 'is':  3
ReplaceAll:  Gopher is simple, Gopher is fast
Split ', ': ["Go is simple" "Go is fast"]
Fields:     ["Go" "is" "simple," "Go" "is" "fast"]
Join:        a-b-c
Repeat:      ababab
```

## Number ⇄ string with `strconv`

```go
strconv.Itoa(42)          // "42"
string(rune(42))          // "*"  ⚠️ the character with code 42, NOT "42"
strconv.Atoi("123")       // 123, nil
strconv.Atoi("12a")       // 0, error
strconv.ParseFloat("3.14", 64)
strconv.ParseBool("true")
```

```text
strconv.Itoa(42) = "42"
string(rune(42)) = "*"  <- not what you want
Atoi("123") = 123 err: <nil>
Atoi("12a") err: strconv.Atoi: parsing "12a": invalid syntax
```

## Gotchas

- `len(s)` is the **byte** count. Use `utf8.RuneCountInString(s)` for characters.
- Use `range` or `[]rune` to work with characters. Never index a non-ASCII string byte by byte.

## Try it yourself

1. Reverse a string correctly, including `"résumé"` (hint: `[]rune`).
2. Count the vowels in a sentence using `strings.ContainsRune`.

---
[← Arrays, Slices, Maps & Loops](../04-arrays-slices-maps-loops/) · [🏠 Index](../README.md) · [Next: Structs & Interfaces →](../06-structs-interfaces/)
