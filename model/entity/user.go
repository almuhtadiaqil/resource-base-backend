package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	PhoneNumber int
	Email       string
	Password    string
	Profile     Profile
	IsActive    bool
	CreatedID   uuid.UUID
	UpdatedID   uuid.UUID
	DeletedID   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
