package quest

import (
	"errors"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) GetQuests(c *gin.Context, id int) ([]*domain.QuestDTO, error) {

	if id == 9 {
		return nil, errors.New("GET ERROR")
	}

	var quests []*domain.QuestDTO

	quests = append(quests, &domain.QuestDTO{Name: "test"})

	return quests, nil
}

func (d *dummyRepo) CreateQuest(c *gin.Context, name string) error {
	if name == "testError" {
		return errors.New("GET ERROR")
	}
	return nil
}

func NewDummyRepository() Repository {
	return &dummyRepo{}
}

func TestServiceGetWithGetErrorShouldFail(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo)

	_, err := service.GetQuests(&gin.Context{}, 9)
	assert.NotNil(t, err)
}

func TestServiceGetShouldReturnOK(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.GetQuests(&gin.Context{}, 1)
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
