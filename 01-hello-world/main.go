// Topic 01: Hello World, program structure, printing and reading input.
//
// Run:        go run ./01-hello-world
// With input: echo "Gopher" | go run ./01-hello-world
package main // every executable program lives in package "main"

import ( // imports are grouped; unused imports are a compile error
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main is the entry point. It takes no arguments and returns nothing.
func main() {
	fmt.Println("Hello, World!")

	section("Print vs Println vs Printf")
	fmt.Print("Print adds no newline. ")
	fmt.Print("See?\n")
	fmt.Println("Println", "adds", "spaces", "and", "a", "newline")
	fmt.Printf("Printf formats values: %s was released in %d\n", "Go", 2009)

	section("Common format verbs")
	gopher := struct {
		Name string
		Age  int
	}{"Gopher", 3}
	fmt.Printf("%%v   -> %v\n", gopher)      // default format
	fmt.Printf("%%+v  -> %+v\n", gopher)     // with field names
	fmt.Printf("%%T   -> %T\n", gopher.Name) // type
	fmt.Printf("%%d   -> %d\n", 42)          // integer
	fmt.Printf("%%5.2f -> %5.2f\n", 3.14159) // float, width 5, 2 decimals
	fmt.Printf("%%q   -> %q\n", "quoted")    // quoted string
	fmt.Printf("%%t   -> %t\n", true)        // boolean
	fmt.Printf("%%b   -> %b\n", 5)           // binary

	section("Sprintf returns a string")
	label := fmt.Sprintf("%s-%03d", "ID", 7)
	fmt.Println(label)

	section("Reading input from stdin")
	greetFromStdin()
}

func section(title string) {
	fmt.Printf("\n--- %s ---\n", title)
}

// greetFromStdin reads one line if something was piped in.
func greetFromStdin() {
	info, err := os.Stdin.Stat()
	if err != nil || info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println(`No input piped. Try: echo "Gopher" | go run ./01-hello-world`)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	name := strings.TrimSpace(line)
	if name == "" {
		fmt.Println(`No input piped. Try: echo "Gopher" | go run ./01-hello-world`)
		return
	}
	fmt.Printf("Nice to meet you, %s!\n", name)
}
