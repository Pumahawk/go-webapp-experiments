package controllers

type ErrorResponseDTO struct {
	Message string `json:"message"`
}

type RoleDTO struct {
	Id string `json:"id"`
}

type IdentityAttributeDTO struct {
	Id string `json:"id"`
}
