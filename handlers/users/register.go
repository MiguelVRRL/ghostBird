package users

import (
	"fmt"
	"net/http"
)

func RegisterHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", "name")
}

