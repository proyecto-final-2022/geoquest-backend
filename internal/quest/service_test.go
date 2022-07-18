package quest

import (
	"errors"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) Get(c *gin.Context, id int) (domain.Quest, error) {
	if id == 9 {
		return domain.Quest{}, errors.New("GET ERROR")
	}
	return domain.Quest{}, nil
}

func (d *dummyRepo) Post(c *gin.Context, name string) (domain.Quest, error) {
	return domain.Quest{}, nil
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
