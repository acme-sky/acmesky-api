package models

import (
	"time"
)

// AvailableFlight model
type AvailableFlight struct {
	Id               uint      `gorm:"column:id" json:"id"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	Airline          string    `gorm:"column:airline" json:"airline"`
	DepartureTime    time.Time `gorm:"column:departure_time" json:"departure_time"`
	DepartureAirport string    `gorm:"column:departure_airport" json:"departure_airport"`
	ArrivalTime      time.Time `gorm:"column:arrival_time" json:"arrival_time"`
	ArrivalAirport   string    `gorm:"column:arrival_airport" json:"arrival_airport"`
	Code             string    `gorm:"column:code" json:"code"`
	Cost             float64   `gorm:"column:cost" json:"cost"`
	InterestId       *int      `json:"-"`
	Interest         *Interest `gorm:"foreignKey:InterestId;null" json:"interest"`
	OfferSent        bool      `gorm:"column:offer_sent" json:"offer_sent"`
	UserId           *int      `json:"-"`
	User             *User     `gorm:"foreignKey:UserId" json:"user"`
}
