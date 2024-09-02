package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/assert"
)

func TestCheckHealthWithApp(t *testing.T) {
	app := newTestApplication(t)

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/v1/healthcheck")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	type jsonResponseBody struct {
		Status     string
		SystemInfo struct {
			environment string
			version     string
		} `json:"system_info"`
	}

	var body jsonResponseBody
	err = json.NewDecoder(rs.Body).Decode(&body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into body, '%v'", rs.Body, err)
	}

	assert.Equal(t, body.Status, "available")
}
