package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elliotgrigor/stdweb/internal"
)

const (
	host = "127.0.0.1"
	port = 3000
)

func main() {
	public := http.NewServeMux()
	{
		public.HandleFunc("GET /", internal.HandleRoot)
		public.HandleFunc("GET /login", internal.HandleLogin)
	}

	protected := http.NewServeMux()
	{
		protected.HandleFunc("GET /dashboard", internal.HandleDashboard)
		protected.HandleFunc("GET /about", internal.HandleAbout)
		protected.HandleFunc("GET /_htmx/snippet", internal.HandleSnippet)
	}
	authProtected := internal.AuthMiddleware(protected)

	app := http.NewServeMux()
	app.Handle("/", public)
	app.Handle("/app/", http.StripPrefix("/app", authProtected))

	log.Printf("Listening on %s:%d\n", host, port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), internal.LoggerMiddleware(app))
}
