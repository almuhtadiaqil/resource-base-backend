package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourceCompetence struct {
	ID           uuid.UUID
	ResourceID   uuid.UUID
	CompetenceID uuid.UUID
	CreatedID    uuid.UUID
	UpdatedID    uuid.UUID
	DeletedID    uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
