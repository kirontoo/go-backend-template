package assert

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Equal[T comparable](t testing.TB, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want:%v", actual, expected)
	}
}

func StringContains(t testing.TB, actual, expectedSubstring string) {
	t.Helper()
	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}

func Status(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("wrong status code; got %d; want %d", actual, expected)
	}
}

// JSONEqual compares the JSON from two Readers.
func JSONEqual(t testing.TB, actual, expected io.Reader) {
	t.Helper()

	var jsonA, jsonB interface{}

	decoded := json.NewDecoder(actual)
	if err := decoded.Decode(&jsonA); err != nil {
		t.Errorf("could not unmarshall json: %+v", actual)
	}

	decoded = json.NewDecoder(expected)
	if err := decoded.Decode(&jsonB); err != nil {
		t.Errorf("could not unmarshall json: %+v", expected)
	}

	eq := reflect.DeepEqual(jsonB, jsonA)
	if !eq {
		t.Errorf("got: %+v; want: %+v", actual, expected)
	}
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(t testing.TB, actual, expected []byte) {
	t.Helper()

	var jsonA, jsonB interface{}

	if err := json.Unmarshal(actual, &jsonA); err != nil {
		t.Errorf("could not unmarshall json: %+v", actual)
	}

	if err := json.Unmarshal(expected, &jsonB); err != nil {
		t.Errorf("could not unmarshall json: %+v", expected)
	}

	eq := reflect.DeepEqual(jsonB, jsonA)
	if !eq {
		t.Errorf("got: %+v; want: %+v", actual, expected)
	}
}

func HeaderEquals(t testing.TB, response *http.Response, header string, expected string) {
	t.Helper()

	actual := response.Header.Get(header)
	if actual != expected {
		t.Errorf("Header '%v' got: %v, want: %v", header, actual, expected)
	}
}
