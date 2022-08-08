package user

import (
	"errors"
	"fmt"
	"time"

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
	CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error
	GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error)
	AddFriend(c *gin.Context, id int, friendID int) error
	GetUserFriends(c *gin.Context, id int) ([]domain.UserDTO, error)
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

func (r *repository) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.First(&user, id); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{ID: id, Email: user.Email, Name: user.Name, Password: user.Password}, nil
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

func (r *repository) CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error {

	coupon := domain.Coupon{UserID: userID, Description: description, ExpirationDate: date}

	if tx := config.MySql.Create(&coupon); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil

}

func (r *repository) GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error) {

	var coupons []domain.Coupon
	if tx := config.MySql.Where("user_id = ?", userID).Find(&coupons); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	couponsDTO := make([]domain.CouponDTO, len(coupons))
	for i := range coupons {
		couponsDTO[i].ID = coupons[i].ID
		couponsDTO[i].Description = coupons[i].Description
		couponsDTO[i].ExpirationDate = coupons[i].ExpirationDate
		couponsDTO[i].Used = coupons[i].Used
	}

	return couponsDTO, nil
}

func (r *repository) AddFriend(c *gin.Context, userID int, friendID int) error {

	userFriend := domain.UserFriends{UserID: userID, FriendID: friendID}

	if tx := config.MySql.Create(&userFriend); tx.Error != nil {
		return errors.New("DB Error")
	}

	return nil

}

func (r *repository) GetUserFriends(c *gin.Context, userID int) ([]domain.UserDTO, error) {

	var users []domain.UserFriends
	if tx := config.MySql.Where("user_id = ?", userID).Find(&users); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	usersDTO := make([]domain.UserDTO, len(users))

	for i := range users {

		var user domain.User
		if tx := config.MySql.Where("id = ?", users[i].FriendID).First(&user); tx.Error != nil {
			return nil, errors.New("DB Error")
		}

		usersDTO[i].ID = user.ID
		usersDTO[i].Name = user.Name
		usersDTO[i].Username = user.Username
		usersDTO[i].Email = user.Email
		usersDTO[i].Password = user.Password
	}

	return usersDTO, nil
}

func (r *repository) DeleteFriend(c *gin.Context, userID int, friendID int) error {
	if tx := config.MySql.Where("user_id = ? AND friend_id = ?", userID, friendID).Delete(&domain.UserFriends{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}
