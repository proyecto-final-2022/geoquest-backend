package handler

import (
	"net/http"
	"strconv"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type userRequest struct {
	ID       int    `json:"user_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	service user.Service
}

func NewUser(s user.Service) *User {
	return &User{service: s}
}

// @Summary New user
// @Schemes
// @Description Save new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body userRequest true "User to save"
// @Success 200 {object} domain.UserDTO
// @Failure 422
// @Failure 500
// @Router /users/ [post]
func (u *User) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req userRequest
		var user domain.UserDTO
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
		}

		if user, err = u.service.Post(c, req.ID, req.Email, req.Name, req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

// @Summary User
// @Schemes
// @Description User info
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.UserDTO
// @Failure 500
// @Router /users/{id} [get]
func (u *User) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.UserDTO
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if user, err = u.service.Get(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}
