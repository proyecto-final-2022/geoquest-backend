package user

import (
	"errors"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) CreateUser(c *gin.Context, email string, name string, username string, password string) error {
	if email == "testError" {
		return errors.New("POST ERROR")
	}
	return nil
}

func (d *dummyRepo) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("POST ERROR")
	}
	return domain.UserDTO{Name: "test", Password: "test", Email: "test"}, nil
}

func (d *dummyRepo) UpdateUser(c *gin.Context, id int, user domain.UserDTO) error {
	if id == 9 {
		return errors.New("GET ERROR")
	}
	return nil
}

func (d *dummyRepo) DeleteUser(c *gin.Context, id int) error {
	if id == 9 {
		return errors.New("GET ERROR")
	}
	return nil
}

func (d *dummyRepo) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	return domain.UserDTO{}, nil
}

func TestServicePostWithGetErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	err := service.CreateUser(&gin.Context{}, "testError", "test", "test", "test")
	assert.NotNil(t, err)
}

func TestServicePostShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	err := service.CreateUser(&gin.Context{}, "test", "test", "test", "test")
	assert.Nil(t, err)
}

func TestServiceGetShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.GetUser(&gin.Context{}, 10)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, "test")
}

func TestServiceGetWithDBErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.GetUser(&gin.Context{}, 9)
	assert.NotNil(t, err)
}
