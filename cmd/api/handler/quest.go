package handler

import (
	"net/http"
	"strconv"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
)

type questRequest struct {
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
// @Success 200 {object} domain.QuestDTO
// @Failure 422
// @Failure 500
// @Router /quests/ [post]
func (u *Quest) CreateQuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req questRequest
		var err error

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
		}

		if err = u.service.CreateQuest(c, req.Name); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, nil)
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

		paramId, _ := strconv.Atoi(c.Param("id"))

		if quests, err = g.service.GetQuests(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, quests)
	}
}
