package controllers

import (
	"log"
	"net/http"
	"simpl-go/users-roles/server"
	"simpl-go/users-roles/services"
)

func IdentityAttributeSearch(r *http.Request) server.RestResponse {
	result := make([]IdentityAttributeDTO, 0)
	params, err := GetIdentityAttributeSearchParams(r)
	if err != nil {
		return ErrorResponse(400, "Invalid query parameters. %v", err)
	}

	idas, err := services.IdentityAttributeSearch(r.Context(), *params)
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

func GetIdentityAttributeSearchParams(r *http.Request) (*services.IdentityAttributeSearchParams, error) {
	var idasp services.IdentityAttributeSearchParams
	q := r.URL.Query()
	if id := q.Get("id"); id != "" {
		idasp.Id = id
	}
	return &idasp, nil
}
