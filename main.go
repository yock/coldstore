package main

import (
  "os"
  "log"
  "net/http"
  "embed"
  "crypto/tls"
  "golang.org/x/crypto/acme/autocert"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "yock.dev/coldstore/cuts"
  "yock.dev/coldstore/home"
  "yock.dev/coldstore/data"
  "yock.dev/coldstore/scan"
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
    port = "8443"
  }

  certManager := autocert.Manager{
    Prompt: autocert.AcceptTOS,
    Cache: autocert.DirCache("certs"),
    HostPolicy: autocert.HostWhitelist("home.yock.dev"),
  }

  router := mux.NewRouter()
  cuts.Router(router.PathPrefix("/{cuts:cuts\\/?}").Subrouter())
  scan.Router(router.PathPrefix("/{scan:scan\\/?}").Subrouter())
  router.HandleFunc("/", home.HomeHandler)
  router.PathPrefix("/static/").Handler(http.FileServerFS(staticFS))
  router.PathPrefix("/favicon.ico").Handler(http.FileServerFS(staticFS))

  server := &http.Server{
    Addr: ":"+port,
    Handler: router,
    TLSConfig: &tls.Config{
      GetCertificate: certManager.GetCertificate,
    },
  }

  go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

  log.Println("Listening on", port)
  log.Fatal(server.ListenAndServeTLS("", ""))
}
