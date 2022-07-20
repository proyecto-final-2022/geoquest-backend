package handler

import (
	"net/http"

	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
	"github.com/proyecto-final-2022/geoquest-backend/internal/quest"

	"github.com/gin-gonic/gin"
)

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
// @Param quest body domain.QuestDTO true "Quest to save"
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
