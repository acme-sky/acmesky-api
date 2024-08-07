package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Offer model
type Offer struct {
	Id           uint      `gorm:"column:id" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	Message      string    `gorm:"column:message" json:"message"`
	Expired      string    `gorm:"column:expired" json:"expired"`
	Token        string    `gorm:"column:token" json:"token"`
	IsUsed       bool      `gorm:"column:is_used" json:"is_used"`
	PaymentLink  string    `gorm:"column:payment_link" json:"payment_link"`
	PaymentPaid  bool      `gorm:"column:payment_paid" json:"payment_paid"`
	RentEndpoint string    `gorm:"column:rent_endpoint" json:"rent_endpoint"`
	RentId       string    `gorm:"column:rent_id" json:"rent_id"`
	JourneyId    int       `json:"-"`
	Journey      Journey   `gorm:"foreignKey:JourneyId" json:"journey"`
	UserId       *int      `json:"-"`
	User         *User     `gorm:"foreignKey:UserId" json:"user"`
}

// Offer struct used to confirm a token
type OfferConfirmInput struct {
	Token string `binding:"required"`
}

// Offer last minute to be saved
type OfferLastMinuteInput struct {
	DepartureAirport string  `binding:"required" json:"departure_airport"`
	DepartureTime    string  `binding:"required" json:"departure_time"`
	ArrivalAirport   string  `binding:"required" json:"arrival_airport"`
	ArrivalTime      string  `binding:"required" json:"arrival_time"`
	Airline          string  `binding:"required" json:"airline"`
	Code             string  `binding:"required" json:"code"`
	Cost             float32 `binding:"required" json:"cost"`
}

// Validates an offer token
func ValidateOfferToken(db *gorm.DB, in OfferConfirmInput, userId uint) error {
	var user User

	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return errors.New("`user_id` does not exist.")
	}

	var offer Offer
	if err := db.Where("token = ?", in.Token).First(&offer).Error; err != nil {
		return errors.New("`token` does not exist.")
	}

	if !(user.IsAdmin || int(user.Id) == *offer.UserId) {
		return errors.New("`token` does not exist.")
	}

	return nil
}
