package methodmux

import (
	"fmt"
	"net/http"
)

// ServeHTTP makes &methodmux.Mux fit the http.Handler interface, and makes methodmux.Mux a type of "middleware".
func (receiver *Mux) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == receiver {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if nil == request {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var method string = request.Method

	var handler http.Handler = receiver.Handler(method)
	if nil == handler {
		// 10.4.6 405 Method Not Allowed
		//
		// The method specified in the Request-Line is not allowed for the resource identified
		// by the Request-URI. The response MUST include an Allow header containing a list of
		// valid methods for the requested resource.
		responseWriter.Header().Set("Allow", receiver.allow())

		http.Error(responseWriter, fmt.Sprintf("Method %q Not Allowed", method), http.StatusMethodNotAllowed)
		return
	}

	handler.ServeHTTP(responseWriter, request)
	return
}
