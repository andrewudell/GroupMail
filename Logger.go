package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, text string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			text,
			time.Since(start),
		)
	})
}