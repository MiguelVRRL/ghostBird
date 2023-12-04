package admin

import (
	"fmt"
	"net/http"
)

func ManageUsersHandler( w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hello %s", "world")
}
