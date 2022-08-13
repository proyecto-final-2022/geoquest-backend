package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/team"
)

type Team struct {
	service team.Service
}

func NewTeam(s team.Service) *Team {
	return &Team{service: s}
}

type TeamRequest struct {
	UserIDs []int `json:"user_ids"`
}

// @Summary New Team
// @Schemes
// @Description Save new team
// @Tags Teams
// @Accept json
// @Produce json
// @Param user body TeamRequest true "Team to save"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /teams/ [post]
func (t *Team) CreateTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TeamRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		id, err := t.service.CreateTeam(c, req.UserIDs)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
