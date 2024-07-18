package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type minimumSkill string

const (
	Beginner     minimumSkill = "beginner"
	Intermediate minimumSkill = "intermediate"
	Advanced     minimumSkill = "advanced"
)

func (ms minimumSkill) IsValid() bool {
	switch ms {
	case Beginner, Intermediate, Advanced:
		return true
	}
	return false
}

type Course struct {
	ID                   uuid.UUID    `gorm:"type:uuid;primaryKey"`
	Title                string       `gorm:"type:varchar(50);not null;unique" validate:"required" json:"title"`
	Description          string       `gorm:"type:varchar(200);not null;unique" validate:"required" json:"text"`
	Weeks                string       `gorm:"type:varchar(50);not null" validate:"required"`
	Tuition              int          `gorm:"not null" validate:"required"`
	MinimumSkill         minimumSkill `gorm:"type:varchar(50);not null;unique" validate:"required"`
	ScholarshipAvailable bool         `gorm:"type:boolean;not null default:false" json:"scholarshipAvailable"`
	CreatedAt            time.Time    `json:"CreatedAt"`
	UpdatedAt            time.Time    `json:"UpdatedAt"`
	User                 User         `gorm:"foreignKey:UserID"`
	UserID               uuid.UUID    `gorm:"type:uuid;not null" json:"userID"`
	Bootcamp             Bootcamp     `gorm:"foreignKey:BootcampID"`
	BootcampID           uuid.UUID    `gorm:"type:uuid;not null" json:"bootcampID"`
}

func (course *Course) BeforeCreate(tx *gorm.DB) (err error) {
	course.ID = uuid.New()
	return
}

func (course *Course) BeforeSave(tx *gorm.DB) (err error) {
	if !course.MinimumSkill.IsValid() {
		return errors.New("invalid minimum skill")
	}
	return
}
