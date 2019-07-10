# go-methodmux

Package **methodmux** provides a HTTP request method oriented ("middleware") HTTP handler, which can hand-off to other HTTP handler for each HTTP request method,
for the Go programming language's built-in `"net/http"` library.

Package **methodmux** is a HTTP request router, and dispatcher.

The name ‘mux’ stands for "HTTP request multiplexer".


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-methodmux

[![GoDoc](https://godoc.org/github.com/reiver/go-methodmux?status.svg)](https://godoc.org/github.com/reiver/go-methodmux)

## Example
```go
import "github.com/reiver/go-methodmux"

// ...

var mux methodmux.Mux

err := mux.HandleGET(getHandler)
err := mux.HandlePATCH(patchHandler)
err := mux.HandlePUT(putHandler)

err := mux.Handle(screamHandler, "SCREAM")

// ...

err := http.ListenAndServe(":8080", mux)
```
