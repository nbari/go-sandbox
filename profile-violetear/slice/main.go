package main

import "fmt"

type (
	param struct {
		name  string
		value interface{}
	}

	Params []param
)

// Add param to Params
func (p *Params) Add(name, value string) {
	for i, param := range *p {
		if param.name == name {
			switch param.value.(type) {
			case string:
				param.value = []string{param.value.(string), value}
			case []string:
				param.value = append(param.value.([]string), value)
			}
			(*p)[i] = param
			return
		}
	}
	*p = append(*p, param{name, value})
}

func main() {
	var params Params
	fmt.Printf("params = %+v\n", params)
	fmt.Printf("params == nil = %+v\n", params == nil)
	if params == nil {
		params = Params{}
	}
	fmt.Printf("params == nil = %+v\n", params == nil)
	params.Add("foo", "bar")
	fmt.Printf("params = %+v\n", params)
	params.Add("foo", "bar")
	params.Add("bar", "foo")
	params.Add("bar", "foo")
	params.Add("*", "catch-all")
	fmt.Printf("params = %+v\n", params)
}
