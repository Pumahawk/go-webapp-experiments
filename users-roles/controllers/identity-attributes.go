package controllers

import (
	"fmt"
	"log"
	"net/http"
	"simpl-go/users-roles/server"
	"simpl-go/users-roles/services"
	"strconv"
	"time"
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
		idasp.Id = &id
	}

	if assignedToParticipant := q.Get("assignedToParticipant"); assignedToParticipant != "" {
		assignedToParticipant, err := strconv.ParseBool(assignedToParticipant)

		if err != nil {
			return nil, fmt.Errorf("controller ida: Unable to extract parameter assignedToParticipant from query param. %w", err)
		}
		idasp.AssignedToParticipant = &assignedToParticipant
	}

	if code := q.Get("code"); code != "" {
		idasp.Code = &code
	}

	if enabled := q.Get("enabled"); enabled != "" {
		enabled, err := strconv.ParseBool(enabled)

		if err != nil {
			return nil, fmt.Errorf("controller ida: Unable to extract parameter enabled from query param. %w", err)
		}
		idasp.Enabled = &enabled
	}

	if id := q.Get("id"); id != "" {
		idasp.Id = &id
	}

	if name := q.Get("name"); name != "" {
		idasp.Name = &name
	}

	if participantTypeIn := q.Get("participantTypeIn"); participantTypeIn != "" {
		idasp.ParticipantTypeIn = &participantTypeIn
	}

	if participantTypeNotIn := q.Get("participantTypeNotIn"); participantTypeNotIn != "" {
		idasp.ParticipantTypeNotIn = &participantTypeNotIn
	}

	if updateTimestampFrom := q.Get("updateTimestampFrom"); updateTimestampFrom != "" {
		updateTimestampFrom, err := time.Parse(time.Layout, updateTimestampFrom)

		if err != nil {
			return nil, fmt.Errorf("controller ida: Unable to extract parameter updateTimestampFrom from query param. %w", err)
		}
		idasp.UpdateTimestampFrom = &updateTimestampFrom
	}

	if updateTimestampTo := q.Get("updateTimestampTo"); updateTimestampTo != "" {
		updateTimestampTo, err := time.Parse(time.Layout, updateTimestampTo)

		if err != nil {
			return nil, fmt.Errorf("controller ida: Unable to extract parameter updateTimestampTo from query param. %w", err)
		}
		idasp.UpdateTimestampTo = &updateTimestampTo
	}

	return &idasp, nil
}
