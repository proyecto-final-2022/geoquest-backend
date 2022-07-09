package game

import (
	"geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Get(c *gin.Context, id int) (domain.Game, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Get(c *gin.Context, id int) (domain.Game, error) {
	return domain.Game{ID: id}, nil
}
