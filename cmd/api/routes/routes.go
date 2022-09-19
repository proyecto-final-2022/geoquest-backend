package routes

import (
	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/handler"
	//	"github.com/proyecto-final-2022/geoquest-backend/cmd/api/middlewares"
	"github.com/proyecto-final-2022/geoquest-backend/internal/client"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"
	"github.com/proyecto-final-2022/geoquest-backend/internal/team"
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
	r.buildTeamRoutes()
	r.buildClientsRoutes()
}

func (r *router) buildGamesRoutes() {
	repo := quest.NewRepository()
	repoUser := user.NewRepository()
	service := quest.NewService(repo, repoUser)
	handler := handler.NewGame(service)

	gGroup := r.r.Group("/quests")
	{
		gGroup.POST("/", handler.CreateQuest())
		gGroup.POST("/:id/completions/:user_id", handler.AddCompletion())
		gGroup.GET("/:id/rankings", handler.GetRanking())
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
		gGroup.POST("/sessions/google", handler.LoginUserGoogle())
		gGroup.POST("/:id/coupons", handler.CreateUserCoupon())
		gGroup.POST("/:id/friends/:friend_id", handler.AddUserFriend())
		gGroup.POST("/:id/notifications/", handler.AddNotification())
		gGroup.DELETE("/:id/notifications/:notification_id", handler.DeleteNotification())
		gGroup.GET("/:id/notifications/", handler.GetNotifications())
		gGroup.GET("/:id/friends", handler.GetUserFriends())
		gGroup.GET("/:id/coupons", handler.GetUserCoupons())
		gGroup.GET("/:id", handler.GetUser())
		gGroup.PUT("/:id", handler.UpdateUser())
		gGroup.DELETE("/:id", handler.DeleteUser())
		gGroup.DELETE("/:id/friends/:friend_id", handler.DeleteUserFriend())
	}
}

func (r *router) buildClientsRoutes() {
	repo := client.NewRepository()
	service := client.NewService(repo)
	handler := handler.NewClient(service)

	gGroup := r.r.Group("/clients")
	{
		gGroup.POST("/", handler.CreateClient())
		gGroup.GET("/", handler.GetClients())
		gGroup.POST("/:id/quests", handler.CreateClientQuest())
		gGroup.POST("/quests/:id", handler.AddTag())
		gGroup.GET("/quests/", handler.GetAllQuests())
		gGroup.GET("/:id/quests", handler.GetClientQuests())
	}
}

func (r *router) buildTeamRoutes() {
	userRepo := user.NewRepository()
	repo := team.NewRepository()
	questRepo := quest.NewRepository()
	service := team.NewService(repo, userRepo, questRepo)
	handler := handler.NewTeam(service)

	gGroup := r.r.Group("/teams")
	{
		gGroup.POST("/", handler.CreateTeam())
		gGroup.POST("/:id/completions/:quest_id", handler.AddCompletion())
		gGroup.PUT("/waitrooms/:team_id/users/:user_id", handler.AcceptQuestTeam())
		gGroup.GET("/rankings/:quest_id", handler.AcceptQuestTeam())
		gGroup.GET("/waitrooms/:team_id/quests/:quest_id/", handler.GetWaitRoomAccepted())
		gGroup.DELETE("/:id/", handler.DeleteTeam())
	}
}
