package server

import (
	"net/http"
	"database/sql"

	"pumahawk.com/webserver/log"
)

type EndpointResult = func(http.ResponseWriter, *http.Request)
type EndpointFunc = func(AppContext) (EndpointResult)

type AppContext struct {
	Log *log.Logger
	DB *sql.DB
}

