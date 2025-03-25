package controllers

import (
	"net/http"
	"simpl-go/users-roles/server"
)

func RoleFindById(r *http.Request) server.RestResponse {
	resp := ErrorResponse(500, "Not yet implemented")
	return resp
}
