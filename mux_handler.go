package methodmux

import (
	"net/http"
)

// Handler returns the http.Handler that would deal with the HTTP request method given by ‘method’.
func (receiver *Mux) Handler(method string) http.Handler {
	if nil == receiver {
		return nil
	}

	receiver.mutex.RLock()
	defer receiver.mutex.RUnlock()

	handler, found := receiver.handlers[method]
	if !found {
		return nil
	}

	return handler
}
