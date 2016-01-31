package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, text string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			text,
		)
	})
}
