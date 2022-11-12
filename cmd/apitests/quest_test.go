package apitests

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type serviceMock struct{}

func (s *serviceMock) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {

	var quests []*domain.QuestDTO

	quests = append(quests, &domain.QuestDTO{QuestID: "test"})

	return quests, nil
}

func (s *serviceMock) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	if id == "9" {
		return domain.QuestDTO{}, errors.New("Get Error")
	}

	return domain.QuestDTO{QuestID: "test"}, nil
}

func (s *serviceMock) CreateQuest(c *gin.Context, id string, scene int, inventory []string, logs []string, points float64) error {
	if id == "testError" {
		return errors.New("GET ERROR")
	}
	return nil
}

func (s *serviceMock) UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error {
	return nil
}

func (s *serviceMock) DeleteQuest(c *gin.Context, id string) error {
	if id == "error" {
		return errors.New("DELETE ERROR")
	}
	return nil
}

func (s *serviceMock) CreateCompletion(c *gin.Context, questID int, userID int, startYear int, startMonth time.Month,
	startDay int, startHour int, startMinutes int, startSeconds int) error {
	return nil
}

func (s *serviceMock) GetRanking(c *gin.Context, id int) ([]domain.QuestCompletionDTO, error) {
	return nil, nil
}

func (s *serviceMock) CreateRating(c *gin.Context, questID int, userID int, rating int) error {
	return nil
}

func (s *serviceMock) GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error) {
	return domain.Rating{}, nil
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
		gGroup.GET("/", handler.GetQuests())
		gGroup.GET("/:id", handler.GetQuest())
		gGroup.POST("/", handler.CreateQuest())
		gGroup.PUT("/:id", handler.UpdateQuest())

	}

	return r
}

func TestGetQuestsShouldReturnOK(t *testing.T) {
	r := createServerGame()

	req, rr := createRequestTest(http.MethodGet, fmt.Sprintf("/quests/"), "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func TestMongolo(t *testing.T) {
	r := createServerGame()

	body := `
	{
		"name": "struung"
	}`

	req, rr := createRequestTest(http.MethodPut, fmt.Sprintf("/quests/23"), body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

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

func TestCreateQuestShouldReturnInternalServerError(t *testing.T) {
	body := `
	{
		"name": "testError"
	}`

	r := createServerGame()

	req, rr := createRequestTest(http.MethodPost, "/quests/", body)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 500, rr.Code)

}
