package templates

import (
	"embed"
	"log"
	"text/template"
)

//go:embed *.tmpl.*
var tmplFiles embed.FS

func LoadTemplateOrFatal() *template.Template {
	tpl, err := template.ParseFS(tmplFiles, "**.tmpl.**")
	if err != nil {
		log.Fatal("Unable to load embed template files", err)
	}
	return tpl;
}
