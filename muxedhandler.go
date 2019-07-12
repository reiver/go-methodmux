package methodmux

import (
	"net/http"
)

// MuxedHandler returns an http.Handler where the HTTP methods (i.e., DELETE, GET, PATCH, POST, PUT)
// are handler by specific methodss of ‘methodHandler’ (if they exist): ServeHTTPDelete(w,r),
// ServeHTTPGet(w, r), ServeHTTPPatch(w, r), ServeHTTPPost(w, r), ServeHTTPPut(w, r)
func MuxedHandler(methodHandler interface{}) http.Handler {
	var muxedHandler internalMuxedHandler

	// DELETE
	if h, casted := methodHandler.(MethodDeleteHandler); casted {
		var handler http.Handler = http.HandlerFunc(h.ServeHTTPDelete)
		if err := muxedHandler.mux.HandleMethodDelete(handler); nil != err {
			return nil
		}
	}

	// GET
	if h, casted := methodHandler.(MethodGetHandler); casted {
		var handler http.Handler = http.HandlerFunc(h.ServeHTTPGet)
		if err := muxedHandler.mux.HandleMethodGet(handler); nil != err {
			return nil
		}
	}

	// PATCH
	if h, casted := methodHandler.(MethodPatchHandler); casted {
		var handler http.Handler = http.HandlerFunc(h.ServeHTTPPatch)
		if err := muxedHandler.mux.HandleMethodPatch(handler); nil != err {
			return nil
		}
	}

	// POST
	if h, casted := methodHandler.(MethodPostHandler); casted {
		var handler http.Handler = http.HandlerFunc(h.ServeHTTPPost)
		if err := muxedHandler.mux.HandleMethodPost(handler); nil != err {
			return nil
		}
	}

	// PUT
	if h, casted := methodHandler.(MethodPutHandler); casted {
		var handler http.Handler = http.HandlerFunc(h.ServeHTTPPut)
		if err := muxedHandler.mux.HandleMethodPut(handler); nil != err {
			return nil
		}
	}

	return &muxedHandler
}

type internalMuxedHandler struct {
	mux Mux
}

func (receiver *internalMuxedHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == receiver {
		http.Error(responseWriter, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	receiver.mux.ServeHTTP(responseWriter, request)
	return
}
