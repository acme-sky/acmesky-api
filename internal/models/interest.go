package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Interest model
type Interest struct {
	Id        uint      `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`

	Flight1DepartureTime    time.Time `gorm:"column:flight1_departure_time" json:"flight1_departure_time"`
	Flight1DepartureAirport string    `gorm:"column:flight1_departure_airport" json:"flight1_departure_airport"`
	Flight1ArrivalTime      time.Time `gorm:"column:flight1_arrival_time" json:"flight1_arrival_time"`
	Flight1ArrivalAirport   string    `gorm:"column:flight1_arrival_airport" json:"flight1_arrival_airport"`

	Flight2DepartureTime    *time.Time `gorm:"column:flight2_departure_time;null" json:"flight2_departure_time"`
	Flight2DepartureAirport *string    `gorm:"column:flight2_departure_airport;null" json:"flight2_departure_airport"`
	Flight2ArrivalTime      *time.Time `gorm:"column:flight2_arrival_time;null" json:"flight2_arrival_time"`
	Flight2ArrivalAirport   *string    `gorm:"column:flight2_arrival_airport;null" json:"flight2_arrival_airport"`

	UserId int   `json:"-"`
	User   *User `gorm:"foreignKey:UserId" json:"user"`
}

// Struct used to get new data for a flight
type InterestInput struct {
	Flight1DepartureTime    time.Time  `json:"flight1_departure_time" binding:"required"`
	Flight1DepartureAirport string     `json:"flight1_departure_airport" binding:"required"`
	Flight1ArrivalTime      time.Time  `json:"flight1_arrival_time" binding:"required"`
	Flight1ArrivalAirport   string     `json:"flight1_arrival_airport" binding:"required"`
	Flight2DepartureTime    *time.Time `json:"flight2_departure_time"`
	Flight2DepartureAirport *string    `json:"flight2_departure_airport"`
	Flight2ArrivalTime      *time.Time `json:"flight2_arrival_time"`
	Flight2ArrivalAirport   *string    `json:"flight2_arrival_airport"`
}

// It validates data from `in` and returns a possible error or not
func ValidateInterest(db *gorm.DB, in InterestInput, userId uint) error {
	var user User

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return errors.New("`user_id` does not exist.")
	}

	if in.Flight1DepartureAirport == in.Flight1ArrivalAirport {
		return errors.New("`flight1`: `departure_airport` can't be equals to `arrival_airport`")
	}

	if in.Flight1DepartureTime.After(in.Flight1ArrivalTime) {
		return errors.New("`flight1`: `departure_time` can't be after `arrival_time`")
	}

	if in.Flight2DepartureAirport != nil && in.Flight2DepartureTime != nil && in.Flight2ArrivalAirport != nil && in.Flight2ArrivalTime != nil {
		if (*in.Flight2DepartureAirport) == (*in.Flight2ArrivalAirport) {
			return errors.New("`flight2`: `departure_airport` can't be equals to `arrival_airport`")
		}

		if (*in.Flight2DepartureTime).After(*in.Flight2ArrivalTime) {
			return errors.New("`flight2`: `departure_time` can't be after `arrival_time`")
		}
	} else if !(in.Flight2DepartureAirport == nil || in.Flight2DepartureTime == nil || in.Flight2ArrivalAirport == nil || in.Flight2ArrivalTime == nil) {
		return errors.New("`flight2`: all fields must be nil or filled")
	}

	return nil
}
