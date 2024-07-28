package cuts

import (
  "net/http"
  "github.com/gorilla/mux"
)

func indexHandler(response http.ResponseWriter, request *http.Request) {}

func Router(router *mux.Router) {
  router.HandleFunc("/", indexHandler).Methods("GET")
}
