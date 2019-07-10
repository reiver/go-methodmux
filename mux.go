package methodmux

import (
	"net/http"
	"sync"
)

// Mux register (a number of) handlers (i.e., http.Handler) to handle (a number of) HTTP methods.
//
// Common HTTP methods include:
//
// • DELETE
//
// • GET
//
// • PATCH
//
// • POST
//
// • PUT
//
// The methods used to register handlers (i.e., http.Handler) for each of thse HTTP request headers is:
//
// • methodmux.Mux.HandleMethodDelete()
//
// • methodmux.Mux.HandleMethodGet()
//
// • methodmux.Mux.HandleMethodPatch()
//
// • methodmux.Mux.HandleMethodPost()
//
// • methodmux.Mux.HandleMethodPut()
//
// For any other methods (including made up methods), one can use:
//
// • methodmux.Mux.HandleMethod()
type Mux struct {
	mutex sync.RWMutex
	handlers map[string]http.Handler
}
