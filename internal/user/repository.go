package user

import (
	"errors"
	"fmt"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, email string, name string, username string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, id int, user domain.UserDTO) error
	DeleteUser(c *gin.Context, id int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateUser(c *gin.Context, email string, name string, username string, password string) error {
	user := domain.User{Email: email, Name: name, Username: username, Password: password}
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

func (r *repository) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.Where("email = ?", email).First(&user); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	fmt.Println(user.Username)
	return domain.UserDTO{Email: user.Email, Name: user.Name, Username: user.Username, Password: user.Password}, nil
}

func (r *repository) UpdateUser(c *gin.Context, id int, user domain.UserDTO) error {
	var userUpd domain.User
	if tx := config.MySql.First(&userUpd, id); tx.Error != nil {
		return errors.New("DB Error")
	}

	userUpd.Name = user.Name
	userUpd.Email = user.Email
	userUpd.Password = user.Password

	if tx := config.MySql.Save(&userUpd); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *repository) DeleteUser(c *gin.Context, id int) error {
	if tx := config.MySql.Delete(&domain.User{}, id); tx.Error != nil {
		return tx.Error
	}
	return nil
}
