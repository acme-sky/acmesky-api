package models

import (
	"time"
)

// Offer model
type Offer struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Message   string    `gorm:"column:message" json:"message"`
	Expired   string    `gorm:"column:expired" json:"expired"`
	Token     string    `gorm:"column:token" json:"token"`
	IsUsed    bool      `gorm:"column:is_used" json:"is_user"`
	JourneyId int       `json:"-"`
	Journey   Journey   `gorm:"foreignKey:JourneyId" json:"journey"`
	UserId    *int      `json:"-"`
	User      *User     `gorm:"foreignKey:UserId" json:"user"`
}
