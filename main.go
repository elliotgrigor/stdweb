package main

import (
	"fmt"
	"log"
	"net/http"
)

const devMode = true

const (
	host = "127.0.0.1"
	port = 3000
)

func main() {
	public := http.NewServeMux()
	{
		public.HandleFunc("GET /", handleRoot)
		public.HandleFunc("GET /login", handleLogin)
	}

	protected := http.NewServeMux()
	{
		protected.HandleFunc("GET /dashboard", handleDashboard)
		protected.HandleFunc("GET /about", handleAbout)
		protected.HandleFunc("GET /_htmx/snippet", handleSnippet)
	}
	authProtected := authMiddleware(protected)

	app := http.NewServeMux()
	app.Handle("/", public)
	app.Handle("/app/", http.StripPrefix("/app", authProtected))

	log.Printf("Listening on %s:%d\n", host, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), loggerMiddleware(app))
}
