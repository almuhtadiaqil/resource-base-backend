package entity

import "github.com/google/uuid"

type ResourceSkill struct {
	ID         uuid.UUID
	ResourceID uuid.UUID
	SkillID    uuid.UUID
	Experience int
}
