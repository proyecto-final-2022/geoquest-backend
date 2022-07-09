package user

import (
	"geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Post(c *gin.Context, id int) (domain.User, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) Post(c *gin.Context, id int) (domain.User, error) {
	user, err := s.repo.Get(c, id)

	return user, err
}
