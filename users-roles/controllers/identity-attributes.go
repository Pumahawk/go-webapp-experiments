package controllers

import (
	"log"
	"net/http"
	"simpl-go/users-roles/server"
	"simpl-go/users-roles/services"
)

func IdentityAttributeSearch(r *http.Request) server.RestResponse {
	result := make([]IdentityAttributeDTO, 0)
	idas, err := services.IdentityAttributeSearch(r.Context())
	if err != nil {
		log.Printf("Error DB search identity attributes. %v", err)
		return ErrorResponse(500, "Unable to search identity attributes")
	}
	for _, ida := range idas {
		var idar IdentityAttributeDTO
		idar.Id = ida.Id
		result = append(result, idar)
	}
	return RestResponse(200, result)
}
