package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID         uuid.UUID
	ResourceID uuid.UUID
	Name       string
	CreatedID  uuid.UUID
	UpdatedID  uuid.UUID
	DeletedID  uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
