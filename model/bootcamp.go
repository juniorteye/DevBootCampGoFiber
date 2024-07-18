package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

// Bootcamp represents the model for a bootcamp
// @swagger:model
type Bootcamp struct {
	// gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"type:varchar(20);not null;unique" validate:"required" json:"name"`
	Slug          string    `gorm:"type:varchar(100)" json:"slug"`
	Description   string    `gorm:"type:varchar(500);not null" validate:"required" json:"description"`
	Website       string    `gorm:"type:varchar(100)" validate:"url" json:"website"`
	Phone         string    `gorm:"type:varchar(20)" json:"phone"`
	Email         string    `gorm:"type:varchar(100)" validate:"email" json:"email"`
	Address       string    `gorm:"type:varchar(200);not null" validate:"required" json:"address"`
	City          string    `gorm:"type:varchar(100)" json:"city"`
	State         string    `gorm:"type:varchar(100)" json:"state"`
	Zipcode       string    `gorm:"type:varchar(20)" json:"zipcode"`
	Country       string    `gorm:"type:varchar(100)" json:"country"`
	Type          string    `gorm:"type:varchar(100);not null" validate:"required" json:"type"`
	Careers       string    `gorm:"type:varchar(200);not null" validate:"required" json:"careers"`
	AverageRating float64   `gorm:"type:float;check:average_rating >= 1 and average_rating <= 10" json:"averageRating"`
	AverageCost   float64   `gorm:"type:float" json:"averageCost"`
	Photo         string    `gorm:"type:varchar(100);default:'no-photo.jpg'" json:"photo"`
	Housing       bool      `gorm:"default:false" json:"housing"`
	JobAssistance bool      `gorm:"default:false" json:"jobAssistance"`
	JobGuarantee  bool      `gorm:"default:false" json:"jobGuarantee"`
	AcceptGi      bool      `gorm:"default:false" json:"acceptGi"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UserID        uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	User          User      `gorm:"foreignKey:UserID"`
}

func (bootcamp *Bootcamp) BeforeCreate(tx *gorm.DB) (err error) {
	bootcamp.ID = uuid.New()
	return
}

type Bootcamps struct {
	Bootcamps []Bootcamp `json:"bootcamps"`
}

// Location struct for embedded location fields
// type Location struct {
// 	Type             string    `gorm:"type:varchar(50);not null" json:"type"`
// 	Coordinates      []float64 `gorm:"type:float[];not null" json:"coordinates"`
// 	FormattedAddress string    `gorm:"type:varchar(100)" json:"formattedAddress"`
// 	Street           string    `gorm:"type:varchar(100)" json:"street"`
// 	City             string    `gorm:"type:varchar(100)" json:"city"`
// 	State            string    `gorm:"type:varchar(100)" json:"state"`
// 	Zipcode          string    `gorm:"type:varchar(20)" json:"zipcode"`
// 	Country          string    `gorm:"type:varchar(100)" json:"country"`
// }
