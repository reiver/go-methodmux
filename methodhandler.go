package methodmux

import (
	"net/http"
)

// When something that fits MethodDeleteHandler is passed to MuxedHandler(), method handler
// will use the ServeHTTPDelete to serve the HTTP DELETE method.
//
// Example
//
//	type MyMethods struct{
//		//@TODO
//	}
//	
//	func (receiver MyMethods) ServeHTTPDelete(w http.ResponseWriter, r *http.Request) {
//		//@TODO
//	}
//	
//	// ...
//	
//	var myMethods MyMethods
//
//	var handler http.Handler = methodmux.MuxedHandler(myMethods)
type MethodDeleteHandler interface {
	ServeHTTPDelete(http.ResponseWriter, *http.Request)
}

// When something that fits MethodGetHandler is passed to MuxedHandler(), method handler
// will use the ServeHTTPGet to serve the HTTP GET method.
//
// Example
//
//	type MyMethods struct{
//		//@TODO
//	}
//	
//	func (receiver MyMethods) ServeHTTPGet(w http.ResponseWriter, r *http.Request) {
//		//@TODO
//	}
//	
//	// ...
//	
//	var myMethods MyMethods
//
//	var handler http.Handler = methodmux.MuxedHandler(myMethods)
type MethodGetHandler interface {
	ServeHTTPGet(http.ResponseWriter, *http.Request)
}

// When something that fits MethodPatchHandler is passed to MuxedHandler(), method handler
// will use the ServeHTTPPatch to serve the HTTP PATCH method.
//
// Example
//
//	type MyMethods struct{
//		//@TODO
//	}
//	
//	func (receiver MyMethods) ServeHTTPPatch(w http.ResponseWriter, r *http.Request) {
//		//@TODO
//	}
//	
//	// ...
//	
//	var myMethods MyMethods
//
//	var handler http.Handler = methodmux.MuxedHandler(myMethods)
type MethodPatchHandler interface {
	ServeHTTPPatch(http.ResponseWriter, *http.Request)
}

// When something that fits MethodPostHandler is passed to MuxedHandler(), method handler
// will use the ServeHTTPPost to serve the HTTP POST method.
//
// Example
//
//	type MyMethods struct{
//		//@TODO
//	}
//	
//	func (receiver MyMethods) ServeHTTPPost(w http.ResponseWriter, r *http.Request) {
//		//@TODO
//	}
//	
//	// ...
//	
//	var myMethods MyMethods
//
//	var handler http.Handler = methodmux.MuxedHandler(myMethods)
type MethodPostHandler interface {
	ServeHTTPPost(http.ResponseWriter, *http.Request)
}

// When something that fits MethodPutHandler is passed to MuxedHandler(), method handler
// will use the ServeHTTPPut to serve the HTTP PUT method.
//
// Example
//
//	type MyMethods struct{
//		//@TODO
//	}
//	
//	func (receiver MyMethods) ServeHTTPPut(w http.ResponseWriter, r *http.Request) {
//		//@TODO
//	}
//	
//	// ...
//	
//	var myMethods MyMethods
//
//	var handler http.Handler = methodmux.MuxedHandler(myMethods)
type MethodPutHandler interface {
	ServeHTTPPut(http.ResponseWriter, *http.Request)
}
