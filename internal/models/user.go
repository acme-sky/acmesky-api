package models

import (
	"time"
)

// User model
// We ignore all the implementation for users having a manually creation. This
// model is used only for login.
type User struct {
	Id                 uint      `gorm:"column:id" json:"id"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"crated_at"`
	Username           string    `gorm:"column:username" gorm:"uniqueIndex" json:"username"`
	Password           string    `gorm:"column:password" json:"password"`
	Name               string    `gorm:"column:name" json:"name"`
	Email              string    `gorm:"column:email" gorm:"uniqueIndex" json:"email"`
	Address            *string   `gorm:"column:address;null" json:"address"`
	ProntogramUsername *string   `gorm:"column:prontogram_username;null" json:"prontogram_username"`
	IsAdmin            bool      `gorm:"column:is_admin" json:"is_admin"`
}

// Struct used to edit an user
type UserInput struct {
	Name               string `binding:"required"`
	Email              string `binding:"required"`
	Address            *string
	ProntogramUsername *string
	IsAdmin            bool
}
