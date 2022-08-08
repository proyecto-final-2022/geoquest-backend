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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserFriends struct {
	gorm.Model
	ID       int `json:"id,identity" gorm:"primary_key"`
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}
