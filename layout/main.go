package layout

import (
  "embed"
  "html/template"
  "golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed layout.html
var layoutFile embed.FS

var funcMap = template.FuncMap {
  "Title": cases.Title(language.English).String,
}

var layout *template.Template = template.Must(template.New("layout.html").Funcs(funcMap).ParseFS(layoutFile, "layout.html"))

func Layout() *template.Template {
  return template.Must(layout.Clone())
}
