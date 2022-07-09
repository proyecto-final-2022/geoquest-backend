package routes

import (
	"geoquest-backend/cmd/server/handler"
	"geoquest-backend/internal/game"

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
}

func (r *router) buildGamesRoutes() {
	repo := game.NewRepository()
	service := game.NewService(repo)
	handler := handler.NewGame(service)

	gGroup := r.r.Group("/games")
	{
		gGroup.GET("/:id", handler.Get())
	}

}
