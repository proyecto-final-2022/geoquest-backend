package handler

import (
	"geoquest-backend/internal/domain"
	"geoquest-backend/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRequest struct {
	ID int `json:"user_id" binding:"required"`
}

type User struct {
	service user.Service
}

func NewUser(s user.Service) *User {
	return &User{service: s}
}

func (u *User) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userRequest
		var user domain.User
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
		}

		if user, err = u.service.Post(c, req.ID); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}
