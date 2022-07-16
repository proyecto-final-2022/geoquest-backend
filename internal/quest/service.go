package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Get(c *gin.Context, id int) (domain.Quest, error)
	Post(c *gin.Context, id int, name string) (domain.Quest, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) Get(c *gin.Context, id int) (domain.Quest, error) {
	quest, err := s.repo.Get(c, id)

	return quest, err
}

func (s *service) Post(c *gin.Context, id int, name string) (domain.Quest, error) {
	quest, err := s.repo.Post(c, id, name)

	return quest, err
}
