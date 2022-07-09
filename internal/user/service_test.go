package user

import (
	"errors"
	"geoquest-backend/internal/domain"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) Post(c *gin.Context, id int) (domain.User, error) {
	if id == 9 {
		return domain.User{}, errors.New("POST ERROR")
	}
	return domain.User{ID: id}, nil
}

func TestServicePostWithGetErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.Post(&gin.Context{}, 9)
	assert.NotNil(t, err)
}

func TestServiceShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.Post(&gin.Context{}, 1)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}
