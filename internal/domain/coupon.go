package domain

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	ID             int       `json:"id,identity" gorm:"primary_key"`
	UserID         int       `json:"user_id"`
	ClientID       int       `json:"client_id"`
	Description    string    `json:"description"`
	Used           bool      `json:"used"`
	ExpirationDate time.Time `json:"expiration_date"`
}

type CouponDTO struct {
	ID             int       `json:"id"`
	Description    string    `json:"description"`
	Used           bool      `json:"used"`
	ExpirationDate time.Time `json:"expiration_date"`
}
