package home

import (
  "log"
  "embed"
  "net/http"
  "html/template"
)

//go:embed templates/*.html
var templateFiles embed.FS

var indexTemplate *template.Template

func Init(layout *template.Template) {
  t, err := template.Must(layout.Clone()).ParseFS(templateFiles, "templates/index.html")
  if err != nil {
    log.Fatal(err.Error())
  }

  indexTemplate = t
}

func HomeHandler(response http.ResponseWriter, request *http.Request) {
  err := indexTemplate.Execute(response, nil)
  if err != nil {
    http.Error(response, err.Error(), http.StatusInternalServerError)
  }
}
