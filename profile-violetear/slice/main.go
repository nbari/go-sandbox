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
func (p *Params) AddSlice(name, value string) {
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

// Params string/interface map used with context
type MapParams map[string]interface{}

// Add param to Params
func (p MapParams) AddMap(k, v string) {
	if param, ok := p[k]; ok {
		switch param.(type) {
		case string:
			param = []string{param.(string), v}
		case []string:
			param = append(param.([]string), v)
		}
		p[k] = param
	} else {
		p[k] = v
	}
}

func main() {
	var params Params
	fmt.Printf("params = %+v\n", params)
	fmt.Printf("params == nil = %+v\n", params == nil)
	if params == nil {
		params = Params{}
	}
	fmt.Printf("params == nil = %+v\n", params == nil)
	params.AddSlice("foo", "bar")
	fmt.Printf("params = %+v\n", params)
	params.AddSlice("foo", "bar")
	params.AddSlice("bar", "foo")
	params.AddSlice("bar", "foo")
	params.AddSlice("*", "catch-all")
	fmt.Printf("params = %+v\n", params)

	var mparams MapParams
	if mparams == nil {
		mparams = MapParams{}
	}
	mparams.AddMap("foo", "bar")
	fmt.Printf("params = %+v\n", mparams)
	mparams.AddMap("foo", "bar")
	mparams.AddMap("bar", "foo")
	mparams.AddMap("bar", "foo")
	mparams.AddMap("*", "catch-all")
	fmt.Printf("params = %+v\n", mparams)
}
