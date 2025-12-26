package main

import (
	"log"
	"net/http"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Log to a file
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthenticated() && isAuthorised() {
			next.ServeHTTP(w, r)
			return
		}
		if !isAuthenticated() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error 401: Unauthorized"))
			return
		}
		if !isAuthorised() {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Error 403: Forbidden"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500: Internal Server Error"))
	})
}
