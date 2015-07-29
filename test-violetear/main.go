package main

import (
	"fmt"
	"github.com/nbari/violetear"
)

func main() {
	config := violetear.New("config.yml")
	fmt.Println(config)

	//	log.Fatal(http.ListenAndServe(":8080", router))
}
