package main

import (
	"fmt"
	"github.com/nbari/violetear"
)

func main() {
	router := violetear.New("config.yml")
	router.Add("/", Index, "")
	fmt.Println(router)

	//	log.Fatal(http.ListenAndServe(":8080", router))
}
