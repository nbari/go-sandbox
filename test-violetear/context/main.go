package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/nbari/violetear"
)

func catchAll(w http.ResponseWriter, r *http.Request) {
	// Get & print the content of named-param *
	fmt.Fprintf(w, "CatchAll value:, %q", r.Context().Value("*"))
}

func handleUUID(w http.ResponseWriter, r *http.Request) {
	// add a key-value pair to the context
	ctx := context.WithValue(r.Context(), "key", "my-value")
	// print current value for :uuid
	fmt.Fprintf(w, "Named parameter: %q, key: %s",
		ctx.Value(":uuid"),
		ctx.Value("key"),
	)
}

func main() {
	router := violetear.New()

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)

	router.HandleFunc("*", catchAll)
	router.HandleFunc("/:uuid", handleUUID, "GET,HEAD")

	log.Fatal(http.ListenAndServe(":8080", router))
}
