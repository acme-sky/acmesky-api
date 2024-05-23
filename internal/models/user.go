package models

import "gorm.io/gorm"

// User model
// We ignore all the implementation for users having a manually creation. This
// model is used only for login.
type User struct {
	gorm.Model
	Username           string  `gorm:"column:username" gorm:"uniqueIndex"`
	Password           string  `gorm:"column:password"`
	Name               string  `gorm:"column:name"`
	Email              string  `gorm:"column:email" gorm:"uniqueIndex"`
	Address            *string `gorm:"colum:address;null"`
	ProntogramUsername *string `gorm:"colum:prontogram_username;null"`
}
