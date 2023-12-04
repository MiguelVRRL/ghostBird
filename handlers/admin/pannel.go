package admin

import (
	"fmt"
	"net/http"
)

func PannelHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s","world")
}
