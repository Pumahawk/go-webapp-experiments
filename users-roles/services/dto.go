package services

import "time"

type RoleInfo struct {
	Id string
}

type IdentityAttributeInfo struct {
	Id string
}

type IdentityAttributeSearchParams struct {
	AssignedToParticipant *bool
	Code                  *string
	Enabled               *bool
	Id                    *string
	Name                  *string
	ParticipantTypeIn     *string
	ParticipantTypeNotIn  *string
	UpdateTimestampFrom   *time.Time
	UpdateTimestampTo     *time.Time
}
