package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Profile struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	UserID          uuid.UUID `gorm:"foreignKey"`
	Nik             int
	BirthPlace      string
	BirthDate       datatypes.Date
	Rt              int
	Rw              int
	SubDistrictID   uuid.UUID `gorm:"foreignKey"`
	DistrictID      uuid.UUID `gorm:"foreignKey"`
	CityID          uuid.UUID `gorm:"foreignKey"`
	ProvinceID      uuid.UUID `gorm:"foreignKey"`
	ReligionID      uuid.UUID `gorm:"foreignKey"`
	Address         string
	DomisiliAddress string
	Gender          string
	CreatedID       uuid.UUID
	UpdatedID       uuid.UUID
	DeletedID       uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
