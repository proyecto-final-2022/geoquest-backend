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

func (d *dummyRepo) CreateCompletion(c *gin.Context, questID int, userID int, completedTime time.Time, hours float64, mins float64, segs float64) error {
	return nil
}

func NewDummyRepository() Repository {
	return &dummyRepo{}
}

func TestServiceGetShouldReturnOK(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.GetQuests(&gin.Context{})
	assert.Nil(t, err)
	assert.Equal(t, result[0].Name, "test")
}

func TestServiceCreateWithPostErrorShouldFail(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo)

	err := service.CreateQuest(&gin.Context{}, "testError")
	assert.NotNil(t, err)
}

func TestServiceCreateShouldReturnOK(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo)

	err := service.CreateQuest(&gin.Context{}, "test")
	assert.Nil(t, err)
}

func TestServiceUpdateK(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo)

	err := service.UpdateQuest(&gin.Context{}, "62d5ea0c6c5608330d5d734b", domain.QuestDTO{Name: "test"})
	assert.Nil(t, err)
}
