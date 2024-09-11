package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/data"
	"github.com/kirontoo/go-backend-template/internal/data/mocks"
	"github.com/kirontoo/go-backend-template/internal/jsonlog"
)

func newTestApplication(t testing.TB) *application {
	t.Helper()

	logger := jsonlog.New(io.Discard, jsonlog.LevelInfo)
	return &application{
		logger: logger,
		models: data.Models{
			Movies:      mocks.MovieModel{},
			Permissions: mocks.PermisionModel{},
			Tokens:      mocks.TokenModel{},
			Users:       mocks.UserModel{},
		},
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)

	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// store any response cookies
	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// Makes a GET request to a given url path using the test server client, and returns the
// response status code, headers and body.
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func setupRecorder(t testing.TB) (*httptest.ResponseRecorder, *http.Request) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	return rr, r
}
