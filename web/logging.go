package web

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := newLoggingResponseWriter(w)

		start := time.Now()
		next.ServeHTTP(lw, r)
		stop := time.Now()

		log.Info("Request",
			zap.String("remoteAddr", r.RemoteAddr),
			zap.String("protocol", r.Proto),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("statusCode", lw.statusCode),
			zap.Duration("responseTime", stop.Sub(start)),
		)
	})
}
