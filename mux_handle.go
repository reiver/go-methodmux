package methodmux

import (
	"net/http"
)

// HandleMethod registers a handler (i.e., http.Handler) to handle an HTTP request method given by ‘method’.
func (receiver *Mux) HandleMethod(handler http.Handler, method string) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == handler {
		return errNilHandler
	}

	switch method {
	case "OPTIONS":
		return errOptionsMethodNotAllowed
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	if nil == receiver.handlers {
		receiver.handlers = map[string]http.Handler{}
	}

	_, found := receiver.handlers[method]
	if found {
		return errFound
	}

	receiver.handlers[method] = handler

	return nil
}

func (receiver *Mux) HandleMethodDelete(handler http.Handler) error {
	const method string = "DELETE"

	return receiver.HandleMethod(handler, method)
}

func (receiver *Mux) HandleMethodGet(handler http.Handler) error {
	const method string = "GET"

	return receiver.HandleMethod(handler, method)
}

func (receiver *Mux) HandleMethodPatch(handler http.Handler) error {
	const method string = "PATCH"

	return receiver.HandleMethod(handler, method)
}

func (receiver *Mux) HandleMethodPost(handler http.Handler) error {
	const method string = "POST"

	return receiver.HandleMethod(handler, method)
}

func (receiver *Mux) HandleMethodPut(handler http.Handler) error {
	const method string = "PUT"

	return receiver.HandleMethod(handler, method)
}
