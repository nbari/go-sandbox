package main

import (
	"fmt"
	"github.com/nbari/violetear"
	"github.com/nbari/violetear/config"
)

func main() {
	config := config.Get("~/projects/go/src/github.com/nbari/violetear/config/config_test.yml")
	router := violetear.New()
	fmt.Println(config, router)

	//	log.Fatal(http.ListenAndServe(":8080", router))
}
