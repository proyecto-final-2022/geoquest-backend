package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
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
	QuestID int   `json:"quest_id"`
}

type TeamCompletionRequest struct {
	StartYear    int `json:"start_year"`
	StartMonth   int `json:"start_month"`
	StartDay     int `json:"start_day"`
	StartHour    int `json:"start_hour"`
	StartMinutes int `json:"start_minutes"`
	StartSeconds int `json:"start_seconds"`
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

		id, err := t.service.CreateTeam(c, req.UserIDs, req.QuestID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, id)
	}
}

// @Summary Add team completion
// @Schemes
// @Description Save new team completion
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path int true "Team ID"
// @Param quest_id path int true "Quest ID"
// @Param user body TeamCompletionRequest true "Add completion to team"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /teams/{id}/completions/{quest_id} [post]
func (t *Team) AddCompletion() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req TeamCompletionRequest

		paramId, _ := strconv.Atoi(c.Param("id"))
		paramQuestId, _ := strconv.Atoi(c.Param("quest_id"))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		if err := t.service.AddCompletion(c, paramId, paramQuestId, req.StartYear, time.Month(req.StartMonth), req.StartDay, req.StartHour, req.StartMinutes, req.StartSeconds); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Ranking of quest
// @Schemes
// @Description Get Ranking of teams by a specific quest
// @Tags Teams
// @Accept json
// @Produce json
// @Param quest_id path int true "Quest ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /teams/rankings/{quest_id} [get]
func (t *Team) GetRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ranking []domain.QuestTeamCompletionDTO

		questId, _ := strconv.Atoi(c.Param("quest_id"))

		ranking, err := t.service.GetRanking(c, questId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, ranking)
	}
}

// @Summary WaitRoom of a quest
// @Schemes
// @Description Get Waitroom to see users that accepted the invitation
// @Tags Teams
// @Accept json
// @Produce json
// @Param team_id path int true "Team ID"
// @Param quest_id path int true "Quest ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /teams/waitrooms/{team_id}/quests/{quest_id} [get]
func (t *Team) GetWaitRoomAccepted() gin.HandlerFunc {
	return func(c *gin.Context) {

		teamId, _ := strconv.Atoi(c.Param("team_id"))
		questId, _ := strconv.Atoi(c.Param("quest_id"))

		waitRoom, err := t.service.GetWaitRoomAccepted(c, teamId, questId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, waitRoom)
	}
}

// @Summary Delete team
// @Schemes
// @Description Delete a team
// @Tags Teams
// @Accept json
// @Produce json
// @Param id path string true "Team ID"
// @Success 200
// @Failure 500
// @Router /teams/{id} [delete]
func (t *Team) DeleteTeam() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		paramId, _ := strconv.Atoi(c.Param("id"))

		if err = t.service.DeleteTeam(c, paramId); err != nil {
			c.JSON(http.StatusInternalServerError, paramId)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}

// @Summary Accept invitation to Quest
// @Schemes
// @Description Accept invitation from a team
// @Tags Teams
// @Accept json
// @Produce json
// @Param team_id path string true "Team ID"
// @Param user_id path string true "User ID"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /teams/waitrooms/{team_id}/users/{user_id} [put]
func (t *Team) AcceptQuestTeam() gin.HandlerFunc {
	return func(c *gin.Context) {

		paramTeamId, _ := strconv.Atoi(c.Param("team_id"))
		paramUserId, _ := strconv.Atoi(c.Param("user_id"))

		if err := t.service.AcceptQuestTeam(c, paramTeamId, paramUserId); err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
