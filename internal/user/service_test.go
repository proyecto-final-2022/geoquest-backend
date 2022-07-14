package user

import (
	"errors"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("POST ERROR")
	}
	return domain.UserDTO{ID: id, Email: email, Name: name, Password: password}, nil
}

func (d *dummyRepo) Get(c *gin.Context, id int) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("POST ERROR")
	}
	return domain.UserDTO{ID: id}, nil
}

func TestServicePostWithGetErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.Post(&gin.Context{}, 9, "test", "test", "test")
	assert.NotNil(t, err)
}

func TestServicePostShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.Post(&gin.Context{}, 1, "test", "test", "test")
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestServiceGetShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.Get(&gin.Context{}, 10)
	assert.Nil(t, err)
	assert.Equal(t, result.ID, 10)
}

func TestServiceGetWithDBErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.Get(&gin.Context{}, 9)
	assert.NotNil(t, err)
}
