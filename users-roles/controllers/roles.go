package controllers

import (
	"database/sql"
	"net/http"
	"simpl-go/users-roles/server"
	"simpl-go/users-roles/services"
)

func RoleFindById(r *http.Request) server.RestResponse {
	id := r.PathValue("id")
	if id == "" {
		return ErrorResponse(400, "Missing id path value")
	}

	role, err := services.RoleFindById(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorResponse(404, "Role not found. Id: %s", id)
		} else {
			return ErrorResponse(500, "Unable to retrieve role from service. %s", err)
		}
	}
	var dto RoleDTO
	dto.Id = role.Id
	return RestResponse(200, dto)
}
