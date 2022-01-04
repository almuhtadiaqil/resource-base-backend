package entity

import "github.com/google/uuid"

type ProjecTechStack struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	TechID    uuid.UUID
}
