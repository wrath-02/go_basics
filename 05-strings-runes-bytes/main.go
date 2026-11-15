// Topic 05: Strings, runes and bytes.
//
// Run: go run ./05-strings-runes-bytes
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	indexingVsRanging()
	runesAndBytes()
	immutabilityAndBuilder()
	stringsPackage()
	conversions()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

func indexingVsRanging() {
	section("A string is a read-only slice of bytes (UTF-8)")
	word := "résumé"
	fmt.Println("word:", word)
	fmt.Println("len(word) =", len(word), "bytes, but", utf8.RuneCountInString(word), "characters")

	// Indexing returns a BYTE (uint8), not a character.
	fmt.Printf("word[0] = %v (%T) -> %c\n", word[0], word[0], word[0])
	fmt.Printf("word[1] = %v (%T) -> first byte of 'é' only\n", word[1], word[1])

	// range decodes UTF-8 and returns RUNES. Notice the index jumps from 1 to 3.
	fmt.Print("range: ")
	for i, r := range word {
		fmt.Printf("[%d]%c ", i, r)
	}
	fmt.Println()
}

func runesAndBytes() {
	section("Converting to []rune and []byte")
	word := "résumé"

	runes := []rune(word) // one element per character
	fmt.Println("[]rune:", runes, "len:", len(runes))
	fmt.Printf("runes[1] = %c (code point U+%04X)\n", runes[1], runes[1])

	bytes := []byte(word) // raw UTF-8 bytes
	fmt.Println("[]byte:", bytes, "len:", len(bytes))

	// byte is an alias for uint8, rune is an alias for int32.
	var b byte = 'A'
	var r rune = '世'
	fmt.Printf("byte 'A' = %d, rune '世' = %d, back to string: %s%s\n", b, r, string(b), string(r))
}

func immutabilityAndBuilder() {
	section("Strings are immutable -> use strings.Builder")
	s := "hello"
	// s[0] = 'H' // compile error: cannot assign to s[0]
	bs := []byte(s)
	bs[0] = 'H'
	fmt.Println("modified copy:", string(bs), "| original:", s)

	// += creates a brand-new string every time (slow in loops).
	// strings.Builder appends into one growing buffer.
	letters := []string{"g", "o", "p", "h", "e", "r"}
	var sb strings.Builder
	for _, l := range letters {
		sb.WriteString(l)
	}
	fmt.Println("built with strings.Builder:", sb.String())
}

func stringsPackage() {
	section("Handy functions from the strings package")
	text := "  Go is simple, Go is fast  "
	trimmed := strings.TrimSpace(text)
	fmt.Printf("TrimSpace:   %q\n", trimmed)
	fmt.Println("ToUpper:    ", strings.ToUpper(trimmed))
	fmt.Println("Contains:   ", strings.Contains(trimmed, "fast"))
	fmt.Println("HasPrefix:  ", strings.HasPrefix(trimmed, "Go"))
	fmt.Println("Count 'Go': ", strings.Count(trimmed, "Go"))
	fmt.Println("Index 'is': ", strings.Index(trimmed, "is"))
	fmt.Println("ReplaceAll: ", strings.ReplaceAll(trimmed, "Go", "Gopher"))
	fmt.Printf("Split ', ': %q\n", strings.Split(trimmed, ", "))
	fmt.Printf("Fields:     %q\n", strings.Fields(trimmed))
	fmt.Println("Join:       ", strings.Join([]string{"a", "b", "c"}, "-"))
	fmt.Println("Repeat:     ", strings.Repeat("ab", 3))
}

func conversions() {
	section("Converting numbers <-> strings with strconv")
	n := 42
	s := strconv.Itoa(n) // int -> string
	fmt.Printf("strconv.Itoa(42) = %q\n", s)

	// string(42) would give "*" (the character with code 42), NOT "42"!
	fmt.Printf("string(rune(42)) = %q  <- not what you want\n", string(rune(n)))

	parsed, err := strconv.Atoi("123") // string -> int
	fmt.Println("Atoi(\"123\") =", parsed, "err:", err)

	_, err = strconv.Atoi("12a")
	fmt.Println("Atoi(\"12a\") err:", err)

	f, _ := strconv.ParseFloat("3.14", 64)
	ok, _ := strconv.ParseBool("true")
	fmt.Println("ParseFloat:", f, "| ParseBool:", ok)
}
