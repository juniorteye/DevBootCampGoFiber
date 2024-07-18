package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Review represents the model for a review
// @swagger:model
type Review struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title      string    `gorm:"type:varchar(50);not null;unique" validate:"required" json:"title"`
	Text       string    `gorm:"type:varchar(200);not null;unique" validate:"required" json:"text"`
	Rating     int       `gorm:"type: interger; check:average_rating >= 1 and average_rating <= 10" json:"rating"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	User       User      `gorm:"foreignKey:UserID"`
	BootcampID uuid.UUID `gorm:"type:uuid;not null" json:"bootcampID"`
	BootCamp   Bootcamp  `gorm:"foreignKey:BootcampID"`
}

func (review *Review) BeforeCreate(tx *gorm.DB) (err error) {
	review.ID = uuid.New()
	return
}

type Reviews struct {
	Reviews []Review `json:"reviews"`
}

// UserReview represents the structure for user response
type UserReview struct {
	ID       uuid.UUID `json:"id"`
	Username string    `gorm:"type:varchar(100);not null" json:"username"`
	Email    string    `gorm:"type:varchar(100);unique;not null" json:"email"`
}

// ReviewResponse represents the structure for review response
type ReviewResponse struct {
	ID         uuid.UUID  `json:"ID"`
	Title      string     `json:"title"`
	Text       string     `json:"text"`
	Rating     int        `json:"rating"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	BootcampID uuid.UUID  `json:"bootcampID"`
	User       UserReview `json:"user"`
}
