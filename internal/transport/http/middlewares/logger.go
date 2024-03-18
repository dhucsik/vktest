package middlewares

import (
	"log"
	"net/http"
)

type LoggerMiddleware struct{}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (m *LoggerMiddleware) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		next(&sw, r)
		log.Printf("%s %s %d", r.Method, r.URL.Path, sw.status)
	}
}
