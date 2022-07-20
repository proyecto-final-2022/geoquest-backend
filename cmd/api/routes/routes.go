package routes

import (
	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"
	"github.com/proyecto-final-2022/geoquest-backend/internal/user"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r *gin.Engine
}

func NewRouter(r *gin.Engine) Router {
	return &router{r: r}
}

func (r *router) MapRoutes() {
	r.buildGamesRoutes()
	r.buildUsersRoutes()

}

func (r *router) buildGamesRoutes() {
	repo := quest.NewRepository()
	service := quest.NewService(repo)
	handler := handler.NewGame(service)

	gGroup := r.r.Group("/quests")
	{
		gGroup.POST("/", handler.CreateQuest())
		gGroup.GET("/", handler.GetQuests())
	}
}

func (r *router) buildUsersRoutes() {
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := handler.NewUser(service)

	gGroup := r.r.Group("/users")
	{
		gGroup.POST("/", handler.CreateUser())
		gGroup.GET("/:id", handler.GetUser())
	}
}
