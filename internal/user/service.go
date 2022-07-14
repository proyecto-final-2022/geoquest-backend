package user

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error)
	Get(c *gin.Context, id int) (domain.UserDTO, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error) {
	user, err := s.repo.Post(c, id, email, name, password)

	return user, err
}

func (s *service) Get(c *gin.Context, id int) (domain.UserDTO, error) {
	user, err := s.repo.Get(c, id)

	return user, err
}
