package main

import (
	_ "fmt"
	"github.com/nbari/violetear"
)

func main() {
	router := violetear.New()

	router.AddRegex(":uuid", `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	router.AddRegex(":ip", `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

	router.Add("/hello/world", "hello_world")
	router.Add("/hello/world/get", "hello_world", "GET")
	router.Add("/", "main")
	router.Add(":uuid", "h_uuid")
	router.Add("/ip/:ip", "h_uuid")

	//	log.Fatal("teste", router)
}
