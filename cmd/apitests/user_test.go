package apitests

import (
	"encoding/json"
	"geoquest-backend/cmd/api/handler"
	"geoquest-backend/internal/user"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerUser() *gin.Engine {

	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := handler.NewUser(service)

	r := gin.Default()

	gGroup := r.Group("/users")
	{
		gGroup.POST("/", handler.Post())
	}

	return r
}

func TestPostUserShouldReturnOK(t *testing.T) {
	objReq := struct {
		ID int `json:"ID"`
	}{}

	r := createServerUser()

	body := `
	{
		"user_id": 6
	}`

	req, rr := createRequestTest(http.MethodPost, "/users/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.ID, 6)

}
