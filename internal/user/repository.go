package user

import (
	"errors"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, email string, name string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateUser(c *gin.Context, email string, name string, password string) error {
	user := domain.User{Email: email, Name: name, Password: password}
	if tx := config.MySql.Create(&user); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.First(&user, id); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{Email: user.Email, Name: user.Name, Password: user.Password}, nil
}
