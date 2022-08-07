package client

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Service interface {
	CreateClient(c *gin.Context, name string, image string) error
	CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string) error
	AddTag(c *gin.Context, questID int, description string) error
	GetClientQuests(c *gin.Context, questID int) ([]domain.QuestInfoDTO, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) CreateClient(c *gin.Context, name string, image string) error {
	err := s.repo.CreateClient(c, name, image)

	return err
}

func (s *service) GetClientQuests(c *gin.Context, questID int) ([]domain.QuestInfoDTO, error) {
	quests, err := s.repo.GetClientQuests(c, questID)

	return quests, err
}

func (s *service) CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string) error {
	err := s.repo.CreateQuest(c, clientID, name, qualification, description, difficulty, duration)

	return err
}

func (s *service) AddTag(c *gin.Context, questID int, description string) error {
	err := s.repo.AddTag(c, questID, description)

	return err
}
