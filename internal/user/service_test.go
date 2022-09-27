package user

import (
	"errors"
	"testing"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) CreateUser(c *gin.Context, email string, name string, username string, image int, manual bool, google bool, facebook bool, password string) error {
	if email == "testError" {
		return errors.New("POST ERROR")
	}
	return nil
}

func (d *dummyRepo) GetUser(c *gin.Context, id int) (domain.UserDTO, domain.User, error) {
	if id == 9 {
		return domain.UserDTO{}, domain.User{}, errors.New("POST ERROR")
	}
	return domain.UserDTO{Name: "test", Password: "test", Email: "test"}, domain.User{}, nil
}

func (d *dummyRepo) UpdateUser(c *gin.Context, user domain.User) error {
	return nil
}

func (d *dummyRepo) DeleteUser(c *gin.Context, id int) error {
	if id == 9 {
		return errors.New("GET ERROR")
	}
	return nil
}

func (d *dummyRepo) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	return domain.UserDTO{}, nil
}

func (d *dummyRepo) CreateCoupon(c *gin.Context, userID int, description string, date time.Time) error {
	return nil
}

func (d *dummyRepo) GetCoupons(c *gin.Context, userID int) ([]domain.Coupon, error) {
	return nil, nil
}

func (d *dummyRepo) AddFriend(c *gin.Context, userID int, friendID int) error {
	return nil

}

func (d *dummyRepo) DeleteFriend(c *gin.Context, userID int, friendID int) error {
	return nil
}

func (d *dummyRepo) GetUserFriends(c *gin.Context, userID int) ([]domain.UserFriends, error) {
	return nil, nil
}

func (d *dummyRepo) AddNotification(c *gin.Context, ID int, senderID int, notificationType string, questName string, teamID int, questID int, actualTime time.Time) error {
	return nil
}

func (d *dummyRepo) GetNotifications(c *gin.Context, userID int) ([]domain.Notification, error) {
	return nil, nil
}

func (d *dummyRepo) DeleteNotification(c *gin.Context, userID int, notificationID int) error {
	return nil
}

func TestServicePostWithGetErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	err := service.CreateUser(&gin.Context{}, "testError", "test", "test", "test")
	assert.NotNil(t, err)
}

func TestServicePostShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	err := service.CreateUser(&gin.Context{}, "test", "test", "test", "test")
	assert.Nil(t, err)
}

func TestServiceGetShouldSuccess(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	result, err := service.GetUser(&gin.Context{}, 10)
	assert.Nil(t, err)
	assert.Equal(t, result.Name, "test")
}

func TestServiceGetWithDBErrorShouldFail(t *testing.T) {
	repo := &dummyRepo{}
	service := NewService(repo)

	_, err := service.GetUser(&gin.Context{}, 9)
	assert.NotNil(t, err)
}
