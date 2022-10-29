package quest

import (
	"errors"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type dummyRepo struct{}

func (d *dummyRepo) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {

	var quests []*domain.QuestDTO

	quests = append(quests, &domain.QuestDTO{QuestID: "test"})

	return quests, nil
}

func (d *dummyRepo) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	if id == "9" {
		return domain.QuestDTO{}, errors.New("Get Error")
	}

	return domain.QuestDTO{QuestID: "test"}, nil
}

func (d *dummyRepo) CreateQuest(c *gin.Context, id string, scene int, inventory []string) error {
	if id == "testError" {
		return errors.New("GET ERROR")
	}
	return nil
}

func (d *dummyRepo) UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error {
	if quest.QuestID == "error" {
		return errors.New("UPDATE ERROR")
	}
	return nil
}

func (d *dummyRepo) DeleteQuest(c *gin.Context, id string) error {
	if id == "error" {
		return errors.New("DELETE ERROR")
	}
	return nil
}

func (d *dummyRepo) GetQuestsCompletions(c *gin.Context, id int) ([]domain.QuestCompletion, error) {
	return nil, nil
}

func (d *dummyRepo) GetCompletion(c *gin.Context, questID int, userID int) (domain.QuestCompletion, error) {
	return domain.QuestCompletion{}, nil
}

func (d *dummyRepo) AddCompletion(c *gin.Context, questID int, userID int, startTime time.Time, endTime time.Time) error {
	return nil
}

func (d *dummyRepo) SaveCompletion(c *gin.Context, completion domain.QuestCompletion) error {
	return nil
}

func (d *dummyRepo) GetQuestInfo(c *gin.Context, questID int) (domain.QuestInfo, error) {
	return domain.QuestInfo{}, nil
}

func (d *dummyRepo) GetQuestInfoByName(c *gin.Context, questName string) (domain.QuestInfo, error) {
	return domain.QuestInfo{}, nil
}

func (d *dummyRepo) UpdateQuestInfo(c *gin.Context, quest domain.QuestInfo) error {
	return nil
}

func (d *dummyRepo) AddRating(c *gin.Context, rating domain.Rating) error {
	return nil
}

func (d *dummyRepo) GetQuestRatings(c *gin.Context, questID int) ([]*domain.Rating, error) {
	return nil, nil
}

func (d *dummyRepo) GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error) {
	return domain.Rating{}, nil
}

func NewDummyRepository() Repository {
	return &dummyRepo{}
}
