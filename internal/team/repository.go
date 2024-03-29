package team

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Repository interface {
	CreateTeam(c *gin.Context, questID int) (int, error)
	AddPlayer(c *gin.Context, teamID int, playerID int, questID int, accepted bool) error
	AddCompletion(c *gin.Context, teamID int, questId int, startTime time.Time, endTime time.Time) error
	GetRanking(c *gin.Context, questId int) ([]domain.QuestTeamCompletion, error)
	GetTeam(c *gin.Context, teamID int) ([]domain.UserXTeam, error)
	GetTeams(c *gin.Context, questID int) ([]domain.Team, error)
	DeleteTeam(c *gin.Context, teamId int) error
	DeletePlayerFromTeam(c *gin.Context, teamId int, userId int) error
	AcceptQuestTeam(c *gin.Context, teamId int, userId int) error
	GetWaitRoomAccepted(c *gin.Context, teamId int, questId int) ([]domain.UserXTeam, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateTeam(c *gin.Context, questID int) (int, error) {
	team := domain.Team{QuestID: questID}
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

func (r *repository) GetTeams(c *gin.Context, questID int) ([]domain.Team, error) {
	var team []domain.Team
	if tx := config.MySql.Where("quest_id = ?", questID).Find(&team); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return team, nil
}

func (r *repository) GetWaitRoomAccepted(c *gin.Context, teamID int, questID int) ([]domain.UserXTeam, error) {
	var team []domain.UserXTeam
	if tx := config.MySql.Where("team_id = ? AND quest_id = ? AND Accept = 1", teamID, questID).Find(&team); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return team, nil
}

func (r *repository) AddPlayer(c *gin.Context, teamID int, playerID int, questID int, accepted bool) error {
	addPlayer := domain.UserXTeam{TeamID: teamID, UserID: playerID, QuestID: questID, Accept: accepted}
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

func (r *repository) DeletePlayerFromTeam(c *gin.Context, teamId int, userId int) error {
	if tx := config.MySql.Where("team_id = ? AND user_id = ?", teamId, userId).Delete(&[]domain.UserXTeam{}); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}

func (r *repository) AcceptQuestTeam(c *gin.Context, teamId int, userId int) error {

	if tx := config.MySql.Where("team_id = ? AND user_id = ? AND Accept = 0", teamId, userId).First(&domain.UserXTeam{}).Update("Accept", true); tx.Error != nil {
		return tx.Error
	}

	return nil
}
