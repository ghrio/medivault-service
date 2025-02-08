package middleware

import (
	"log"
	"net/http"
	"time"
)

type WrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeNow := time.Now()

		wrw := &WrappedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrw, r)
		log.Println(wrw.statusCode, r.Method, r.URL.Path, time.Since(timeNow))
	})
}
