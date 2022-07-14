package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
