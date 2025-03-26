package controllers

import (
	"fmt"
	"simpl-go/users-roles/server"
)

func ErrorResponse(code int, format string, a ...any) server.RestResponse {
	return server.RestResponse{
		Code: code,
		Body: ErrorResponseDTO{
			Message: fmt.Sprintf(format, a...),
		},
	}
}

func RestResponse(code int, body any) server.RestResponse {
	return server.RestResponse{
		Code: code,
		Body: body,
	}
}
