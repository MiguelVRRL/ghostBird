package users

import (
	"fmt"
	"net/http"
)

func HomeHandler (w http.ResponseWriter, r *http.Request) {
  uri := r.URL.Path[1:]
  fmt.Fprintf(w, "Hello, %s", uri)
}
