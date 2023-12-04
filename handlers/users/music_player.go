package users

import (
	"fmt"
	"net/http"
)

func MusicPlayerHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", "name")
}

