package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(c *gin.Context, email string, name string, username string, image int, manual bool, google bool, facebook bool, password string, firebaseToken string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
	GetUsers(c *gin.Context) ([]domain.UserDTO, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, id int, email string, name string, password string, username string, image int, firebaseToken string) error
	DeleteUser(c *gin.Context, id int) error
	HashPassword(password string) (string, error)
	CheckPassword(providedPassword string, userPassword string) error
	CreateCoupon(c *gin.Context, userID int, description string, expirationYear int, expirationMonth time.Month, expirationDay int, expirationHour int) error
	GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error)
	UpdateCoupon(c *gin.Context, userID int, couponID int) error
	AddFriend(c *gin.Context, id int, friendID int) error
	GetUserFriends(c *gin.Context, id int) ([]domain.UserDTO, error)
	DeleteFriend(c *gin.Context, id int, friendID int) error
	AddNotification(c *gin.Context, ID int, senderID int, notificationType string, questName string, image int, teamID int, questID int) error
	GetNotifications(c *gin.Context, ID int) ([]domain.NotificationDTO, error)
	DeleteNotification(c *gin.Context, id int, notificationID int) error
	SendUpdateNewFriend(c *gin.Context, receiverID, senderID int) error
	SendUpdateQuestAccept(c *gin.Context, userID int, notificationID int) error
	SendUpdateAcceptFriend(c *gin.Context, userID int, friendID int) error
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) CreateUser(c *gin.Context, email string, name string, username string, image int, manual bool, google bool, facebook bool, password string, firebaseToken string) error {
	err := s.repo.CreateUser(c, email, name, username, image, manual, google, facebook, password, firebaseToken)

	return err
}

func (s *service) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	user, _, err := s.repo.GetUser(c, id)

	return user, err
}

func (s *service) GetUsers(c *gin.Context) ([]domain.UserDTO, error) {
	users, err := s.repo.GetUsers(c)

	usersDTO := make([]domain.UserDTO, len(users))
	for i := range users {
		usersDTO[i].ID = users[i].ID
		usersDTO[i].Name = users[i].Name
		usersDTO[i].Username = users[i].Username
		usersDTO[i].Email = users[i].Email
		usersDTO[i].Image = users[i].Image
		usersDTO[i].FirebaseToken = users[i].FirebaseToken
	}

	return usersDTO, err
}

func (s *service) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	user, err := s.repo.GetUserByEmail(c, email)

	return user, err
}

func (s *service) UpdateUser(c *gin.Context, id int, email string, name string, password string, username string, image int, firebaseToken string) error {

	_, user, err := s.repo.GetUser(c, id)
	if err != nil {
		return err
	}

	if name != "" {
		user.Name = name
	}

	if email != "" {
		user.Email = email
	}

	if password != "" {
		user.Password = password
	}

	if username != "" {
		user.Username = username
	}

	if image != 0 {
		user.Image = image
	}

	if firebaseToken != "" {
		user.FirebaseToken = firebaseToken
	}

	err = s.repo.UpdateUser(c, user)

	return err
}

func (s *service) DeleteUser(c *gin.Context, id int) error {
	err := s.repo.DeleteUser(c, id)

	return err
}

