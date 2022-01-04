package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Competence struct {
	ID         uuid.UUID
	CategoryID uuid.UUID
	Competence string
	CreatedID  uuid.UUID
	UpdatedID  uuid.UUID
	DeletedID  uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
