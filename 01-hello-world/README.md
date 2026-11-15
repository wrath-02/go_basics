# 01 · Hello World & Program Structure

> **Definition:** A Go program is a set of **packages**. An executable program must have a package named `main` with a function `func main()`, which is where execution starts. Code is compiled to a single native binary.

```bash
go run ./01-hello-world                         # compile + run
echo "Gopher" | go run ./01-hello-world         # with piped input
go build -o hello ./01-hello-world && ./hello   # build a binary
```

## Key concepts

| Concept | What it means |
|---|---|
| `package main` | Marks this package as an executable (not a library). |
| `import "fmt"` | Brings in another package. **Unused imports are a compile error.** |
| `func main()` | Entry point. No arguments, no return value. |
| `go mod init <path>` | Creates `go.mod`, which names your module and lists dependencies. |
| `gofmt` | The official formatter. Every Go file looks the same; use it on save. |
| Exported names | Start with a capital letter (`fmt.Println`). Lowercase = private to the package. |

## Examples

### Print, Println, Printf

```go
fmt.Print("Print adds no newline. ")
fmt.Print("See?\n")
fmt.Println("Println", "adds", "spaces", "and", "a", "newline")
fmt.Printf("Printf formats values: %s was released in %d\n", "Go", 2009)
```

```text
Print adds no newline. See?
Println adds spaces and a newline
Printf formats values: Go was released in 2009
```

### Format verbs

```go
gopher := struct {
	Name string
	Age  int
}{"Gopher", 3}
fmt.Printf("%%v   -> %v\n", gopher)
fmt.Printf("%%+v  -> %+v\n", gopher)
fmt.Printf("%%T   -> %T\n", gopher.Name)
fmt.Printf("%%5.2f -> %5.2f\n", 3.14159)
fmt.Printf("%%q   -> %q\n", "quoted")
fmt.Printf("%%b   -> %b\n", 5)
```

```text
%v   -> {Gopher 3}
%+v  -> {Name:Gopher Age:3}
%T   -> string
%5.2f ->  3.14
%q   -> "quoted"
%b   -> 101
```

| Verb | Use |
|---|---|
| `%v` / `%+v` / `%#v` | default / with field names / Go syntax |
| `%T` | type of the value |
| `%d` `%b` `%x` | integer: decimal / binary / hex |
| `%f` `%.2f` `%e` | float |
| `%s` `%q` | string / quoted string |
| `%t` | bool |
| `%p` | pointer address |

### Sprintf returns a string instead of printing

```go
label := fmt.Sprintf("%s-%03d", "ID", 7)
fmt.Println(label) // ID-007
```

### Reading input

```go
reader := bufio.NewReader(os.Stdin)
line, _ := reader.ReadString('\n')
name := strings.TrimSpace(line)
fmt.Printf("Nice to meet you, %s!\n", name)
```

```bash
$ echo "Gopher" | go run ./01-hello-world
...
Nice to meet you, Gopher!
```

> `fmt.Scan(&x)` and `fmt.Scanln(&x)` are simpler alternatives for reading space-separated values.

## Gotchas

- An unused **variable** or **import** fails compilation. Use `_` to discard values on purpose.
- The opening `{` must be on the same line as `func`, `if`, `for`, and so on (semicolons are inserted automatically).

## Try it yourself

1. Print your name and age using a single `Printf`.
2. Read two numbers from stdin with `fmt.Scan` and print their sum.

---
[🏠 Index](../README.md) · [Next: Variables, Constants & Types →](../02-variables-constants-types/)
