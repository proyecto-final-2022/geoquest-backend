package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Get(c *gin.Context, id int) (domain.Quest, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Get(c *gin.Context, id int) (domain.Quest, error) {
	return domain.Quest{ID: id}, nil
}
