package home

import (
  "embed"
  "net/http"
  "html/template"
  "yock.dev/coldstore/layout"
)

//go:embed templates/*.html
var templateFiles embed.FS

var indexTemplate *template.Template = template.Must(layout.Layout().ParseFS(templateFiles, "templates/index.html"))

func HomeHandler(response http.ResponseWriter, request *http.Request) {
  err := indexTemplate.Execute(response, nil)
  if err != nil {
    http.Error(response, err.Error(), http.StatusInternalServerError)
  }
}
