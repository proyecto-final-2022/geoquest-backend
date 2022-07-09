package game

import (
	"errors"
	"geoquest-backend/internal/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) Get(c *gin.Context, id int) (domain.Game, error) {
	if id == 9 {
		return domain.Game{}, errors.New("GET ERROR")
	}
	return domain.Game{ID: id}, nil
}

func TestServiceGetWithGetErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.Get(&gin.Context{}, 9)
	assert.NotNil(t, err)
}

func TestServiceShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.Get(&gin.Context{}, 1)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}
