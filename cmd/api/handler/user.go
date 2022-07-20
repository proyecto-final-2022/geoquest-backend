package handler

import (
	"net/http"
	"strconv"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

	"github.com/gin-gonic/gin"
)

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
// @Param user body domain.UserDTO true "User to save"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/ [post]
func (u *User) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.UserDTO
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.CreateUser(c, req.Email, req.Name, req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, nil)
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

		if user, err = u.service.GetUser(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// @Summary User
// @Schemes
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param quest body domain.UserDTO true "User to update"
// @Param id path string true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id} [put]
func (g *User) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.UserDTO
		var err error

		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = g.service.UpdateUser(c, paramId, req); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Users
// @Schemes
// @Description Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 500
// @Router /users/{id} [delete]
func (g *User) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = g.service.DeleteUser(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
