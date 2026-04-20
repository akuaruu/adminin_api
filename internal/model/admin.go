package model

import (
	"time"
)

type Admin struct {
	ID        	int64  `gorm:"primaryKey" json:"id"`

	Username 	string `gorm:"unique;not null" json:"username"`
	Email		string `gorm:"unique;not null" json:"email"`
	Password 	string `gorm:"not null" json:"-"`

	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}