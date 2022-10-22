package quest

import (
	"errors"
	"testing"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {

	var quests []*domain.QuestDTO

	quests = append(quests, &domain.QuestDTO{Name: "test"})

	return quests, nil
}

func (d *dummyRepo) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	if id == "9" {
		return domain.QuestDTO{}, errors.New("Get Error")
	}

	return domain.QuestDTO{Name: "test"}, nil
}

func (d *dummyRepo) CreateQuest(c *gin.Context, name string) error {
	if name == "testError" {
		return errors.New("GET ERROR")
	}
	return nil
}

func (d *dummyRepo) UpdateQuest(c *gin.Context, id string, quest domain.QuestDTO) error {
	if id == "error" {
		return errors.New("UPDATE ERROR")
	}
	return nil
}

func (d *dummyRepo) DeleteQuest(c *gin.Context, id string) error {
	if id == "error" {
		return errors.New("DELETE ERROR")
	}
	return nil
}

func (d *dummyRepo) GetQuestsCompletions(c *gin.Context, id int) ([]domain.QuestCompletion, error) {
	return nil, nil
}

func (d *dummyRepo) GetCompletion(c *gin.Context, questID int, userID int) (domain.QuestCompletion, error) {
	return domain.QuestCompletion{}, nil
}

func (d *dummyRepo) AddCompletion(c *gin.Context, questID int, userID int, startTime time.Time, endTime time.Time) error {
	return nil
}

func (d *dummyRepo) SaveCompletion(c *gin.Context, completion domain.QuestCompletion) error {
	return nil
}

func (d *dummyRepo) GetQuestInfo(c *gin.Context, questID int) (domain.QuestInfo, error) {
	return domain.QuestInfo{}, nil
}

func (d *dummyRepo) GetQuestInfoByName(c *gin.Context, questName string) (domain.QuestInfo, error) {
	return domain.QuestInfo{}, nil
}

func (d *dummyRepo) UpdateQuestInfo(c *gin.Context, quest domain.QuestInfo) error {
	return nil
}

func (d *dummyRepo) AddRating(c *gin.Context, rating domain.Rating) error {
	return nil
}

func (d *dummyRepo) GetQuestRatings(c *gin.Context, questID int) ([]*domain.Rating, error) {
	return nil, nil
}

func (d *dummyRepo) GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error) {
	return domain.Rating{}, nil
}

func NewDummyRepository() Repository {
	return &dummyRepo{}
}

func TestServiceGetShouldReturnOK(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo, nil)

	result, err := service.GetQuests(&gin.Context{})
	assert.Nil(t, err)
	assert.Equal(t, result[0].Name, "test")
}

func TestServiceCreateWithPostErrorShouldFail(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo, nil)

	err := service.CreateQuest(&gin.Context{}, "testError")
	assert.NotNil(t, err)
}

func TestServiceCreateShouldReturnOK(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo, nil)

	err := service.CreateQuest(&gin.Context{}, "test")
	assert.Nil(t, err)
}

func TestServiceUpdateK(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo, nil)

	err := service.UpdateQuest(&gin.Context{}, "62d5ea0c6c5608330d5d734b", domain.QuestDTO{Name: "test"})
	assert.Nil(t, err)
}
