package controllers

import "simpl-go/users-roles/server"

func ErrorResponse(code int, message string) server.RestResponse {
	return server.RestResponse{
		Code: code,
		Body: ErrorResponseDTO{
			Message: message,
		},
	}
}
