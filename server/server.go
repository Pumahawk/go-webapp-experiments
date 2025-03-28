package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"pumahawk.com/webserver/log"
)

type EndpointResult = func(http.ResponseWriter, *http.Request)
type EndpointFunc = func(AppContext) EndpointResult

type AppContext struct {
	Log      *log.Logger
	DB       *sql.DB
	Template *template.Template
}

func ErroResponse(ctx *AppContext, w io.Writer, format string, a ...any) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		log := ctx.Log
		log.Error("Unable to write error %w", err)
	}
}

func ErrorResponseFunc(ctx *AppContext, w io.Writer) func(string, ...any) {
	return func(format string, a ...any) {
		ErroResponse(ctx, w, format, a...)
	}
}

func JsonEndpoint(controller func(r *http.Request) HttpResponse) EndpointResult {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := controller(r)
		w.WriteHeader(resp.Code)
		body := resp.Body
		if body != nil {
			json.NewEncoder(w).Encode(body)
		}
	}
}

type HttpResponse struct {
	Code int
	Body any
}
