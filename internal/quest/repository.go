package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetQuests(c *gin.Context) ([]*domain.QuestDTO, error)
	GetQuest(c *gin.Context, id string) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, name string) error
	UpdateQuest(c *gin.Context, id string, quest domain.QuestDTO) error
	DeleteQuest(c *gin.Context, id string) error
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

func (r *repository) CreateQuest(c *gin.Context, name string) error {

	var err error

	_, err = collection.InsertOne(c, domain.Quest{Name: name})

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateQuest(c *gin.Context, id string, quest domain.QuestDTO) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name": quest.Name,
		},
	}

	_, err = collection.UpdateOne(c, filter, update)

	if err != nil {
		return err
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
