package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// RequestLog is a middleware function that will produce log of incoming request
func RequestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(rw, r)
		log.WithFields(log.Fields{
			"start_time": startTime,
			"end_time":   time.Now(),
			"method":     r.Method,
			"url":        r.URL.RequestURI(),
		}).Info("successfully serve incoming request")
	})
}
