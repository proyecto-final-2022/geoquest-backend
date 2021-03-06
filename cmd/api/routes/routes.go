package routes

import (
	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/middlewares"
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

	gGroup := r.r.Group("/quests").Use(middlewares.Auth())
	{
		gGroup.POST("/", handler.CreateQuest())
		gGroup.GET("/", handler.GetQuests())
		gGroup.GET("/:id", handler.GetQuest())
		gGroup.PUT("/:id", handler.UpdateQuest())
		gGroup.DELETE("/:id", handler.DeleteQuest())
	}
}

func (r *router) buildUsersRoutes() {
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := handler.NewUser(service)

	//por ahora lo dejamos como que con el token no puedas acceder a endpoints relacionados a las quests
	//tienen que ser los endpoints relacionados a las interacciones que hace el usuario
	gGroup := r.r.Group("/users")
	{
		gGroup.POST("/", handler.CreateUser())
		gGroup.POST("/sessions/", handler.LoginUser())
		gGroup.GET("/:id", handler.GetUser())
		gGroup.PUT("/:id", handler.UpdateUser())
		gGroup.DELETE("/:id", handler.DeleteUser())
	}
}
