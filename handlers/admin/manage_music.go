package admin

import (
	"fmt"
	"net/http"
)

func ManageMusicHandler( w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello %s", "world")
}
