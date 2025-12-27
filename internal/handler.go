package internal

import "net/http"

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if isAuthenticated() {
		http.Redirect(w, r, "/app/dashboard", http.StatusFound)
		return
	}
	tmplRender(w, "login", nil)
}

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "dashboard", nil)
}

func HandleAbout(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "about", nil)
}

func HandleSnippet(w http.ResponseWriter, r *http.Request) {
	tmplRender(w, "_htmx/snippet", nil)
}
