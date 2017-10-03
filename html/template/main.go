package main

//go:generate go-bindata-assetfs static/... templates/...

import (
	"html/template"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

// Model of stuff to render a page
type Model struct {
	Title string
	Name  string
}

// Templates with functions available to them
var templates = template.New("").Funcs(templateMap)

// Parse all of the bindata templates
func init() {
	for _, path := range AssetNames() {
		bytes, err := Asset(path)
		if err != nil {
			log.Panicf("Unable to parse: path=%s, err=%s", path, err)
		}
		templates.New(path).Parse(string(bytes))
	}
}

// Render a template given a model
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Well hello there
func hello(w http.ResponseWriter, r *http.Request) {
	model := Model{Name: violetear.GetParam("*", r)}
	//	model := violetear.GetParam("*", r)
	renderTemplate(w, "templates/hello.html", &model)
}

// The server itself
func main() {
	// mux handler
	router := violetear.New()

	// Example route that takes one rest style option
	router.HandleFunc("/hello/*", hello)
	router.Handle("/static/*",
		//		http.StripPrefix("/static", http.FileServer(http.Dir("./static"))),
		http.StripPrefix("/static", http.FileServer(assetFS())),
	)

	log.Fatal(http.ListenAndServe(":8080", router))
}
