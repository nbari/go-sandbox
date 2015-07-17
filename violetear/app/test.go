package resource

import (
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get(":name")
	w.Write([]byte("Test Hello " + name))
}
