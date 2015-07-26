package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"github.com/nbari/violetear/config"
)

func main() {
	//	config := config.Get("/Volumes/Raid/Google Drive/projects/go/src/github.com/nbari/violetear/router.yml")
	config := config.Get("~/projects/go/src/github.com/nbari/violetear/router.yml")
	router := violetear.New()
	fmt.Println(config, router)

	//	log.Fatal(http.ListenAndServe(":8080", router))
}
