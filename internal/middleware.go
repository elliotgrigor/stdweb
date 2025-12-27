package internal

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Log to a file
		log.Println(r.RemoteAddr, r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isAuthenticated() && isAuthorised() {
			next.ServeHTTP(w, r)
			return
		}
		if !isAuthenticated() {
			http.Error(w, "Error 401: Unauthorized", http.StatusInternalServerError)
			return
		}
		if !isAuthorised() {
			http.Error(w, "Error 403: Forbidden", http.StatusInternalServerError)
			return
		}
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
	})
}
