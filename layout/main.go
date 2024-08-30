package layout

import (
  "embed"
  "image"
  "bytes"
  "image/png"
  "log"
  "encoding/base64"
  "html/template"
  "golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed layout.html
var layoutFile embed.FS

var funcMap = template.FuncMap {
  "Title": cases.Title(language.English).String,
  "AsPng": AsPng,
}

var layout *template.Template = template.Must(template.New("layout.html").Funcs(funcMap).ParseFS(layoutFile, "layout.html"))

func Layout() *template.Template {
  return template.Must(layout.Clone())
}

func AsPng(i image.Image) string {
  buffer := new(bytes.Buffer)
  if err := png.Encode(buffer, i); err != nil {
    log.Println("Unable to encode image")
  }
  enc := base64.StdEncoding.EncodeToString(buffer.Bytes())
  return enc
}
