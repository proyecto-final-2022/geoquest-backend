package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetQuest(c *gin.Context, id int) (domain.QuestDTO, error)
	CreateQuest(c *gin.Context, name string) (domain.QuestDTO, error)
}

type repository struct {
}

var collection = config.GetCollection("quest")

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetQuest(c *gin.Context, id int) (domain.QuestDTO, error) {
	return domain.QuestDTO{}, nil
}

func (r *repository) CreateQuest(c *gin.Context, name string) (domain.QuestDTO, error) {

	var err error

	_, err = collection.InsertOne(c, domain.Quest{Name: name})

	if err != nil {
		return domain.QuestDTO{}, err
	}

	return domain.QuestDTO{Name: name}, nil
}