func (s *service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *service) CheckPassword(providedPassword string, userPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateCoupon(c *gin.Context, userID int, description string, expirationYear int, expirationMonth time.Month, expirationDay int, expirationHour int) error {

	tm := time.Date(expirationYear, expirationMonth, expirationDay, expirationHour, 00, 00, 00, time.UTC)

	if err := s.repo.CreateCoupon(c, userID, description, tm); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateCoupon(c *gin.Context, userID int, couponID int) error {

	coupon, err := s.repo.GetCoupon(c, couponID)

	if err != nil {
		return err
	}

	coupon.Used = true

	if err := s.repo.UpdateCoupon(c, coupon); err != nil {
		return err
	}

	if err := s.repo.UnlockAchivement(c, userID, "UsedCoupon_ac"); err != nil {
		return err
	}

	return nil
}

func (s *service) GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error) {

	coupons, err := s.repo.GetCoupons(c, userID)
	if err != nil {
		return nil, err
	}

	var couponsDTO []domain.CouponDTO

	for i := range coupons {
		if !coupons[i].Used {
			var couponDTO domain.CouponDTO
			couponDTO.ID = coupons[i].ID
			couponDTO.Description = coupons[i].Description
			couponDTO.ClientID = coupons[i].ClientID
			couponDTO.UserID = coupons[i].UserID
			couponDTO.ExpirationDate = coupons[i].ExpirationDate
			couponDTO.Used = coupons[i].Used
			couponsDTO = append(couponsDTO, couponDTO)
		}
	}

	return couponsDTO, nil
}

func (s *service) AddFriend(c *gin.Context, id int, friendID int) error {

	userFriends, err := s.repo.GetUserFriends(c, id)
	if err != nil {
		return err
	}

	for _, user := range userFriends {
		if user.FriendID == friendID {
			return nil
		}
	}

	if err := s.repo.AddFriend(c, id, friendID); err != nil {
		return err
	}

	if err := s.repo.UnlockAchivement(c, id, "MadeFriend_ac"); err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserFriends(c *gin.Context, id int) ([]domain.UserDTO, error) {

	users, err := s.repo.GetUserFriends(c, id)
	if err != nil {
		return nil, err
	}

	usersDTO := make([]domain.UserDTO, len(users))

	for i := range users {

		_, user, err := s.repo.GetUser(c, users[i].FriendID)

		if err != nil {
			return nil, err
		}

		usersDTO[i].ID = user.ID
		usersDTO[i].Name = user.Name
		usersDTO[i].Username = user.Username
		usersDTO[i].Email = user.Email
		usersDTO[i].Password = user.Password
		usersDTO[i].Image = user.Image
	}

	return usersDTO, nil
}

func (s *service) DeleteFriend(c *gin.Context, id int, friendID int) error {

	err := s.repo.DeleteFriend(c, id, friendID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) AddNotification(c *gin.Context, id int, senderID int, notificationType string, questName string, image int, teamID int, questID int) error {

	actualTime := time.Now()

	err := s.repo.AddNotification(c, id, senderID, notificationType, questName, image, teamID, questID, actualTime)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetNotifications(c *gin.Context, id int) ([]domain.NotificationDTO, error) {

	notifications, err := s.repo.GetNotifications(c, id)
	if err != nil {
		return nil, err
	}

	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].SentTime.After(notifications[j].SentTime)
	})

	notificationsDTO := make([]domain.NotificationDTO, len(notifications))

	for i := range notifications {
		senderDTO, _, err := s.repo.GetUser(c, notifications[i].SenderID)

		if err != nil {
			return nil, err
		}

		notificationsDTO[i].ID = notifications[i].ID
		notificationsDTO[i].SenderID = notifications[i].SenderID
		notificationsDTO[i].Type = notifications[i].Type
		notificationsDTO[i].SenderName = senderDTO.Username
		notificationsDTO[i].QuestName = notifications[i].QuestName
		notificationsDTO[i].TeamID = notifications[i].TeamID
		notificationsDTO[i].SenderName = senderDTO.Username
		notificationsDTO[i].QuestID = notifications[i].QuestID
		notificationsDTO[i].SenderImage = notifications[i].SenderImage
	}

	return notificationsDTO, nil
}

func (s *service) DeleteNotification(c *gin.Context, id int, notificationID int) error {
	err := s.repo.DeleteNotification(c, id, notificationID)

	return err
}

func (s *service) SendUpdateNewFriend(c *gin.Context, receiverID int, senderID int) error {

	senderDTO, _, err := s.repo.GetUser(c, senderID)
	receiverDTO, _, err := s.repo.GetUser(c, receiverID)

	if err != nil {
		return nil
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"sender_id":   senderID,
		"token":       receiverDTO.FirebaseToken,
		"sender_name": senderDTO.Username,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(config.GetConfig("prod").APP_NOTIFICATIONS_URL+"notifications/friend_request", "application/json", responseBody)
	//Handle Error
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (s *service) SendUpdateQuestAccept(c *gin.Context, userID int, notificationID int) error {
	fmt.Println("******Notification ID: ", notificationID)
	notificationDTO, err := s.repo.GetNotification(c, notificationID)
	if err != nil {
		return err
	}
	senderDTO, _, err := s.repo.GetUser(c, userID)
	if err != nil {
		return err
	}
	fmt.Println("******Quest ID: ", notificationDTO.QuestID)
	receiverDTO, _, err := s.repo.GetUser(c, notificationDTO.SenderID)
	if err != nil {
		return err
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"quest_id":    notificationDTO.QuestID,
		"token":       receiverDTO.FirebaseToken,
		"sender_name": senderDTO.Username,
		"team_id":     notificationDTO.TeamID,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(config.GetConfig("prod").APP_NOTIFICATIONS_URL+"notifications/quest_accept", "application/json", responseBody)
	fmt.Println("******Response status: ", resp.StatusCode)
	//Handle Error
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (s *service) SendUpdateAcceptFriend(c *gin.Context, friendID int, userID int) error {

	userDTO, _, err := s.repo.GetUser(c, userID)
	friendDTO, _, err := s.repo.GetUser(c, friendID)

	if err != nil {
		return nil
	}

	postBody, _ := json.Marshal(map[string]interface{}{
		"sender_name": userDTO.Username,
		"token":       friendDTO.FirebaseToken,
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(config.GetConfig("prod").APP_NOTIFICATIONS_URL+"notifications/friend_accept", "application/json", responseBody)
	//Handle Error
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
