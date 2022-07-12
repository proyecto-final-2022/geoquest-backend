package apitests

import (
	"encoding/json"
	"geoquest-backend/cmd/api/handler"
	"geoquest-backend/internal/quest"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerGame() *gin.Engine {
	repo := quest.NewRepository()
	service := quest.NewService(repo)
	handler := handler.NewGame(service)
	r := gin.Default()
	gGroup := r.Group("/games")
	{
		gGroup.GET("/:id", handler.Get())
	}

	return r
}

func TestGetGameShouldReturnOK(t *testing.T) {
	objReq := struct {
		ID int `json:"ID"`
	}{}

	r := createServerGame()

	req, rr := createRequestTest(http.MethodGet, "/games/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.ID, 1)

}
