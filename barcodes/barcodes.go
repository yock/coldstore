package barcodes

import (
  "embed"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "yock.dev/coldstore/layout"
  "yock.dev/coldstore/data"
)

//go:embed templates/*
var templateFiles embed.FS

type IndexModel struct {
  Barcodes []data.Barcode
}

var templates = map[string]*template.Template {
  "index": template.Must(layout.Layout().ParseFS(templateFiles, "templates/index.html")),
}

func indexHandler(response http.ResponseWriter, request *http.Request) {
  var barcodes []data.Barcode
  data.Conn.Find(&barcodes)
  model := IndexModel {
    Barcodes: barcodes,
  }
  err := templates["index"].Execute(response, model)
  if err != nil {
    http.Error(response, err.Error(), http.StatusInternalServerError)
  }
}

func Router(router *mux.Router) {
  router.HandleFunc("", indexHandler).Methods("GET")
}
