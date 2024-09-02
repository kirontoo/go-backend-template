package main

import (
	"io"
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
