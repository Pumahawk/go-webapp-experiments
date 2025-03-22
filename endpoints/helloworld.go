package endpoints

import (
	_ "embed"
	"net/http"

	"pumahawk.com/webserver/server"
)

func HelloWorlsEndpoint(ctx *server.AppContext) server.EndpointResult {
	ctx.Log.Info("Create HelloWorlsEndpoint")
	return func(res http.ResponseWriter, req *http.Request) {
		ctx.Log.Info("Hello from %s", req.URL.Path)
		body := PageHtml
		res.WriteHeader(200)
		res.Write(body)
	}
}
