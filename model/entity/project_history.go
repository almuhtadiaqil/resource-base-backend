package entity

import "github.com/google/uuid"

type ProjectHistory struct {
	ID          uuid.UUID
	ResourceID  uuid.UUID
	Name        string
	Description string
	TechStack   ProjecTechStack
}
