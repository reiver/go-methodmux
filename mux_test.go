package methodmux_test

import (
	"github.com/reiver/go-methodmux"

	"net/http"

	"testing"
)

func MuxTypeAsHandler(t *testing.T) {

	var handler http.Handler = new(methodmux.Mux) // THIS IS THE LINE THAT ACTUALLY MATTERS.

	if nil == handler {
		t.Error("This should never happen.")
	}
}
