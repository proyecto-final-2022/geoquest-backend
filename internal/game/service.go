package game

import (
	"geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Get(c *gin.Context, id int) (domain.Game, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) Get(c *gin.Context, id int) (domain.Game, error) {
	game, err := s.repo.Get(c, id)

	return game, err
}
