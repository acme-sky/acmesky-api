package models

import (
	"time"
)

// Invoice model
type Invoice struct {
	Id                uint      `gorm:"column:id" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	RentId            string    `gorm:"column:rent_id" json:"rent_id"`
	RentCustomerName  string    `gorm:"column:rent_customer_name" json:"rent_customer_name"`
	RentPickupAddress string    `gorm:"column:rent_pickup_address" json:"rent_pickup_address"`
	RentPickupDate    string    `gorm:"column:rent_pickup_date" json:"rent_pickup_date"`
	RentAddress       string    `gorm:"column:rent_address" json:"rent_address"`
	JourneyId         int       `json:"-"`
	Journey           Journey   `gorm:"foreignKey:JourneyId" json:"journey"`
	UserId            int       `json:"-"`
	User              User      `gorm:"foreignKey:UserId" json:"user"`
}
