package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/assert"
	"github.com/kirontoo/go-backend-template/internal/jsonlog"
)

func TestCheckHealthHandler(t *testing.T) {
	response := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	logger := jsonlog.New(io.Discard, jsonlog.LevelInfo)
	app := &application{
		logger: logger,
	}

	app.healthcheckHandler(response, request)

	result := response.Result()

	assert.Status(t, result.StatusCode, http.StatusOK)

	type jsonResponseBody struct {
		Status     string
		SystemInfo struct {
			environment string
			version     string
		} `json:"system_info"`
	}

	var body jsonResponseBody
	err = json.NewDecoder(response.Body).Decode(&body)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into body, '%v'", response.Body, err)
	}

	assert.Equal(t, body.Status, "available")
}
