# Go Basics 🐹

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white)
![CI](https://github.com/wrath-02/go_basics/actions/workflows/go.yml/badge.svg)
![Dependencies](https://img.shields.io/badge/dependencies-0-brightgreen)

My notes and runnable examples for learning **Go from scratch**, from Hello World to a working REST API.
Each topic has its own folder with:

- 📖 a **README** with the definition, key concepts, code snippets and their **actual output**, gotchas and exercises
- ▶️ a runnable **`main.go`** with every example commented

Everything uses **only the standard library**.

## 📚 Topics

| # | Topic | What you'll learn | Run |
|---|---|---|---|
| 01 | [Hello World](01-hello-world/) | program structure, `fmt`, format verbs, reading input | `make run T=01` |
| 02 | [Variables, Constants & Types](02-variables-constants-types/) | int/float/string/bool, zero values, conversion, `const`, `iota` | `make run T=02` |
| 03 | [Functions & Control Flow](03-functions-control-flow/) | multiple returns, errors, variadic, closures, `if`, `switch` | `make run T=03` |
| 04 | [Arrays, Slices, Maps & Loops](04-arrays-slices-maps-loops/) | `append`, `len`/`cap`, `make`, comma-ok, `for`/`range`, preallocation | `make run T=04` |
| 05 | [Strings, Runes & Bytes](05-strings-runes-bytes/) | UTF-8, `[]rune`, `strings.Builder`, `strings`, `strconv` | `make run T=05` |
| 06 | [Structs & Interfaces](06-structs-interfaces/) | methods, embedding, implicit interfaces, type switches | `make run T=06` |
| 07 | [Pointers](07-pointers/) | `&`, `*`, `new`, pass by value vs pointer, pointer receivers | `make run T=07` |
| 08 | [Goroutines](08-goroutines/) | `go`, `WaitGroup`, `Mutex`, `RWMutex`, `atomic`, race detector | `make run T=08` |
| 09 | [Channels](09-channels/) | buffered/unbuffered, `close`, `range`, `select`, timeouts, worker pool | `make run T=09` |
| 10 | [Generics](10-generics/) | type parameters, constraints, `~`, `comparable`, generic types | `make run T=10` |
| 11 | [Building an API](11-building-an-api/) | `net/http` routing, middleware, JSON, `slog`, graceful shutdown, `httptest` | `make api` |
| 12 | [Error Handling](12-error-handling/) ⭐ | `%w` wrapping, `errors.Is`/`As`/`Join`, custom errors | `make run T=12` |
| 13 | [Defer, Panic & Recover](13-defer-panic-recover/) ⭐ | cleanup with `defer`, LIFO order, recovering from panics | `make run T=13` |
| 14 | [Testing](14-testing/) ⭐ | table-driven tests, subtests, benchmarks, examples, coverage | `make test-topic T=14` |

⭐ = bonus topics beyond the core tutorial

## 🚀 Getting started

```bash
git clone https://github.com/wrath-02/go_basics.git
cd go_basics

go run ./01-hello-world          # run any topic directly...
make run T=05                    # ...or use the Makefile shortcut
make api                         # start the API on :8000
make test                        # run every test (with the race detector)
make help                        # list all commands
```

Requires **Go 1.25+** (uses `range` over ints, `wg.Go`, `b.Loop`). `go.mod` targets 1.27.

## 🗂️ Project structure

```text
go_basics/
├── 01-hello-world/
│   ├── README.md          # definition + examples + output
│   └── main.go            # runnable code
├── 02-variables-constants-types/
├── ...
├── 11-building-an-api/
│   ├── cmd/api/           # entry point
│   ├── api/               # shared types
│   └── internal/          # handlers, middleware, tools (db)
├── 14-testing/            # library package + *_test.go
├── playground/            # scratch space for experiments
├── .github/workflows/     # CI: fmt, vet, build, test
├── Makefile
└── go.mod
```

## 🧠 Go cheat sheet

```go
x := 5                          // declare + assign
var s []int                     // nil slice
s = append(s, 1)                // grow a slice
m := map[string]int{"a": 1}     // map literal
v, ok := m["a"]                 // comma-ok
for i, v := range s {}          // loop
func f(a int) (int, error) {}   // multiple returns
type T struct{ Name string }    // struct
func (t *T) Set(n string) {}    // method (pointer receiver)
type I interface{ M() }         // interface (implicit)
go work()                       // goroutine
ch := make(chan int); ch <- 1   // channel send
defer f.Close()                 // run on return
func Map[T, U any](...) {}      // generics
```

## 🔗 Resources

- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Standard library docs](https://pkg.go.dev/std)
