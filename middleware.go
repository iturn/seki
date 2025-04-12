package seki

import (
	"context"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](Chain(f, m[1:cap(m)]...))
}

func ContextMiddleware(s *Seki) Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ctx = context.WithValue(ctx, "user", "testUser1") // temp fixed key value

			newRequest := r.WithContext(ctx)

			nextHandler.ServeHTTP(w, newRequest)
		})
	}
}

func ApiKeyHeader(s *Seki) Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKeyHeader := r.Header.Get("x-api-key")

			if apiKeyHeader != "1234" { // temp fixed api key
				s.Log.Warn("ApiKeyHeader middleware failed")
				s.UnauthorizedResponse(w)
				return
			}

			nextHandler.ServeHTTP(w, r)
		})
	}
}

func RequestLogger(s *Seki) Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ew := newExtendResponseWriter(w)

			nextHandler.ServeHTTP(ew, r)

			s.Log.Info("request", "method", r.Method, "status", ew.StatusCode, "ms", int(time.Since(start).Milliseconds()), "url", r.URL.RequestURI())
		})
	}
}

// used in request logger middleware
type extendedResponseWriter struct {
	responseWriter http.ResponseWriter
	StatusCode     int
}

func newExtendResponseWriter(w http.ResponseWriter) *extendedResponseWriter {
	return &extendedResponseWriter{w, http.StatusOK}
}

func (w *extendedResponseWriter) Write(b []byte) (int, error) {
	return w.responseWriter.Write(b)
}

func (w *extendedResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *extendedResponseWriter) WriteHeader(statusCode int) {
	// receive status code from this method
	w.StatusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)
}
