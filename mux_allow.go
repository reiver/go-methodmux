package methodmux

import (
	"sort"
	"strings"
)

// allow returns the value for the HTTP response “Allow” header.
//
//
// 10.4.6 405 Method Not Allowed
//
// The method specified in the Request-Line is not allowed for the resource identified
// by the Request-URI. The response MUST include an Allow header containing a list of
// valid methods for the requested resource.
func (receiver *Mux) allow() string {
	if nil == receiver {
		return ""
	}

	var allowedMethods []string
	func(){
		receiver.mutex.RLock()
		defer receiver.mutex.RUnlock()

		for method, _ := range receiver.handlers {
			allowedMethods = append(allowedMethods, method)
		}

		allowedMethods = append(allowedMethods, "OPTIONS")
	}()
	sort.Strings(allowedMethods)

	return strings.Join(allowedMethods, ", ")
}
