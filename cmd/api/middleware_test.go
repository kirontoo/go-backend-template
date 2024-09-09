package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/assert"
	"github.com/kirontoo/go-backend-template/internal/data"
)

func TestRequirePermission(t *testing.T) {
	t.Run("user activated", func(t *testing.T) {
		app := newTestApplication(t)

		rr := httptest.NewRecorder()

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		user := &data.User{
			ID:        1,
			Activated: true,
		}

		// mock context
		ctx := context.WithValue(r.Context(), userContextKey, user)
		r = r.WithContext(ctx)

		// mock HTTP Handler to pass to RequirePermissions middleware
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusOK)

	})

	t.Run("user not activated", func(t *testing.T) {
		app := newTestApplication(t)

		rr := httptest.NewRecorder()

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		user := &data.User{
			ID:        1,
			Activated: false,
		}

		// mock context
		ctx := context.WithValue(r.Context(), userContextKey, user)
		r = r.WithContext(ctx)

		// mock HTTP Handler to pass to RequirePermissions middleware
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusForbidden)
	})

	t.Run("user does not have permission", func(t *testing.T) {
		app := newTestApplication(t)

		rr := httptest.NewRecorder()

		r, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		user := &data.User{
			ID:        3,
			Activated: false,
		}

		// mock context
		ctx := context.WithValue(r.Context(), userContextKey, user)
		r = r.WithContext(ctx)

		// mock HTTP Handler to pass to RequirePermissions middleware
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusForbidden)
	})
}
