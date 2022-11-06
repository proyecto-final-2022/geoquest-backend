package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"
	"gorm.io/datatypes"

	"github.com/gin-gonic/gin"
)

type Quest struct {
	service quest.Service
}

type CompletionRequest struct {
	StartYear    int `json:"start_year"`
	StartMonth   int `json:"start_month"`
	StartDay     int `json:"start_day"`
	StartHour    int `json:"start_hour"`
	StartMinutes int `json:"start_minutes"`
	StartSeconds int `json:"start_seconds"`
}

type QuestRequest struct {
	ID        string   `json:"id"`
	Scene     int      `json:"scene"`
	Inventory []string `json:"inventory"`
}

type QuestProgressRequest struct {
	Scene     int            `json:"scene"`
	Logs      []string       `json:"logs"`
	Inventory []string       `json:"inventory"`
	Points    float64        `json:"points"`
	Objects   map[string]int `json:"objects"`
}

type Rating struct {
	Rating int `json:"rating"`
}

type WaitRoomRequest struct {
	UserIDS []int `json:"user_ids"`
}

/*
t1 := time.Now()
t2 := t1.Add(time.Second * 341)

fmt.Println(t1)
fmt.Println(t2)

diff := t2.Sub(t1)
fmt.Println(diff)
*/

func NewGame(s quest.Service) *Quest {
	return &Quest{service: s}
}

// @Summary New quest
// @Schemes
// @Description Save new quest
// @Tags Quests
// @Accept json
// @Produce json
// @Param quest body domain.QuestDTO true "Quest to save"
// @Param Authorization header string true "Auth token"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /quests/ [post]
func (u *Quest) CreateQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.QuestDTO
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.CreateQuest(c, req.QuestID, req.Scene, req.Inventory, req.Logs, req.Points); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary New quest progression
// @Schemes
// @Description Save new quest progression
// @Tags Quests
// @Accept json
// @Produce json
// @Param quest body QuestProgressRequest true "Quest progress to save"
// @Param id path string true "Quest ID"
// @Param team_id path string true "Team ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /quests/{id}/progression/{team_id} [post]
func (u *Quest) CreateQuestProgression() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req QuestProgressRequest
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramTeamId, _ := strconv.Atoi(c.Param("team_id"))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.CreateQuestProgression(c, paramId, paramTeamId, req.Scene, req.Inventory, req.Logs, req.Objects, req.Points); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Get Team progression
// @Schemes
// @Description Team progression
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path string true "Quest ID"
// @Param team_id path string true "Team ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/progression/{team_id} [get]
func (g *Quest) GetQuestProgression() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var questProgress datatypes.JSON

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramTeamId, _ := strconv.Atoi(c.Param("team_id"))

		if questProgress, err = g.service.GetQuestProgression(c, paramId, paramTeamId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, questProgress)
	}
}

// @Summary Update quest progression
// @Schemes
// @Description Update quest progression
// @Tags Quests
// @Accept json
// @Produce json
// @Param quest body QuestProgressRequest true "Quest progress to update"
// @Param id path string true "Quest ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /quests/{id}/progression [put]
func (u *Quest) UpdateQuestProgression() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req QuestProgressRequest
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = u.service.UpdateQuestProgression(c, paramId, req.Scene, req.Inventory, req.Logs, req.Objects, req.Points); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Quests
// @Schemes
// @Description Quest info
// @Tags Quests
// @Accept json
// @Produce json
// @Success 200 {object} domain.QuestDTO
// @Failure 500
// @Router /quests/ [get]
func (g *Quest) GetQuests() gin.HandlerFunc {
	return func(c *gin.Context) {
		var quests []*domain.QuestDTO
		var err error

		if quests, err = g.service.GetQuests(c); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}

// @Summary Quests
// @Schemes
// @Description Quest info
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path string true "Quest ID"
// @Param Authorization header string true "Auth token"
// @Success 200 {object} domain.QuestDTO
// @Failure 500
// @Router /quests/{id} [get]
func (g *Quest) GetQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var quests domain.QuestDTO
		var err error

		paramId := c.Param("id")

		if quests, err = g.service.GetQuest(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}

// @Summary Quests
// @Schemes
// @Description Quest info
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path string true "Quest ID"
// @Param quest body domain.QuestDTO true "Quest to update"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /quests/{id} [put]
func (g *Quest) UpdateQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.QuestDTO
		var err error

		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramId := c.Param("id")

		if err = g.service.UpdateQuest(c, req, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Quests
// @Schemes
// @Description Delete a quest
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path string true "Quest ID"
// @Success 200
// @Failure 500
// @Router /quests/{id} [delete]
func (g *Quest) DeleteQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId := c.Param("id")

		if err = g.service.DeleteQuest(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Completion
// @Schemes
// @Description Completion of a quest
// @Tags Quests
// @Accept json
// @Produce json
// @Param completion body CompletionRequest true "Quest completed by a User"
// @Param Authorization header string true "Auth token"
// @Param id path string true "Quest ID"
// @Param user_id path string true "User ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/completions/{user_id} [post]
func (g *Quest) AddCompletion() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var req CompletionRequest

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramUserId, _ := strconv.Atoi(c.Param("user_id"))

		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = g.service.CreateCompletion(c, paramId, paramUserId, req.StartYear, time.Month(req.StartMonth), req.StartDay, req.StartHour, req.StartMinutes, req.StartSeconds); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Completion
// @Schemes
// @Description Get a User Rating
// @Tags Quests
// @Accept json
// @Produce json
// @Param completion body CompletionRequest true "Quest completed by a User"
// @Param Authorization header string true "Auth token"
// @Param id path string true "Quest ID"
// @Param user_id path string true "User ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/rating/{user_id} [post]
func (g *Quest) GetRating() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var rating domain.Rating

		paramQuestId, _ := strconv.Atoi(c.Param("id"))
		paramUserId, _ := strconv.Atoi(c.Param("user_id"))

		if rating, err = g.service.GetRating(c, paramQuestId, paramUserId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, rating)
	}
}

// @Summary Completion
// @Schemes
// @Description Rate a quest
// @Tags Quests
// @Accept json
// @Produce json
// @Param completion body CompletionRequest true "Quest completed by a User"
// @Param Authorization header string true "Auth token"
// @Param id path string true "Quest ID"
// @Param user_id path string true "User ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/rating/{user_id} [post]
func (g *Quest) AddRating() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var req Rating

		paramQuestId, _ := strconv.Atoi(c.Param("id"))
		paramUserId, _ := strconv.Atoi(c.Param("user_id"))

		if err = c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err = g.service.CreateRating(c, paramQuestId, paramUserId, req.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Ranking
// @Schemes
// @Description Quest ranking
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path string true "Quest ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/rankings [get]
func (g *Quest) GetRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var quests []domain.QuestCompletionDTO

		paramId, _ := strconv.Atoi(c.Param("id"))

		if quests, err = g.service.GetRanking(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}
