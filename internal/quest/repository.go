package quest

import (
	"errors"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	GetQuestInfo(c *gin.Context, questID int) (domain.QuestInfo, error)
	GetQuestInfoByName(c *gin.Context, questName string) (domain.QuestInfo, error)
	UpdateQuestInfo(c *gin.Context, quest domain.QuestInfo) error
	CreateQuest(c *gin.Context, id string, scene int, inventory []string) error
	UpdateQuest(c *gin.Context, quest domain.QuestDTO) error
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

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	result := collection.FindOne(c, filter)

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

func (r *repository) CreateQuest(c *gin.Context, id string, scene int, inventory []string) error {

	var err error

	_, err = collection.InsertOne(c, domain.Quest{QuestID: id, Scene: scene, Inventory: inventory})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateQuest(c *gin.Context, quest domain.QuestDTO) error {

	var err error

	//	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"questid": quest.QuestID}

	update := bson.M{
		"$set": bson.M{
			"quest_id":  quest.QuestID,
			"scene":     quest.Scene,
			"inventory": quest.Inventory,
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
