package user

import (
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(c *gin.Context, email string, name string, username string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, id int, user domain.UserDTO) error
	DeleteUser(c *gin.Context, id int) error
	HashPassword(password string) (string, error)
	CheckPassword(providedPassword string, userPassword string) error
	CreateCoupon(c *gin.Context, userID int, description string, expirationYear int, expirationMonth time.Month, expirationDay int, expirationHour int) error
	GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error)
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
	user, err := s.repo.GetUser(c, id)

	return user, err
}

func (s *service) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	user, err := s.repo.GetUserByEmail(c, email)

	return user, err
}

func (s *service) UpdateUser(c *gin.Context, id int, user domain.UserDTO) error {

	err := s.repo.UpdateUser(c, id, user)

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

	return coupons, nil
}
