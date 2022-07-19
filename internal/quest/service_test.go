package quest

import (
	"errors"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) GetQuest(c *gin.Context, id int) (domain.QuestDTO, error) {
	if id == 9 {
		return domain.QuestDTO{}, errors.New("GET ERROR")
	}
	return domain.QuestDTO{}, nil
}

func (d *dummyRepo) CreateQuest(c *gin.Context, name string) (domain.QuestDTO, error) {
	return domain.QuestDTO{}, nil
}

func NewDummyRepository() Repository {
	return &dummyRepo{}
}

func TestServiceGetWithGetErrorShouldFail(t *testing.T) {
	repo := NewDummyRepository()
	service := NewService(repo)

	_, err := service.GetQuest(&gin.Context{}, 9)
	assert.NotNil(t, err)
}

/*
func TestServiceShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.Get(&gin.Context{}, 1)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}
*/
