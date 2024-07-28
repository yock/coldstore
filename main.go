package main

import (
  "os"
  "log"
  "net/http"
  "html/template"
  "embed"
  "github.com/gorilla/mux"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
  "yock.dev/coldstore/cuts"
  "yock.dev/coldstore/home"
)

//go:embed templates/layout.html
var layoutTemplate embed.FS

func main () {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  db, err := gorm.Open(sqlite.Open("coldstore.sqlite"))
  if err != nil {
    log.Fatal("Could not open database")
  }

  db.AutoMigrate(&cuts.Cut{})

  layout, err := template.ParseFS(layoutTemplate, "templates/layout.html")
  if err != nil {
    log.Fatal(err.Error())
  }

  home.Init(layout)

  router := mux.NewRouter()
  cuts.Router(router.PathPrefix("/cuts").Subrouter())
  router.HandleFunc("/", home.HomeHandler)

  log.Println("Listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, router))
}
