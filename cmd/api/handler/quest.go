package handler

import (
	"net/http"
	"strconv"

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
func (g *Quest) Get() gin.HandlerFunc {
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
