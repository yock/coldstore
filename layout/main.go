package layout

import (
  "embed"
  "html/template"
)

//go:embed layout.html
var layoutFile embed.FS

var layout *template.Template = template.Must(template.ParseFS(layoutFile, "layout.html"))

func Layout() *template.Template {
  return template.Must(layout.Clone())
}
