package internal

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	templateFileExtension = ".tmpl"
	templateLayoutPath    = "templates/layout/"
	templatePagesPath     = "templates/pages/"
	templateHtmxPath      = "templates/_htmx/"
	templateHtmxStrip     = "templates/"
)

var templateCache = map[string]*template.Template{}

func createCacheKey(fullPath, templateDirPath string) string {
	key := strings.TrimPrefix(fullPath, templateDirPath)
	key = strings.TrimSuffix(key, templateFileExtension)
	return key
}

func cacheBuilder() {
	pages, err := filepath.Glob(templatePagesPath + "*" + templateFileExtension)
	if err != nil {
		log.Fatalln(err)
	}
	for _, page := range pages {
		key := createCacheKey(page, templatePagesPath)
		templateCache[key] = template.Must(template.ParseFiles(page,
			templateLayoutPath+"base.tmpl",
			templateLayoutPath+"header.tmpl",
		))
	}
	partials, err := filepath.Glob(templateHtmxPath + "*" + templateFileExtension)
	if err != nil {
		log.Fatalln(err)
	}
	for _, partial := range partials {
		// Keep the `_htmx/` prefix
		key := createCacheKey(partial, templateHtmxStrip)
		templateCache[key] = template.Must(template.ParseFiles(partial))
	}
}

func tmplRender(w http.ResponseWriter, tmplName string, data any) {
	if devMode {
		cacheBuilder() // DEV: Hot template reloading in development mode
	}
	err := templateCache[tmplName].Execute(w, data)
	if err != nil {
		http.Error(w, "Error 500: Internal Server Error", http.StatusInternalServerError)
		log.Println("Error: Failed to build template cache:", err.Error())
		return
	}
}

func init() {
	cacheBuilder()
}
