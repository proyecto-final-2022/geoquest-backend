package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int      `json:"id,identity" gorm:"primary_key"`
	Name     string   `json:"name"`
	Username string   `json:"username" gorm:"unique"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"password"`
	Coupons  []Coupon `json:"coupons" gorm:"foreignKey:UserID;references:ID"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
