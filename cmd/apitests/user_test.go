package apitests

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type dummyRepo struct{}

func (d *dummyRepo) Post(c *gin.Context, id int, email string, name string, password string) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("POST ERROR")
	}
	return domain.UserDTO{ID: id, Email: email, Name: name, Password: password}, nil
}

func (d *dummyRepo) Get(c *gin.Context, id int) (domain.UserDTO, error) {
	if id == 9 {
		return domain.UserDTO{}, errors.New("GET ERROR")
	}
	return domain.UserDTO{ID: id, Email: "test", Name: "test", Password: "test"}, nil
}

func createServerUser() *gin.Engine {

	repo := &dummyRepo{}
	service := user.NewService(repo)
	handler := handler.NewUser(service)

	r := gin.Default()

	gGroup := r.Group("/users")
	{
		gGroup.POST("/", handler.Post())
		gGroup.GET("/:id", handler.GetUser())
	}

	return r
}

func TestPostUserShouldReturnOK(t *testing.T) {
	objReq := struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

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

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.ID, 10)

}

func TestPostUserWithDBErrorShouldReturnInternalServerError(t *testing.T) {
	r := createServerUser()
	body := `
	{
		"user_id": 9,
		"email": "test",
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
		ID int `json:"id"`
	}{}

	r := createServerUser()

	req, rr := createRequestTest(http.MethodGet, "/users/10", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.ID, 10)

}

func TestGetUserWithDBErrorShouldReturnInternalServerError(t *testing.T) {
	r := createServerUser()

	req, rr := createRequestTest(http.MethodGet, "/users/9", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Code)

}
