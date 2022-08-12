package user

import (
	"errors"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, email string, name string, username string, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, domain.User, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, user domain.User) error
	DeleteUser(c *gin.Context, id int) error
	CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error
	GetCoupons(c *gin.Context, userID int) ([]domain.Coupon, error)
	AddFriend(c *gin.Context, id int, friendID int) error
	GetUserFriends(c *gin.Context, id int) ([]domain.UserFriends, error)
	DeleteFriend(c *gin.Context, id int, friendID int) error
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

func (r *repository) GetUser(c *gin.Context, id int) (domain.UserDTO, domain.User, error) {
	var user domain.User
	if tx := config.MySql.First(&user, id); tx.Error != nil {
		return domain.UserDTO{}, domain.User{}, errors.New("DB Error")
	}
	return domain.UserDTO{ID: id, Email: user.Email, Name: user.Name, Password: user.Password}, user, nil
}

func (r *repository) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.Where("email = ?", email).First(&user); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{Email: user.Email, Name: user.Name, Username: user.Username, Password: user.Password}, nil
}

func (r *repository) UpdateUser(c *gin.Context, user domain.User) error {

	if tx := config.MySql.Save(&user); tx.Error != nil {
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

func (r *repository) CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error {

	coupon := domain.Coupon{UserID: userID, Description: description, ExpirationDate: date}

	if tx := config.MySql.Create(&coupon); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil

}

func (r *repository) GetCoupons(c *gin.Context, userID int) ([]domain.Coupon, error) {
	var coupons []domain.Coupon
	if tx := config.MySql.Where("user_id = ?", userID).Find(&coupons); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return coupons, nil
}

func (r *repository) AddFriend(c *gin.Context, userID int, friendID int) error {

	userFriend := domain.UserFriends{UserID: userID, FriendID: friendID}

	if tx := config.MySql.Create(&userFriend); tx.Error != nil {
		return errors.New("DB Error")
	}

	return nil

}

func (r *repository) GetUserFriends(c *gin.Context, userID int) ([]domain.UserFriends, error) {
	var users []domain.UserFriends
	if tx := config.MySql.Where("user_id = ?", userID).Find(&users); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return users, nil
}

func (r *repository) DeleteFriend(c *gin.Context, userID int, friendID int) error {
	if tx := config.MySql.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&domain.UserFriends{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}
