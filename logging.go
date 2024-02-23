package main

import (
	"net/http"
	"time"
)

func (db *tempDB) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		db.logger.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Str("remote", r.RemoteAddr).
			Dur("duration", duration).
			Msg("request handled")
	})
}

func (db *tempDB) logError(r *http.Request, statusCode int, errMsg string, err error) {
	event := db.logger.Error().
		Str("ip", r.RemoteAddr).
		Str("method", r.Method).
		Str("url", r.URL.Path).
		Int("status", statusCode)

	if err != nil {
		event.Err(err)
	}

	if errMsg != "" {
		event.Str("error", errMsg)
	}

	event.Msg("request failed")
}
