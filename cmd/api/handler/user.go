package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/auth"
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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginGoogleRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CouponRequest struct {
	Description     string `json:"description"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationDay   int    `json:"expiration_day"`
	ExpirationHour  int    `json:"expiration_hour"`
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
		var pass string
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if pass, err = u.service.HashPassword(req.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = u.service.CreateUser(c, req.Email, req.Name, req.Username, pass); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		tokenString, err := auth.GenerateJWT(req.Email, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tokenString)
	}
}

// @Summary New coupon
// @Schemes
// @Description Save new coupon
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body CouponRequest true "Coupon to save"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/coupons  [post]
func (u *User) CreateUserCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CouponRequest
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = u.service.CreateCoupon(c, paramId, req.Description, req.ExpirationYear, time.Month(req.ExpirationMonth), req.ExpirationDay, req.ExpirationHour); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Coupon
// @Schemes
// @Description Coupon
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/coupons  [get]
func (u *User) GetUserCoupons() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		coupons, err := u.service.GetCoupons(c, paramId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, coupons)
	}
}

// @Summary Login user
// @Schemes
// @Description Login user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User to log in"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/sessions [post]
func (u *User) LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		var user domain.UserDTO
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if user, err = u.service.GetUserByEmail(c, req.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = u.service.CheckPassword(req.Password, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		tokenString, err := auth.GenerateJWT(user.Email, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tokenString)

	}
}

func (u *User) LoginUserGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginGoogleRequest
		var err error

		tokenGoogle := c.GetHeader("Authorization")
		if tokenGoogle == "" {
			c.JSON(401, gin.H{"error": "request does not contain a Google access token"})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		/*
			// Validate the JWT is valid
			claims, err := auth.ValidateGoogleJWT(tokenGoogle)
			if err != nil {
				c.JSON(403, gin.H{"error": "Invalid google auth"})
				return
			}
			if claims.Email != req.Email {
				c.JSON(403, gin.H{"error": "Emails don't match"})
				return
			}
		*/
		// create a JWT for OUR app and give it back to the client for future requests
		tokenString, err := auth.GenerateJWT(req.Email, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tokenString)
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
