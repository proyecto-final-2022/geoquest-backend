package apitests

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id,identity" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
