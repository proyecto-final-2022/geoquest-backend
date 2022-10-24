package client

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Service interface {
	CreateClient(c *gin.Context, name string, image string) error
	GetClients(c *gin.Context) ([]domain.ClientDTO, error)
	CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string, image string) error
	AddTag(c *gin.Context, questID int, description []string) error
	GetClientQuests(c *gin.Context, clientID int) ([]domain.QuestInfoDTO, error)
	GetAllQuests(c *gin.Context) ([]domain.QuestInfoDTO, error)
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

func (s *service) GetClients(c *gin.Context) ([]domain.ClientDTO, error) {
	clients, err := s.repo.GetClients(c)

	clientsDTO := make([]domain.ClientDTO, len(clients))

	for i := range clients {
		clientsDTO[i].ID = clients[i].ID
		clientsDTO[i].Name = clients[i].Name
		clientsDTO[i].Image = clients[i].Image
	}

	return clientsDTO, err
}

func (s *service) GetAllQuests(c *gin.Context) ([]domain.QuestInfoDTO, error) {
	quests, err := s.repo.GetAllQuests(c)

	questsDTO := make([]domain.QuestInfoDTO, len(quests))

	for i := range quests {
		questsDTO[i].ID = int(quests[i].ID)
		questsDTO[i].Name = quests[i].Name
		questsDTO[i].Qualification = quests[i].Qualification
		questsDTO[i].Description = quests[i].Description
		questsDTO[i].Difficulty = quests[i].Difficulty
		questsDTO[i].Duration = quests[i].Duration
		questsDTO[i].Completions = int(quests[i].Completions)
		questsDTO[i].Image = quests[i].Image
	}

	return questsDTO, err
}

func (s *service) GetClientQuests(c *gin.Context, clientID int) ([]domain.QuestInfoDTO, error) {

	quests, err := s.repo.GetClientQuests(c, clientID)

	questsDTO := make([]domain.QuestInfoDTO, len(quests))
	for i := range quests {
		id := quests[i].ID

		questsDTO[i].ID = id
		questsDTO[i].Name = quests[i].Name
		questsDTO[i].Qualification = quests[i].Qualification
		questsDTO[i].Description = quests[i].Description
		questsDTO[i].Difficulty = quests[i].Difficulty
		questsDTO[i].Duration = quests[i].Duration
		questsDTO[i].Image = quests[i].Image
		questsDTO[i].Completions = quests[i].Completions

		tags, err := s.repo.GetTags(c, id)
		if err != nil {
			return nil, err
		}

		for j := range tags {
			questsDTO[i].Tags = append(questsDTO[i].Tags, tags[j].Description)
		}

	}

	return questsDTO, err
}

func (s *service) CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string, image string) error {
	err := s.repo.CreateQuest(c, clientID, name, qualification, description, difficulty, duration, image)

	return err
}

func (s *service) AddTag(c *gin.Context, questID int, description []string) error {
	err := s.repo.AddTag(c, questID, description)

	return err
}
