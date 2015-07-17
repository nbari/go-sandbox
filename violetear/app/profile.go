package resource

import (
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get(":name")
	w.Write([]byte("Hello " + name))
}
