package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

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

		if err = u.service.CreateQuest(c, req.Name); err != nil {
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
// @Param quest body domain.QuestDTO true "Quest to update"
// @Param id path string true "Quest ID"
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

		if err = g.service.UpdateQuest(c, paramId, req); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
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

// @Summary Ranking
// @Schemes
// @Description Quest ranking
// @Tags Quests
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth token"
// @Param id path string true "Quest ID"
// @Success 200
// @Failure 500
// @Router /quests/{id}/rankings [get]
func (g *Quest) GetRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var quests []domain.QuestCompletion

		paramId, _ := strconv.Atoi(c.Param("id"))

		if quests, err = g.service.GetRanking(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, quests)
	}
}
