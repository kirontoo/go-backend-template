package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/assert"
)

func TestCheckHealthWithApp(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/v1/healthcheck")

	assert.Equal(t, code, http.StatusOK)

	type healthCheckResponseBody struct {
		Status     string `json:"status"`
		SystemInfo struct {
			environment string
			version     string
		} `json:"system_info"`
	}

	var actual healthCheckResponseBody
	if err := json.Unmarshal([]byte(body), &actual); err != nil {

		t.Fatalf("Unable to parse response from server %q into body, '%v'", body, err)
	}

	assert.Equal(t, actual.Status, "available")
}
