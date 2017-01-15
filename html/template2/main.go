package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

// Model of stuff to render a page
type Model struct {
	Title string
	Name  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value(violetear.ParamsKey).(violetear.Params)
	model := Model{
		Title: "xxx",
		Name:  params["*"].(string),
	}
	renderTemplate(w, "indexPage", &model)
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
