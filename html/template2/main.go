package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/nbari/violetear"
)

// compile all templates and cache them
var (
	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob(
		filepath.Join("templates", "*.tpl")),
	)
	funcMap = template.FuncMap{
		"Upper": toUpper,
		"title": strings.Title,
	}
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}

type Page struct {
	Title string
	Name  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	page := Page{
		Title: "test title",
		Name:  params["*"].(string),
	}
	renderTemplate(w, "index", &page)
}

// The server itself
func main() {
	router := violetear.New()
	router.HandleFunc("/hello/*", hello)
	router.Handle("/static/*",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("./static")),
		),
	)
	log.Fatal(http.ListenAndServe(":8080", router))
}
