package main

import (
  "net/http"
)

var addr string = "6969"

func main() {

  setupDB()
  setupRouter()
  http.ListenAndServe(":"+addr, nil)
}
