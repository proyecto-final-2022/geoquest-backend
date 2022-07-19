package apitests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type serviceMock struct{}

func (s *serviceMock) GetQuest(c *gin.Context, id int) (domain.QuestDTO, error) {
	if id == 9 {
		return domain.QuestDTO{}, errors.New("GET ERROR")
	}
	return domain.QuestDTO{Name: "test"}, nil
}

func (s *serviceMock) CreateQuest(c *gin.Context, name string) (domain.QuestDTO, error) {
	return domain.QuestDTO{}, nil
}

func NewServiceMock() quest.Service {
	return &serviceMock{}
}

func createServerGame() *gin.Engine {
	service := NewServiceMock()
	handler := handler.NewGame(service)
	r := gin.Default()
	gGroup := r.Group("/quests")
	{
		gGroup.GET("/:id", handler.GetQuest())
		gGroup.POST("/", handler.CreateQuest())
	}

	return r
}

func TestGetQuestShouldReturnOK(t *testing.T) {
	id := "1"
	objReq := struct {
		Name string `json:"name"`
	}{}

	r := createServerGame()

	req, rr := createRequestTest(http.MethodGet, fmt.Sprintf("/quests/%s", id), "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

	err := json.Unmarshal(rr.Body.Bytes(), &objReq)
	assert.Nil(t, err)
	assert.Equal(t, objReq.Name, "test")

}

func TestGetQuestShouldReturnInternalServerError(t *testing.T) {
	id := "9"

	r := createServerGame()

	req, rr := createRequestTest(http.MethodGet, fmt.Sprintf("/quests/%s", id), "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Code)

}

func TestCreateQuestShouldReturnOK(t *testing.T) {
	body := `
	{
		"name": "string"
	}`

	r := createServerGame()

	req, rr := createRequestTest(http.MethodPost, "/quests/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}
