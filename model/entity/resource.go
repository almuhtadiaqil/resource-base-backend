package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resource struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	ProfileID          uuid.UUID `gorm:"foreignKey"`
	JobStatus          bool
	IsFulltime         bool
	IsFreelance        bool
	IsOnsite           bool
	IsOfficialHour     bool
	IsProfileCompleted bool
	CategoryID         uuid.UUID `gorm:"foreignKey"`
	CreatedID          uuid.UUID
	UpdatedID          uuid.UUID
	DeletedID          uuid.UUID
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}
