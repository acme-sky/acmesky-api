package models

import (
	"time"
)

// Journey model
type Journey struct {
	Id        uint             `gorm:"column:id" json:"id"`
	CreatedAt time.Time        `gorm:"column:created_at" json:"created_at"`
	Flight1Id int              `json:"-"`
	Flight1   AvailableFlight  `gorm:"foreignKey:Flight1Id;null" json:"flight1"`
	Flight2Id *int             `json:"-"`
	Flight2   *AvailableFlight `gorm:"foreignKey:Flight2Id;null" json:"flight2"`
	Cost      float64          `gorm:"column:cost" json:"cost"`
	UserId    *int             `json:"-"`
	User      *User            `gorm:"foreignKey:UserId" json:"user"`
}
