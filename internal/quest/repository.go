package quest

import (
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Get(c *gin.Context, id int) (domain.Quest, error)
	Post(c *gin.Context, id int, name string) (domain.Quest, error)
}

type repository struct {
}

var collection = config.GetCollection("quest")

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Get(c *gin.Context, id int) (domain.Quest, error) {
	return domain.Quest{ID: id}, nil
}

func (r *repository) Post(c *gin.Context, id int, name string) (domain.Quest, error) {

	var err error

	_, err = collection.InsertOne(c, domain.Quest{
		ID:   id,
		Name: name})

	if err != nil {
		return domain.Quest{}, err
	}

	return domain.Quest{ID: id, Name: name}, nil
}
