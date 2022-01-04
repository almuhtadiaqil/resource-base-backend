package entity

import "github.com/google/uuid"

type JobHistory struct {
	ID           uuid.UUID
	ResourceID   uuid.UUID
	OfficeName   string
	Experience   int
	CompetenceID uuid.UUID
}
