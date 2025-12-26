package main

import "net/http"

func handleRoot(w http.ResponseWriter, r *http.Request) {
	handleLogin(w, r)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated() {
		http.Redirect(w, r, "/app/dashboard", http.StatusFound)
		return
	}
	tmplRender(w, "login", nil)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "dashboard", nil)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "about", nil)
}

func handleSnippet(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "_htmx/snippet", nil)
}
