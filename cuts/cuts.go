package cuts

import (
  "embed"
  "net/http"
  "html/template"
  "time"
  "strconv"
  "github.com/gorilla/mux"
  "yock.dev/coldstore/layout"
  "yock.dev/coldstore/data"
)

//go:embed templates/*
var templateFiles embed.FS

var templates = map[string]*template.Template {
  "index": template.Must(layout.Layout().ParseFS(templateFiles, "templates/index.html")),
}

type Option struct {
  Value string
  Label string
}

type IndexModel struct {
  Title string
  Cuts []data.Cut
  MeatTypes []Option
}

var meatTypes = []Option {
  { Value: "beef", Label: "Beef" },
  { Value: "chicken", Label: "Chicken" },
  { Value: "pork", Label: "Pork" },
}


func indexHandler(response http.ResponseWriter, request *http.Request) {
  var cuts []data.Cut
  data.Conn.Find(&cuts)
  data := IndexModel {
    MeatTypes: meatTypes,
    Cuts: cuts,
  }
  err := templates["index"].Execute(response, data)
  if err != nil {
    http.Error(response, err.Error(), http.StatusInternalServerError)
  }
}

func createHandler(response http.ResponseWriter, request *http.Request) {
  request.ParseForm()
  weight, err := strconv.ParseInt(request.FormValue("weight"), 10, 64)
  if err != nil {
  }
  cut := data.Cut {
    Name: request.FormValue("name"),
    Meat: request.FormValue("meat_type"),
    Weight: weight,
    Unit: request.FormValue("unit"),
    AddedAt: time.Now(),
  }

  data.Conn.Create(&cut)
  http.Redirect(response, request, "/cuts", http.StatusSeeOther)
}

func Router(router *mux.Router) {
  router.HandleFunc("", indexHandler).Methods("GET")
  router.HandleFunc("", createHandler).Methods("POST")
}
