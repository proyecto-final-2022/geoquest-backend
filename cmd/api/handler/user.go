package handler

import (
	"fmt"
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
	Email         string `json:"email"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	FirebaseToken string `json:"firebaseToken"`
	Image         int    `json:"image"`
}

type LoginFacebookRequest struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	FirebaseToken string `json:"firebaseToken"`
	Image         int    `json:"image"`
}

type UserResponse struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Username      string   `json:"username"`
	FirebaseToken string   `json:"firebaseToken"`
	Image         int      `json:"image"`
	Token         string   `json:"token"`
	Manual        bool     `json:"manual"`
	Google        bool     `json:"google"`
	Facebook      bool     `json:"facebook"`
	Friends       int      `json:"friends"`
	Notifications int      `json:"notifications"`
	Achivements   []string `json:"achivements"`
}

type UserRequest struct {
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebaseToken"`
	Image         int    `json:"image"`
	Manual        bool   `json:"manual"`
	Google        bool   `json:"google"`
	Facebook      bool   `json:"facebook"`
}

type UserPasswordChangeRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type CouponRequest struct {
	Description     string `json:"description"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationDay   int    `json:"expiration_day"`
	ExpirationHour  int    `json:"expiration_hour"`
}

type NotificationRequest struct {
	SenderID         int    `json:"sender_id"`
	SenderImage      int    `json:"sender_image"`
	NotificationType string `json:"type"`
	QuestName        string `json:"quest_name"`
	TeamID           int    `json:"team_id"`
	QuestID          int    `json:"quest_id"`
}

type NotificationDeleteRequest struct {
	NotificationType string `json:"type,omitempty"`
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
		var createdUser domain.UserDTO

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

		if req.Image == 0 {
			req.Image = 1
		}

		if err = u.service.CreateUser(c, req.Email, req.Name, req.Username, req.Image, true, false, false, pass, req.FirebaseToken); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		tokenString, err := auth.GenerateJWT(req.Email, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if createdUser, err = u.service.GetUserByEmail(c, req.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userResponse := UserResponse{
			ID:            createdUser.ID,
			Name:          createdUser.Name,
			Email:         createdUser.Email,
			Username:      createdUser.Username,
			FirebaseToken: createdUser.FirebaseToken,
			Image:         createdUser.Image,
			Manual:        createdUser.Manual,
			Google:        createdUser.Google,
			Facebook:      createdUser.Facebook,
			Token:         tokenString,
			Friends:       0,
			Notifications: 0,
		}

		c.JSON(http.StatusOK, userResponse)
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

// @Summary New friend
// @Schemes
// @Description Add new friend
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param friend_id path int true "User ID of user's friend"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/friends/{friend_id}  [post]
func (u *User) AddUserFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramFriendId, _ := strconv.Atoi(c.Param("friend_id"))

		if err = u.service.AddFriend(c, paramId, paramFriendId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err = u.service.AddFriend(c, paramFriendId, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		u.service.SendUpdateAcceptFriend(c, paramId, paramFriendId)

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Delete friend
// @Schemes
// @Description Delete friend from user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param friend_id path int true "User ID of user's friend to delete"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/friends/{friend_id}  [delete]
func (u *User) DeleteUserFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramFriendId, _ := strconv.Atoi(c.Param("friend_id"))

		if err = u.service.DeleteFriend(c, paramId, paramFriendId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary User's friends
// @Schemes
// @Description Get friends from user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/friends  [get]
func (u *User) GetUserFriends() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		users, err := u.service.GetUserFriends(c, paramId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// @Summary Get all users
// @Schemes
// @Description Get users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/ [get]
func (u *User) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		users, err := u.service.GetUsers(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, users)
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

		var friends []domain.UserDTO
		if friends, err = u.service.GetUserFriends(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var notifications []domain.NotificationDTO
		if notifications, err = u.service.GetNotifications(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userResponse := UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			Username:      user.Username,
			FirebaseToken: user.FirebaseToken,
			Image:         user.Image,
			Manual:        user.Manual,
			Google:        user.Google,
			Facebook:      user.Facebook,
			Token:         tokenString,
			Friends:       len(friends),
			Notifications: len(notifications),
		}

		c.JSON(http.StatusOK, userResponse)
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
// @Router /users/sessions/google [post]
func (u *User) LoginUserGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginGoogleRequest
		var createdUser domain.UserDTO

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

		var image int
		if req.Image == 0 {
			image = 1
		}

		//If user is not created, create it. If it is, get
		if createdUser, err = u.service.GetUserByEmail(c, req.Email); err != nil {
			if err = u.service.CreateUser(c, req.Email, req.Name, req.Username, image, false, true, false, "", req.FirebaseToken); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			if createdUser, err = u.service.GetUserByEmail(c, req.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// create a JWT for OUR app and give it back to the client for future requests
		tokenString, err := auth.GenerateJWT(req.Email, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var friends []domain.UserDTO
		if friends, err = u.service.GetUserFriends(c, createdUser.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var notifications []domain.NotificationDTO
		if notifications, err = u.service.GetNotifications(c, createdUser.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userResponse := UserResponse{
			ID:            createdUser.ID,
			Email:         createdUser.Email,
			Name:          createdUser.Name,
			Username:      createdUser.Username,
			FirebaseToken: createdUser.FirebaseToken,
			Image:         createdUser.Image,
			Manual:        createdUser.Manual,
			Google:        createdUser.Google,
			Facebook:      createdUser.Facebook,
			Token:         tokenString,
			Friends:       len(friends),
			Notifications: len(notifications),
		}

		fmt.Println(userResponse)

		c.JSON(http.StatusOK, userResponse)
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
// @Router /users/sessions/facebook [post]
func (u *User) LoginUserFacebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginFacebookRequest
		var createdUser domain.UserDTO

		var err error

		tokenFacebook := c.GetHeader("Authorization")
		if tokenFacebook == "" {
			c.JSON(401, gin.H{"error": "request does not contain a Google access token"})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		var image int
		if req.Image == 0 {
			image = 1
		}

		//If user is not created, create it. If it is, get
		if createdUser, err = u.service.GetUserByEmail(c, req.Email); err != nil {
			if err = u.service.CreateUser(c, req.Email, req.Name, req.Username, image, false, false, true, "", req.FirebaseToken); err != nil {
				c.JSON(http.StatusInternalServerError, err)
				return
			}

			if createdUser, err = u.service.GetUserByEmail(c, req.Email); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// create a JWT for OUR app and give it back to the client for future requests
		tokenString, err := auth.GenerateJWT(req.Email, req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var friends []domain.UserDTO
		if friends, err = u.service.GetUserFriends(c, createdUser.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var notifications []domain.NotificationDTO
		if notifications, err = u.service.GetNotifications(c, createdUser.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userResponse := UserResponse{
			ID:            createdUser.ID,
			Email:         createdUser.Email,
			Name:          createdUser.Name,
			Username:      createdUser.Username,
			FirebaseToken: createdUser.FirebaseToken,
			Image:         createdUser.Image,
			Manual:        createdUser.Manual,
			Google:        createdUser.Google,
			Facebook:      createdUser.Facebook,
			Token:         tokenString,
			Friends:       len(friends),
			Notifications: len(notifications),
		}

		fmt.Println(userResponse)

		c.JSON(http.StatusOK, userResponse)
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

		// create a JWT for OUR app and give it back to the client for future requests
		tokenString, err := auth.GenerateJWT(user.Email, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var friends []domain.UserDTO
		if friends, err = u.service.GetUserFriends(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var notifications []domain.NotificationDTO
		if notifications, err = u.service.GetNotifications(c, user.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var achivements []string

		//Achivements
		if user.MadeFriend_ac == true {
			achivements = append(achivements, "MadeFriend_ac")
		}
		if user.StartedQuest_ac == true {
			achivements = append(achivements, "StartedQuest_ac")
		}
		if user.FinishedQuest_ac == true {
			achivements = append(achivements, "FinishedQuest_ac")
		}
		if user.FinishedTeamQuest_ac == true {
			achivements = append(achivements, "FinishedTeamQuest_ac")
		}
		if user.RatedQuest_ac == true {
			achivements = append(achivements, "RatedQuest_ac")
		}

		userResponse := UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			Name:          user.Name,
			Username:      user.Username,
			FirebaseToken: user.FirebaseToken,
			Image:         user.Image,
			Manual:        user.Manual,
			Google:        user.Google,
			Facebook:      user.Facebook,
			Token:         tokenString,
			Friends:       len(friends),
			Notifications: len(notifications),
			Achivements:   achivements,
		}

		c.JSON(http.StatusOK, userResponse)
	}
}

// @Summary User
// @Schemes
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UserRequest true "User to update"
// @Param id path string true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id} [put]
func (g *User) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserRequest
		var err error

		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = g.service.UpdateUser(c, paramId, req.Email, req.Name, req.Password, req.Username, req.Image, req.FirebaseToken); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary User
// @Schemes
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param user body UserPasswordChangeRequest true "User to update"
// @Param id path string true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/updateUserPassword [put]
func (g *User) UpdateUserPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req UserPasswordChangeRequest
		var user domain.UserDTO
		var pass string
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId, _ := strconv.Atoi(c.Param("id"))

		if user, err = g.service.GetUser(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if err = g.service.CheckPassword(req.OldPassword, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if pass, err = g.service.HashPassword(req.NewPassword); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = g.service.UpdateUser(c, paramId, user.Email, user.Name, pass, user.Username, user.Image, user.FirebaseToken); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Coupon
// @Schemes
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param coupon_id path string true "Coupon ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/coupons/{coupon_id} [put]
func (g *User) UpdateCouponUsed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramCouponId, _ := strconv.Atoi(c.Param("coupon_id"))

		if err = g.service.UpdateCoupon(c, paramId, paramCouponId); err != nil {
			c.JSON(http.StatusInternalServerError, paramCouponId)
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

// @Summary New notification
// @Schemes
// @Description Add new notification
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body NotificationRequest true "Notification: Specify Sender and Type of notification: 'friend_request' or 'quest_invite'"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/notifications/  [post]
func (u *User) AddNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var req NotificationRequest

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.AddNotification(c, paramId, req.SenderID, req.NotificationType, req.QuestName, req.SenderImage, req.TeamID, req.QuestID); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		if req.NotificationType == "friend_request" {
			u.service.SendUpdateNewFriend(c, paramId, req.SenderID)
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Get notifications
// @Schemes
// @Description Get notifications
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /users/{id}/notifications/  [get]
func (u *User) GetNotifications() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		notifications, err := u.service.GetNotifications(c, paramId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, notifications)
	}
}

// @Summary Users
// @Schemes
// @Description Delete a notification
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param notification_id path string true "Notification ID"
// @Param user body NotificationDeleteRequest true "Notification: Specify Sender and Type of notification: 'friend_request' or 'quest_invite'"
// @Success 200
// @Failure 500
// @Router /users/{id}/notifications/{notification_id} [delete]
func (g *User) DeleteNotification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var req NotificationDeleteRequest

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramNotificationId, _ := strconv.Atoi(c.Param("notification_id"))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		fmt.Println(paramNotificationId)
		if err = g.service.DeleteNotification(c, paramId, paramNotificationId); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		if req.NotificationType == "quest_accept" {
			g.service.SendUpdateQuestAccept(c, paramId, paramNotificationId)
		}

		c.JSON(http.StatusOK, "")
	}
}
