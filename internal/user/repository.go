package user

import (
	"errors"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error)
	Get(c *gin.Context, id int) (domain.UserDTO, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error) {
	user := domain.User{ID: id, Email: email, Name: name, Password: password}
	if tx := config.MySql.Create(&user); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{ID: id, Email: email, Name: name, Password: password}, nil
}

func (r *repository) Get(c *gin.Context, id int) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.First(&user, id); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{ID: user.ID, Email: user.Email, Name: user.Name, Password: user.Password}, nil
}
