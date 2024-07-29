package main

import (
  "os"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "yock.dev/coldstore/cuts"
  "yock.dev/coldstore/home"
)

func main () {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router := mux.NewRouter()
  cuts.Router(router.PathPrefix("/{cuts:cuts\\/?}").Subrouter())
  router.HandleFunc("/", home.HomeHandler)

  log.Println("Listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, router))
}
