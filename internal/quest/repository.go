package quest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"

	"gorm.io/datatypes"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	GetQuestInfo(c *gin.Context, questID int) (domain.QuestInfo, error)
	GetQuestInfoByName(c *gin.Context, questName string) (domain.QuestInfo, error)
	UpdateQuestInfo(c *gin.Context, quest domain.QuestInfo) error
	CreateQuest(c *gin.Context, id string, scene int, inventory []string, logs []string, points float64) error
	CreateQuestProgression(c *gin.Context, id int, teamId int, scene int, inventory []string, logs []string, objects map[string]int, points float64) error
	GetQuestProgression(c *gin.Context, id int, teamId int) (datatypes.JSON, error)
	UpdateQuestProgression(c *gin.Context, id int, scene int, inventory []string, logs []string, objects map[string]int, points float64) error
	UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error
	DeleteQuest(c *gin.Context, id string) error
	GetQuestsCompletions(c *gin.Context, questID int) ([]domain.QuestCompletion, error)
	GetCompletion(c *gin.Context, questID int, userID int) (domain.QuestCompletion, error)
	AddCompletion(c *gin.Context, questID int, userID int, startTime time.Time, endTime time.Time) error
	GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error)
	AddRating(c *gin.Context, rating domain.Rating) error
	GetQuestRatings(c *gin.Context, questID int) ([]*domain.Rating, error)
	SaveCompletion(c *gin.Context, completion domain.QuestCompletion) error
}

type repository struct {
}

var collection = config.GetCollection("quest")

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetQuest(c *gin.Context, id string) (domain.QuestDTO, error) {
	var quest domain.QuestDTO

	filter := bson.M{"questid": id}

	result := collection.FindOne(c, filter)

	fmt.Println("result: ", result)

	if err := result.Decode(&quest); err != nil {
		return domain.QuestDTO{}, err
	}

	return quest, nil
}

func (r *repository) GetQuests(c *gin.Context) ([]*domain.QuestDTO, error) {
	var quests []*domain.QuestDTO

	filter := bson.D{}

	coll, err := collection.Find(c, filter)

	if err != nil {
		return nil, err
	}

	for coll.Next(c) {
		var quest domain.QuestDTO
		if err = coll.Decode(&quest); err != nil {
			return nil, err
		}
		quests = append(quests, &quest)
	}

	return quests, nil
}

func (r *repository) CreateQuest(c *gin.Context, id string, scene int, inventory []string, logs []string, points float64) error {

	var err error

	_, err = collection.InsertOne(c, domain.Quest{QuestID: id, Scene: scene, Inventory: inventory, Logs: logs, Points: points})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateQuestProgression(c *gin.Context, id int, teamId int, scene int, inventory []string, logs []string, objects map[string]int, points float64) error {

	questInfo := map[string]interface{}{
		"quest_id":  id,
		"team_id":   teamId,
		"scene":     scene,
		"inventory": inventory,
		"logs":      logs,
		"points":    points,
		"objects":   objects,
	}

	jsonQuest, _ := json.Marshal(questInfo)

	questProgress := domain.QuestProgress{QuestID: id, TeamID: teamId, Info: datatypes.JSON(string(jsonQuest))}
	if tx := config.MySql.Create(&questProgress); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) GetQuestProgression(c *gin.Context, id int, teamId int) (datatypes.JSON, error) {

	var questProgress domain.QuestProgress
	if tx := config.MySql.Where("quest_id = ? AND team_id = ?", id, teamId).First(&questProgress); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return questProgress.Info, nil
}

func (r *repository) UpdateQuestProgression(c *gin.Context, id int, scene int, inventory []string, logs []string, objects map[string]int, points float64) error {
	questInfo := map[string]interface{}{
		"quest_id":  id,
		"scene":     scene,
		"inventory": inventory,
		"logs":      logs,
		"points":    points,
		"objects":   objects,
	}

	jsonQuest, _ := json.Marshal(questInfo)

	questProgress := domain.QuestProgress{Info: datatypes.JSON(string(jsonQuest))}

	if tx := config.MySql.Save(&questProgress).Where("quest_id = ?", id); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *repository) UpdateQuest(c *gin.Context, quest domain.QuestDTO, paramId string) error {

	var err error

	//	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"questid": paramId}

	update := bson.M{
		"$set": bson.M{
			"scene":     quest.Scene,
			"inventory": quest.Inventory,
			"objects":   quest.Objects,
			"logs":      quest.Logs,
			"points":    quest.Points,
		},
	}

	_, err = collection.UpdateOne(c, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetQuestInfo(c *gin.Context, questID int) (domain.QuestInfo, error) {
	var quest domain.QuestInfo
	if tx := config.MySql.Where("id = ?", questID).First(&quest); tx.Error != nil {
		return domain.QuestInfo{}, errors.New("DB Error")
	}
	return quest, nil
}

func (r *repository) GetQuestInfoByName(c *gin.Context, questName string) (domain.QuestInfo, error) {
	var quest domain.QuestInfo
	if tx := config.MySql.Where("name = ?", questName).First(&quest); tx.Error != nil {
		return domain.QuestInfo{}, errors.New("DB Error")
	}
	return quest, nil
}

func (r *repository) UpdateQuestInfo(c *gin.Context, quest domain.QuestInfo) error {
	if tx := config.MySql.Save(&quest); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *repository) GetCompletion(c *gin.Context, questID int, userID int) (domain.QuestCompletion, error) {
	var completion domain.QuestCompletion
	if tx := config.MySql.Where("user_id = ? AND quest_id = ?", userID, questID).First(&completion); tx.Error != nil {
		return domain.QuestCompletion{}, errors.New("DB Error")
	}
	return completion, nil
}

func (r *repository) AddCompletion(c *gin.Context, questID int, userID int, startTime time.Time, endTime time.Time) error {
	completionSave := domain.QuestCompletion{QuestID: questID, UserID: userID, StartTime: startTime, EndTime: endTime}
	if tx := config.MySql.Create(&completionSave); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) GetRating(c *gin.Context, questID int, userID int) (domain.Rating, error) {
	var rating domain.Rating
	if tx := config.MySql.Where("user_id = ? AND quest_id = ?", userID, questID).First(&rating); tx.Error != nil {
		return domain.Rating{}, errors.New("DB Error")
	}
	return rating, nil
}

func (r *repository) AddRating(c *gin.Context, rating domain.Rating) error {
	if tx := config.MySql.Save(&rating); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) GetQuestRatings(c *gin.Context, questID int) ([]*domain.Rating, error) {
	var ratings []*domain.Rating

	if tx := config.MySql.Where("quest_id = ?", questID).Find(&ratings); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return ratings, nil
}

func (r *repository) SaveCompletion(c *gin.Context, completion domain.QuestCompletion) error {
	if tx := config.MySql.Save(&completion); tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *repository) DeleteQuest(c *gin.Context, id string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(c, filter)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetQuestsCompletions(c *gin.Context, questID int) ([]domain.QuestCompletion, error) {
	var questsCompletions []domain.QuestCompletion
	if tx := config.MySql.Where("quest_id = ?", questID).Find(&questsCompletions); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return questsCompletions, nil
}
