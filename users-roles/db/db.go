package db

import (
	"embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"text/template"
)

//go:embed queries/*
var queriesTmpl embed.FS
var QueryTmpl *template.Template

func init() {
	funcs := template.FuncMap{
		"param": func(p any)string{return ""},
	}

	tmpl, err := template.New("").Funcs(funcs).ParseFS(queriesTmpl, "queries/*")
	if err != nil {
		log.Fatalf("Unable to load queries: %v", err)
	}
	QueryTmpl = tmpl
}

func Query(name string, data any) (*QueryGen, error) {
	sb := strings.Builder{}
	var qg QueryGen

	funcs := template.FuncMap{
		"param": qg.Param,
	}

	tmpl := QueryTmpl.Funcs(funcs)
	err := tmpl.ExecuteTemplate(&sb, name, data)
	if err != nil {
		return nil, fmt.Errorf("Unable to execute query %s: %w", name, err)
	}

	qg.Sql = sb.String()
	return &qg, nil
}

type QueryGen struct {
	ParamCounter int
	Sql string
	Params []any
}

func (qg *QueryGen) Param(p any) string {
	qg.ParamCounter = qg.ParamCounter + 1
	res := "$" + strconv.Itoa(qg.ParamCounter)
	qg.Params = append(qg.Params, p)
	return res
}
