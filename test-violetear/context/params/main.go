package main

import "fmt"

type Params map[string]interface{}

func main() {

	params := make(Params)

	params["a"] = "aa"
	params["z"] = []string{"a", "a"}
	fmt.Printf("params = %+v\n", params)

	x := params["z"].([]string)
	fmt.Printf("x = %+v\n", x[0])
	fmt.Printf("x = %+v\n", x[1])
}
