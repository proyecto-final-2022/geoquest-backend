package apitests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (s *serviceMock) CreateUser(c *gin.Context, email string, name string, username string, password string) error {
	if email == "testError" {
		return errors.New("POST ERROR")
	}
	return nil
}

func (s *serviceMock) GetUser(c *gin.Context, id int) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("GET ERROR")
	}
	return domain.UserDTO{Email: "test", Name: "test", Password: "test"}, nil
}

func (s *serviceMock) UpdateUser(c *gin.Context, id int, email string, name string, password string, username string) error {
	if id == 9 {
		return errors.New("GET ERROR")
	}
	return nil
}

func (s *serviceMock) DeleteUser(c *gin.Context, id int) error {
	if id == 9 {
		return errors.New("GET ERROR")
	}
	return nil
}

func (s *serviceMock) HashPassword(password string) (string, error) {
	return "", nil
}

func (s *serviceMock) GetUserByEmail(c *gin.Context, email string) (domain.UserDTO, error) {
	return domain.UserDTO{}, nil
}

func (s *serviceMock) CheckPassword(providedPassword string, userPassword string) error {
	return nil
}

func (s *serviceMock) CreateCoupon(c *gin.Context, userID int, description string, expirationYear int, expirationMonth time.Month, expirationDay int, expirationHour int) error {
	return nil
}

func (s *serviceMock) GetCoupons(c *gin.Context, userID int) ([]domain.CouponDTO, error) {
	return nil, nil
}

func (s *serviceMock) AddFriend(c *gin.Context, userID int, friendID int) error {
	return nil
}

func (s *serviceMock) DeleteFriend(c *gin.Context, userID int, friendID int) error {
	return nil
}

func (s *serviceMock) GetUserFriends(c *gin.Context, userID int) ([]domain.UserDTO, error) {
	return nil, nil
}

func (s *serviceMock) AddNotification(c *gin.Context, userID int, senderID int, notificationType string) error {
	return nil
}

func (s *serviceMock) GetNotifications(c *gin.Context, userID int) ([]domain.NotificationDTO, error) {
	return nil, nil
}

func createServerUser() *gin.Engine {

	service := &serviceMock{}
	handler := handler.NewUser(service)

	r := gin.Default()

	gGroup := r.Group("/users")
	{
		gGroup.POST("/", handler.CreateUser())
		gGroup.GET("/:id", handler.GetUser())
	}

	return r
}

func TestPostUserShouldReturnOK(t *testing.T) {

	r := createServerUser()

	body := `
	{
		"user_id": 10,
		"email": "test",
		"password": "test",
		"name": "test"
	}`

	req, rr := createRequestTest(http.MethodPost, "/users/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func TestPostUserWithDBErrorShouldReturnInternalServerError(t *testing.T) {
	r := createServerUser()
	body := `
	{
		"user_id": 9,
		"email": "testError",
		"password": "test",
		"name": "test"
	}`

	req, rr := createRequestTest(http.MethodPost, "/users/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Code)
}

func TestPostUserWithInvalidJSONShouldReturnStatusUnprocessableEntity(t *testing.T) {
	r := createServerUser()
	body := `ble`

	req, rr := createRequestTest(http.MethodPost, "/users/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 422, rr.Code)
}

func TestGetUserShouldReturnOK(t *testing.T) {
	objReq := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	r := createServerUser()

	req, rr := createRequestTest(http.MethodGet, "/users/10", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Name, "test")
	assert.Equal(t, objReq.Email, "test")
	assert.Equal(t, objReq.Password, "test")

}

func TestGetUserWithDBErrorShouldReturnInternalServerError(t *testing.T) {
	r := createServerUser()

	id := "9"

	req, rr := createRequestTest(http.MethodGet, fmt.Sprintf("/users/%s", id), "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Code)

}
