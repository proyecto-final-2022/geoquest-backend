package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetQuests(c *gin.Context, id int) ([]*domain.QuestDTO, error)
	CreateQuest(c *gin.Context, name string) error
}

type repository struct {
}

var collection = config.GetCollection("quest")

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetQuests(c *gin.Context, id int) ([]*domain.QuestDTO, error) {
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
