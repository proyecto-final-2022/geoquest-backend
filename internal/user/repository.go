package user

import (
	"geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Get(c *gin.Context, id int) (domain.User, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Get(c *gin.Context, id int) (domain.User, error) {
	return domain.User{ID: id}, nil
}
