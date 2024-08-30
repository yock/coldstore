package scan

import (
  "embed"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "yock.dev/coldstore/layout"
)

//go:embed templates/*
var templateFiles embed.FS

var templates = map[string]*template.Template {
  "index": template.Must(layout.Layout().ParseFS(templateFiles, "templates/index.html")),
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
  err := templates["index"].Execute(response, nil)
  if err != nil {
    http.Error(response, err.Error(), http.StatusInternalServerError)
  }
}

func Router(router *mux.Router) {
  router.HandleFunc("", indexHandler).Methods("GET")
}

