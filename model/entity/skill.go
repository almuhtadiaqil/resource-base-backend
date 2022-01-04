package entity

import "github.com/google/uuid"

type Skill struct {
	ID              uuid.UUID
	SkillCategoryID uuid.UUID
	Name            string
}
