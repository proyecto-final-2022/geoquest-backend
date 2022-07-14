package user

import (
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Post(c *gin.Context, id int) (domain.User, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Post(c *gin.Context, id int) (domain.User, error) {
	return domain.User{ID: id}, nil
}
