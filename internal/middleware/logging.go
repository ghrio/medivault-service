package middleware

import (
	"log"
	"net/http"
	"time"
)

// WrappedResponseWriter wraps the standard http.ResponseWriter to capture the status code
type WrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// NewWrappedResponseWriter initializes a WrappedResponseWriter with a default status code
func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	// Default status code to 200 in case WriteHeader is not called
	return &WrappedResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

// WriteHeader captures the status code and delegates to the original ResponseWriter
func (wrw *WrappedResponseWriter) WriteHeader(code int) {
	wrw.statusCode = code
	wrw.ResponseWriter.WriteHeader(code)
}

// Logging is a middleware that logs the HTTP request method, URL path, status code, and duration
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Initialize the wrapped response writer
		wrw := NewWrappedResponseWriter(w)

		// Call the next handler with the wrapped response writer
		next.ServeHTTP(wrw, r)

		// Log the request details
		log.Printf("%3d|%13v|%15s|%-7s%s\n",
			wrw.statusCode,
			time.Since(startTime),
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
		)
	})
}
