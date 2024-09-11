package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/kirontoo/go-backend-template/internal/assert"
	"github.com/kirontoo/go-backend-template/internal/data"
)

func setupMockRouteWithUserContext(
	t testing.TB,
	r *http.Request,
	user *data.User,
) (http.HandlerFunc, *http.Request) {
	t.Helper()

	ctx := context.WithValue(r.Context(), userContextKey, user)
	r = r.WithContext(ctx)

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return next, r
}

func TestRequirePermission(t *testing.T) {
	t.Run("user activated", func(t *testing.T) {
		app := newTestApplication(t)

		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        1,
			Activated: true,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusOK)

	})

	t.Run("user not activated", func(t *testing.T) {
		app := newTestApplication(t)

		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        1,
			Activated: false,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusForbidden)
	})

	t.Run("user does not have permission", func(t *testing.T) {
		app := newTestApplication(t)

		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        3,
			Activated: false,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requirePermission("movies:read", next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusForbidden)
	})
}

func TestAuthenticatedUser(t *testing.T) {
	t.Run("should send 401 if user is not authenticated", func(t *testing.T) {
		app := newTestApplication(t)

		rr, r := setupRecorder(t)

		user := data.AnonymousUser

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requireAuthenticatedUser(next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusUnauthorized)
	})

	t.Run("should send 200 if user is authenticated", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        3,
			Activated: false,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requireAuthenticatedUser(next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusOK)
	})
}

func TestRequireActivatedUser(t *testing.T) {
	t.Run("user is activated", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        1,
			Activated: true,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requireActivatedUser(next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusOK)
	})

	t.Run("user is not activated", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		user := &data.User{
			ID:        3,
			Activated: false,
		}

		next, r := setupMockRouteWithUserContext(t, r, user)

		app.requireActivatedUser(next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.Equal(t, rs.StatusCode, http.StatusForbidden)
	})
}

func TestEnableCORS(t *testing.T) {
	t.Run("sets the 'Vary' header", func(t *testing.T) {

		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.enableCORS(next).ServeHTTP(rr, r)

		rs := rr.Result()

		expectedVaryValues := []string{"Origin", "Access-Control-Allow-Origin"}
		assert.HeaderValuesEquals(t, rs, "Vary", expectedVaryValues)
	})

	t.Run("should set Access-Control-Allow-Origin header on trusted origin", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		// set origin header
		r.Header.Set("Origin", "http://localhost:9000")
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.enableCORS(next).ServeHTTP(rr, r)

		rs := rr.Result()

		expectedVaryValues := []string{"Origin", "Access-Control-Allow-Origin"}
		assert.HeaderValuesEquals(t, rs, "Vary", expectedVaryValues)

		expectedOrigin := "http://localhost:9000"
		assert.HeaderEquals(t, rs, "Access-Control-Allow-Origin", expectedOrigin)
	})

	t.Run("should not set Access-Control-Allow-Origin header on untrusted Origin", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		// set origin header
		r.Header.Set("Origin", "http://localhost:4000")
		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.enableCORS(next).ServeHTTP(rr, r)

		rs := rr.Result()

		expectedOrigin := "http://localhost:4000"
		assert.HeaderNotEqual(t, rs, "Access-Control-Allow-Origin", expectedOrigin)
		assert.Equal(t, rs.StatusCode, http.StatusOK)
	})

	t.Run("should set headers for pre-flight request on trusted origin", func(t *testing.T) {
		app := newTestApplication(t)
		rr, r := setupRecorder(t)

		r.Method = http.MethodOptions
		r.Header.Set("Origin", "http://localhost:9000")
		r.Header.Set("OPTIONS", "GET")
		r.Header.Set("Access-Control-Request-Method", "PUT")

		next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})

		app.enableCORS(next).ServeHTTP(rr, r)

		rs := rr.Result()

		assert.HeaderEquals(t, rs, "Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
		assert.HeaderEquals(t, rs, "Access-Control-Allow-Headers", "Authorization, Content-Type")
	})
}
