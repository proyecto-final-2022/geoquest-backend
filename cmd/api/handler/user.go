package handler

import (
	"net/http"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

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

// @Summary New user
// @Schemes
// @Description Save new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body userRequest true "User to save"
// @Success 200 {object} domain.User
// @Failure 500
// @Router /users/ [post]
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
