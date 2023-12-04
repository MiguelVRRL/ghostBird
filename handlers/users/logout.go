package users

import (
	"fmt"
	"net/http"
)

func LogoutHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", "name")
}

