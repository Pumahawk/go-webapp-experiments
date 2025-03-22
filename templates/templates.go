package templates

import (
	"embed"
	"fmt"
	htmltemplate "html/template"
	"log"
	"strings"
)

//go:embed templates/*.tmpl.*
var tmplFiles embed.FS

func LoadTemplateOrFatal() *htmltemplate.Template {
	tpl := htmltemplate.New("")
	funcMap := htmltemplate.FuncMap{
		"nilV": NilV,
	}
	tpl.Funcs(funcMap)
	tpl, err := tpl.ParseFS(tmplFiles, "templates/**.tmpl.**")
	if err != nil {
		log.Fatal("Unable to load embed template files", err)
	}
	return tpl
}

func NilV(v any) string {
	if v != nil {
		if s, ok := v.(string); ok && !strings.EqualFold(s, "<nil>") {
			return fmt.Sprintf("%v", v)
		} else {
			return "nil"
		}
	} else {
		return "nil"
	}
}
