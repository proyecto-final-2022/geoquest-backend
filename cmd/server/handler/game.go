package handler

import (
	"geoquest-backend/internal/domain"
	"geoquest-backend/internal/game"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Game struct {
	service game.Service
}

func NewGame(s game.Service) *Game {
	return &Game{service: s}
}

func (g *Game) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var game domain.Game
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if game, err = g.service.Get(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, game)
	}
}
