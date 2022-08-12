package client

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Repository interface {
	CreateClient(c *gin.Context, name string, image string) error
	CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string) error
	AddTag(c *gin.Context, questID int, description string) error
	GetClientQuests(c *gin.Context, questID int) ([]domain.QuestInfo, error)
	GetClients(c *gin.Context) ([]domain.Client, error)
	GetTags(c *gin.Context, questID int) ([]domain.Tag, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateClient(c *gin.Context, name string, image string) error {

	client := domain.Client{Name: name, Image: image}
	if tx := config.MySql.Create(&client); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) CreateQuest(c *gin.Context, clientID int, name string, qualification float32, description string, difficulty string, duration string) error {

	quest := domain.QuestInfo{ClientID: clientID, Name: name, Qualification: qualification, Description: description, Difficulty: difficulty, Duration: duration}

	if tx := config.MySql.Create(&quest); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil

}

func (r *repository) AddTag(c *gin.Context, questID int, description string) error {

	tag := domain.Tag{QuestID: questID, Description: description}

	if tx := config.MySql.Create(&tag); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil

}

func (r *repository) GetTags(c *gin.Context, questID int) ([]domain.Tag, error) {

	var tags []domain.Tag

	if tx := config.MySql.Where("quest_id = ?", questID).Find(&tags); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	fmt.Println(questID)
	fmt.Println(tags)

	return tags, nil

}

func (r *repository) GetClients(c *gin.Context) ([]domain.Client, error) {
	var clients []domain.Client
	if tx := config.MySql.Find(&clients); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return clients, nil
}

func (r *repository) GetClientQuests(c *gin.Context, questID int) ([]domain.QuestInfo, error) {
	var quests []domain.QuestInfo
	if tx := config.MySql.Where("client_id = ?", questID).Find(&quests); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return quests, nil
}
