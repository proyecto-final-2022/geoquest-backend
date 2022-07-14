package apitests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServerGame() *gin.Engine {
	repo := quest.NewRepository()
	service := quest.NewService(repo)
	handler := handler.NewGame(service)
	r := gin.Default()
	gGroup := r.Group("/quests")
	{
		gGroup.GET("/:id", handler.GetQuest())
	}

	return r
}

func TestGetGameShouldReturnOK(t *testing.T) {
	objReq := struct {
		ID int `json:"ID"`
	}{}

	r := createServerGame()

	req, rr := createRequestTest(http.MethodGet, "/quests/1", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.ID, 1)

}
