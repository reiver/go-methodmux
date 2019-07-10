package methodmux_test

import (
	"github.com/reiver/go-methodmux"

	"net/http"
	"net/http/httptest"

	"testing"
)

func TestMuxServerHTTPMethodNotAllowed(t *testing.T) {

	const headerName string = "Allow"

	var mux methodmux.Mux

	{
		var httpResponseRecorder httptest.ResponseRecorder
		var httpRequest *http.Request = httptest.NewRequest("PUNCH", "/path/to/file.html", nil)
		mux.ServeHTTP(&httpResponseRecorder, httpRequest)

		if expected, actual := http.StatusMethodNotAllowed, httpResponseRecorder.Code; expected != actual {
			t.Error("Did not get expected HTTP response code.")
			t.Logf("EXPECTED HTTP response code: %d", expected)
			t.Logf("ACTUAL   HTTP response code: %d", actual)
			return
		}

		if expected, actual := "OPTIONS", httpResponseRecorder.HeaderMap.Get(headerName); expected != actual {
			t.Errorf("For HTTP header %q, did not get expected value.", headerName)
			t.Logf("EXPECTED HTTP %s header value: %q", headerName, expected)
			t.Logf("ACTUAL   HTTP %s header value: %q", headerName, actual)
		}
	}

	{
		var handler http.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request){})

		if err := mux.HandleMethod(handler, "GET"); nil != err {
			t.Errorf("Did not expect to get an error, but actually got one: (%T) %q", err, err)
			return
		}
	}
	{
		var httpResponseRecorder httptest.ResponseRecorder
		var httpRequest *http.Request = httptest.NewRequest("PUNCH", "/path/to/file.html", nil)
		mux.ServeHTTP(&httpResponseRecorder, httpRequest)

		if expected, actual := http.StatusMethodNotAllowed, httpResponseRecorder.Code; expected != actual {
			t.Error("Did not get expected HTTP response code.")
			t.Logf("EXPECTED HTTP response code: %d", expected)
			t.Logf("ACTUAL   HTTP response code: %d", actual)
			return
		}

		if expected, actual := "GET, OPTIONS", httpResponseRecorder.HeaderMap.Get(headerName); expected != actual {
			t.Errorf("For HTTP header %q, did not get expected value.", headerName)
			t.Logf("EXPECTED HTTP %s header value: %q", headerName, expected)
			t.Logf("ACTUAL   HTTP %s header value: %q", headerName, actual)
		}
	}


	{
		var handler http.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request){})

		if err := mux.HandleMethod(handler, "POST"); nil != err {
			t.Errorf("Did not expect to get an error, but actually got one: (%T) %q", err, err)
			return
		}
	}
	{
		var httpResponseRecorder httptest.ResponseRecorder
		var httpRequest *http.Request = httptest.NewRequest("PUNCH", "/path/to/file.html", nil)
		mux.ServeHTTP(&httpResponseRecorder, httpRequest)

		if expected, actual := http.StatusMethodNotAllowed, httpResponseRecorder.Code; expected != actual {
			t.Error("Did not get expected HTTP response code.")
			t.Logf("EXPECTED HTTP response code: %d", expected)
			t.Logf("ACTUAL   HTTP response code: %d", actual)
			return
		}

		if expected, actual := "GET, OPTIONS, POST", httpResponseRecorder.HeaderMap.Get(headerName); expected != actual {
			t.Errorf("For HTTP header %q, did not get expected value.", headerName)
			t.Logf("EXPECTED HTTP %s header value: %q", headerName, expected)
			t.Logf("ACTUAL   HTTP %s header value: %q", headerName, actual)
		}
	}


	{
		var handler http.Handler = http.HandlerFunc(func(http.ResponseWriter, *http.Request){})

		if err := mux.HandleMethod(handler, "SCREAM"); nil != err {
			t.Errorf("Did not expect to get an error, but actually got one: (%T) %q", err, err)
			return
		}
	}
	{
		var httpResponseRecorder httptest.ResponseRecorder
		var httpRequest *http.Request = httptest.NewRequest("PUNCH", "/path/to/file.html", nil)
		mux.ServeHTTP(&httpResponseRecorder, httpRequest)

		if expected, actual := http.StatusMethodNotAllowed, httpResponseRecorder.Code; expected != actual {
			t.Error("Did not get expected HTTP response code.")
			t.Logf("EXPECTED HTTP response code: %d", expected)
			t.Logf("ACTUAL   HTTP response code: %d", actual)
			return
		}

		if expected, actual := "GET, OPTIONS, POST, SCREAM", httpResponseRecorder.HeaderMap.Get(headerName); expected != actual {
			t.Errorf("For HTTP header %q, did not get expected value.", headerName)
			t.Logf("EXPECTED HTTP %s header value: %q", headerName, expected)
			t.Logf("ACTUAL   HTTP %s header value: %q", headerName, actual)
		}
	}
}
