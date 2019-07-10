package methodmux

import (
	"net/http"

	"testing"
)

func TestMuxAllow(t *testing.T) {

	var mux Mux

	if expected, actual := "OPTIONS", mux.allow(); expected != actual {
		t.Error("The value for the HTTP response Allow header is not what was expected.")
		t.Logf("EXPECTED: %q", expected)
		t.Logf("ALLOW:    %q", actual)
	}

	tests := []struct{
		Method   string
		Expected string
	}{
		{
			Method:   "GET",
			Expected: "GET, OPTIONS",
		},
		{
			Method:                 "POST",
			Expected: "GET, OPTIONS, POST",
		},
		{
			Method:                       "PUT",
			Expected: "GET, OPTIONS, POST, PUT",
		},
		{
			Method:                 "PATCH",
			Expected: "GET, OPTIONS, PATCH, POST, PUT",
		},
		{
			Method:   "DELETE",
			Expected: "DELETE, GET, OPTIONS, PATCH, POST, PUT",
		},
		{
			Method:                                "POP",
			Expected: "DELETE, GET, OPTIONS, PATCH, POP, POST, PUT",
		},
		{
			Method:                                           "PUSH",
			Expected: "DELETE, GET, OPTIONS, PATCH, POP, POST, PUSH, PUT",
		},
		{
			Method:                                                      "SCREAM",
			Expected: "DELETE, GET, OPTIONS, PATCH, POP, POST, PUSH, PUT, SCREAM",
		},
		{
			Method:   "ACQUIRE",
			Expected: "ACQUIRE, DELETE, GET, OPTIONS, PATCH, POP, POST, PUSH, PUT, SCREAM",
		},
	}

	for testNumber, test := range tests {
		var handler http.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request){})

		if err := mux.HandleMethod(handler, test.Method); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected, actual := test.Expected, mux.allow(); expected != actual {
			t.Errorf("For test #%d, the value for the HTTP response Allow header is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ALLOW:    %q", actual)
			continue
		}
	}
}
