package user

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateUser(c *gin.Context, email string, name string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) CreateUser(c *gin.Context, email string, name string, password string) error {
	err := s.repo.CreateUser(c, email, name, password)

	return err
}

func (s *service) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	user, err := s.repo.GetUser(c, id)

	return user, err
}
