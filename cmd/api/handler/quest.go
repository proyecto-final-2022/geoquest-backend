package handler

import (
	"net/http"
	"strconv"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
)

type questRequest struct {
	ID   int    `json:"quest_id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Quest struct {
	service quest.Service
}

func NewGame(s quest.Service) *Quest {
	return &Quest{service: s}
}

// @Summary New quest
// @Schemes
// @Description Save new quest
// @Tags Quests
// @Accept json
// @Produce json
// @Param quest body questRequest true "Quest to save"
// @Success 200 {object} domain.Quest
// @Failure 422
// @Failure 500
// @Router /quests/ [post]
func (u *Quest) CreateQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req questRequest
		var user domain.Quest
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
		}

		if user, err = u.service.Post(c, req.ID, req.Name); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, user)
	}
}

// @Summary Quest
// @Schemes
// @Description Quest info
// @Tags Quests
// @Accept json
// @Produce json
// @Param id path int true "Quest ID"
// @Success 200 {object} domain.Quest
// @Failure 500
// @Router /quests/{id} [get]
func (g *Quest) GetQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var game domain.Quest
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if game, err = g.service.Get(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, game)
	}
}