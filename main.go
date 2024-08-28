package main

import (
  "os"
  "log"
  "net/http"
  "embed"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "yock.dev/coldstore/cuts"
  "yock.dev/coldstore/home"
  "yock.dev/coldstore/data"
)

//go:embed static/*
//go:embed favicon.ico
var staticFS embed.FS

func main () {
  log.Println("Loading environment file")
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Could not load .env file")
  }

  data.Connect()
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  router := mux.NewRouter()
  cuts.Router(router.PathPrefix("/{cuts:cuts\\/?}").Subrouter())
  router.HandleFunc("/", home.HomeHandler)
  router.PathPrefix("/static/").Handler(http.FileServerFS(staticFS))
  router.PathPrefix("/favicon.ico").Handler(http.FileServerFS(staticFS))

  log.Println("Listening on", port)
  log.Fatal(http.ListenAndServe(":"+port, router))
}
