package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetQuest(c *gin.Context, id int) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, name string) error
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) GetQuest(c *gin.Context, id int) (domain.QuestDTO, error) {
	quest, err := s.repo.GetQuest(c, id)

	return quest, err
}

func (s *service) CreateQuest(c *gin.Context, name string) error {
	err := s.repo.CreateQuest(c, name)

	return err
}
