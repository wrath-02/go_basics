# 11 · Building an API

> **Definition:** A **REST API** is an HTTP server that exposes resources at URLs (endpoints) and exchanges data, usually as JSON. Go's standard library `net/http` is production-ready, and since Go 1.22 its router supports HTTP methods and path wildcards, so this project needs **zero third-party dependencies**.

This project is a small **coin-balance service**. A user sends their username and an auth token, and the server returns their balance.

## Run it

```bash
go run ./11-building-an-api/cmd/api
```

In another terminal:

```bash
curl localhost:8000/health
curl -H "Authorization: 123ABC" "localhost:8000/account/coins?username=alex"
curl -H "Authorization: WRONG"  "localhost:8000/account/coins?username=alex"
curl -X POST "localhost:8000/account/coins?username=alex"
```

```text
{"status":"ok"}
{"code":200,"balance":100,"username":"alex"}
{"code":401,"message":"invalid username or token"}
Method Not Allowed
```

Server logs (structured logging with `log/slog`):

```text
level=INFO msg="starting GO API service" addr=:8000
level=INFO msg=request method=GET path=/account/coins status=200 duration=22.199ms
level=INFO msg=request method=GET path=/account/coins status=401 duration=11.057ms
level=INFO msg=request method=POST path=/account/coins status=405 duration=40µs
level=INFO msg="shutting down"          # after Ctrl+C (graceful shutdown)
```

Test users (in the mock database):

| username | token | coins |
|---|---|---|
| alex | `123ABC` | 100 |
| jason | `456DEF` | 200 |
| marie | `789GHI` | 300 |

## Project structure

```text
11-building-an-api/
├── cmd/api/main.go                  # entry point: config, server, graceful shutdown
├── api/api.go                       # shared request/response types + JSON error helpers
└── internal/                        # private to this module (can't be imported from outside)
    ├── handlers/
    │   ├── handlers.go              # router: URL pattern -> handler
    │   ├── get_balance.go           # GET /account/coins
    │   └── handlers_test.go         # tests with net/http/httptest
    ├── middleware/
    │   ├── authorization.go         # checks the Authorization header
    │   └── logger.go                # logs every request
    └── tools/
        ├── database.go              # DatabaseInterface (the abstraction)
        └── mockdb.go                # in-memory implementation
```

This follows the common Go project layout: `cmd/` for executables, `internal/` for private packages, and a public `api/` package for shared types.

## Request flow

```text
request ─► Logger ─► ServeMux ─► Authorization ─► GetCoinBalance ─► DatabaseInterface
                                     │ 401 if bad token        │ 400 / 404 / 200 JSON
```

## Key ideas

### 1. Routing with the standard library

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /health", healthHandler)
mux.Handle("GET /account/coins", auth(GetCoinBalance(db)))
// Wildcards: mux.HandleFunc("GET /users/{id}", ...) then r.PathValue("id")
```

### 2. Middleware = `func(http.Handler) http.Handler`

```go
func Authorization(db tools.DatabaseInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ...check token...
			next.ServeHTTP(w, r) // continue the chain
		})
	}
}
```

### 3. Depend on interfaces (dependency injection)

Handlers receive a `tools.DatabaseInterface`, not a concrete database. To use a real database (Postgres, for example), write a type with the same methods. Nothing else has to change.

### 4. JSON in and out

```go
type CoinBalanceResponse struct {
	Code    int    `json:"code"`     // struct tags control JSON field names
	Balance int64  `json:"balance"`
	User    string `json:"username"`
}
json.NewEncoder(w).Encode(resp)
```

### 5. Graceful shutdown

```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
<-ctx.Done()        // wait for Ctrl+C
srv.Shutdown(ctx)   // finish in-flight requests, then stop
```

## Tests

```bash
go test -v ./11-building-an-api/...
```

```text
--- PASS: TestGetCoinBalance (0.04s)
    --- PASS: TestGetCoinBalance/valid_user_and_token
    --- PASS: TestGetCoinBalance/wrong_token
    --- PASS: TestGetCoinBalance/missing_token
    --- PASS: TestGetCoinBalance/unknown_user
    --- PASS: TestGetCoinBalance/missing_username
    --- PASS: TestGetCoinBalance/wrong_method
--- PASS: TestHealth
--- PASS: TestGetCoinBalanceMissingUsername
```

## Try it yourself

1. Add `POST /account/coins/add?username=...&amount=...` to increase a balance (protect the map with a `sync.Mutex`).
2. Add `GET /users/{username}` using `r.PathValue("username")`.
3. Popular third-party alternatives for bigger projects: `chi` (router), `gorilla/schema` (query decoding), `logrus` / `zap` (logging).

---
[← Generics](../10-generics/) · [🏠 Index](../README.md) · [Next (bonus): Error Handling →](../12-error-handling/)
