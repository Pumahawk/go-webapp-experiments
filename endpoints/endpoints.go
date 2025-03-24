package endpoints

import (
	"fmt"

	"pumahawk.com/webserver/server"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RestError(code int, format string, a ...any) server.HttpResponse {
	return server.HttpResponse{
		Code: code,
		Body: ErrorResponse{Message: fmt.Sprintf(format, a...)},
	}
}

func RestSuccess(body any) server.HttpResponse {
	return server.HttpResponse{
		Code: 200,
		Body: body,
	}
}

func DBRestError(err error) server.HttpResponse {
	return RestError(500, "Unable to open database connection [%v]", err)
}
