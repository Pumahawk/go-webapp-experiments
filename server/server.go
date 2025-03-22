package server

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"pumahawk.com/webserver/log"
)

type EndpointResult = func(http.ResponseWriter, *http.Request)
type EndpointFunc = func(AppContext) (EndpointResult)

type AppContext struct {
	Log *log.Logger
	DB *sql.DB
}

func ErroResponse(ctx *AppContext, w io.Writer, format string, a ...any) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		log := ctx.Log
		log.Error("Unable to write error %w", err)
	}
}
