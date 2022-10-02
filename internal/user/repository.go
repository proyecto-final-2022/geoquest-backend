package user

import (
	"errors"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	CreateUser(c *gin.Context, email string, name string, username string, image int, manual bool, google bool, facebook bool, password string) error
	GetUser(c *gin.Context, id int) (domain.UserDTO, domain.User, error)
	GetUsers(c *gin.Context) ([]domain.User, error)
	GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error)
	UpdateUser(c *gin.Context, user domain.User) error
	UpdateCoupon(c *gin.Context, coupon domain.Coupon) error
	GetCoupon(c *gin.Context, couponID int) (domain.Coupon, error)
	DeleteUser(c *gin.Context, id int) error
	CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error
	GetCoupons(c *gin.Context, userID int) ([]domain.Coupon, error)
	AddFriend(c *gin.Context, id int, friendID int) error
	GetUserFriends(c *gin.Context, id int) ([]domain.UserFriends, error)
	DeleteFriend(c *gin.Context, id int, friendID int) error
	AddNotification(c *gin.Context, ID int, senderID int, notificationType string, questName string, teamID int, questID int, actualTime time.Time) error
	GetNotifications(c *gin.Context, userID int) ([]domain.Notification, error)
	DeleteNotification(c *gin.Context, id int, notificationID int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateUser(c *gin.Context, email string, name string, username string, image int, manual bool, google bool, facebook bool, password string) error {
	user := domain.User{Email: email, Name: name, Username: username, Image: image, Manual: manual, Google: google, Facebook: facebook, Password: password}
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
	return domain.UserDTO{ID: id, Username: user.Username, Email: user.Email, Name: user.Name, Password: user.Password, Manual: user.Manual, Google: user.Google, Facebook: user.Facebook, Image: user.Image}, user, nil
}

func (r *repository) GetUsers(c *gin.Context) ([]domain.User, error) {
	var users []domain.User
	if tx := config.MySql.Find(&users); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return users, nil
}

func (r *repository) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	var user domain.User
	if tx := config.MySql.Where("email = ?", email).First(&user); tx.Error != nil {
		return domain.UserDTO{}, errors.New("DB Error")
	}
	return domain.UserDTO{ID: user.ID, Email: user.Email, Name: user.Name, Username: user.Username, Password: user.Password, Manual: user.Manual, Google: user.Google, Facebook: user.Facebook, Image: user.Image}, nil
}

func (r *repository) UpdateUser(c *gin.Context, user domain.User) error {

	if tx := config.MySql.Save(&user); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *repository) UpdateCoupon(c *gin.Context, coupon domain.Coupon) error {

	if tx := config.MySql.Save(&coupon); tx.Error != nil {
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

func (r *repository) GetCoupon(c *gin.Context, couponID int) (domain.Coupon, error) {
	var coupon domain.Coupon
	if tx := config.MySql.Where("id = ?", couponID).Find(&coupon); tx.Error != nil {
		return domain.Coupon{}, errors.New("DB Error")
	}
	return coupon, nil
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

func (r *repository) DeleteNotification(c *gin.Context, userID int, notificationID int) error {
	if tx := config.MySql.Where("receiver_id = ? AND id = ?", userID, notificationID).Delete(&domain.Notification{}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *repository) AddNotification(c *gin.Context, userID int, senderID int, notificationType string, questName string, teamID int, questID int, actualTime time.Time) error {

	notification := domain.Notification{SenderID: senderID, ReceiverID: userID, Type: notificationType, SentTime: actualTime, QuestName: questName, TeamID: teamID, QuestID: questID}

	if tx := config.MySql.Create(&notification); tx.Error != nil {
		return errors.New("DB Error")
	}

	return nil

}

func (r *repository) GetNotifications(c *gin.Context, userID int) ([]domain.Notification, error) {
	var notifications []domain.Notification
	if tx := config.MySql.Where("receiver_id = ?", userID).Find(&notifications); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return notifications, nil
}
