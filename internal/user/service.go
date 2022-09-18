package user

import (
	"sort"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(c *gin.Context, email string, name string, username string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, id int, email string, name string, password string, username string) error
	DeleteUser(c *gin.Context, id int) error
	HashPassword(password string) (string, error)
	CheckPassword(providedPassword string, userPassword string) error
	CreateCoupon(c *gin.Context, userID int, description string, expirationYear int, expirationMonth time.Month, expirationDay int, expirationHour int) error
	GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error)
	AddFriend(c *gin.Context, id int, friendID int) error
	GetUserFriends(c *gin.Context, id int) ([]domain.UserDTO, error)
	DeleteFriend(c *gin.Context, id int, friendID int) error
	AddNotification(c *gin.Context, ID int, senderID int, notificationType string) error
	GetNotifications(c *gin.Context, ID int) ([]domain.NotificationDTO, error)
	DeleteNotification(c *gin.Context, id int, notificationID int) error
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) CreateUser(c *gin.Context, email string, name string, username string, password string) error {
	err := s.repo.CreateUser(c, email, name, username, password)

	return err
}

func (s *service) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	user, _, err := s.repo.GetUser(c, id)

	return user, err
}

func (s *service) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	user, err := s.repo.GetUserByEmail(c, email)

	return user, err
}

func (s *service) UpdateUser(c *gin.Context, id int, email string, name string, password string, username string) error {

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

func (s *service) GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error) {

	coupons, err := s.repo.GetCoupons(c, userID)
	if err != nil {
		return nil, err
	}

	//servicio
	couponsDTO := make([]domain.CouponDTO, len(coupons))
	for i := range coupons {
		couponsDTO[i].ID = coupons[i].ID
		couponsDTO[i].Description = coupons[i].Description
		couponsDTO[i].ExpirationDate = coupons[i].ExpirationDate
		couponsDTO[i].Used = coupons[i].Used
	}

	return couponsDTO, nil
}

func (s *service) AddFriend(c *gin.Context, id int, friendID int) error {

	err := s.repo.AddFriend(c, id, friendID)
	if err != nil {
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

func (s *service) AddNotification(c *gin.Context, id int, senderID int, notificationType string) error {

	actualTime := time.Now()

	err := s.repo.AddNotification(c, id, senderID, notificationType, actualTime)
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
	}

	return notificationsDTO, nil
}

func (s *service) DeleteNotification(c *gin.Context, id int, notificationID int) error {
	err := s.repo.DeleteNotification(c, id, notificationID)

	return err
}
