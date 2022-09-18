package team

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Repository interface {
	CreateTeam(c *gin.Context) (int, error)
	AddPlayer(c *gin.Context, teamID, playerID int, questID int) error
	AddCompletion(c *gin.Context, teamID int, questId int, startTime time.Time, endTime time.Time) error
	GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletion, error)
	GetTeam(c *gin.Context, teamID int) ([]domain.UserXTeam, error)
	DeleteTeam(c *gin.Context, teamId int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateTeam(c *gin.Context) (int, error) {
	team := domain.Team{}
	if tx := config.MySql.Create(&team); tx.Error != nil {
		return 0, errors.New("DB Error")
	}

	return team.ID, nil
}

func (r *repository) GetTeam(c *gin.Context, teamID int) ([]domain.UserXTeam, error) {
	var team []domain.UserXTeam
	if tx := config.MySql.Where("team_id = ?", teamID).Find(&team); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return team, nil
}

func (r *repository) AddPlayer(c *gin.Context, teamID int, playerID int, questID int) error {
	addPlayer := domain.UserXTeam{TeamID: teamID, UserID: playerID, QuestID: questID, Accept: false}
	if tx := config.MySql.Create(&addPlayer); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) AddCompletion(c *gin.Context, teamID int, questId int, startTime time.Time, endTime time.Time) error {
	addCompletion := domain.QuestTeamCompletion{TeamID: teamID, QuestID: questId, StartTime: startTime, EndTime: endTime}
	if tx := config.MySql.Create(&addCompletion); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletion, error) {
	var completions []domain.QuestTeamCompletion
	if tx := config.MySql.Where("quest_id = ?", questId).Find(&completions); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return completions, nil
}

func (r *repository) DeleteTeam(c *gin.Context, teamId int) error {
	if tx := config.MySql.Where("id = ?", teamId).Delete(&domain.Team{}); tx.Error != nil {
		return errors.New("DB Error")
	}
	if tx := config.MySql.Where("team_id = ?", teamId).Delete(&[]domain.UserXTeam{}); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}
