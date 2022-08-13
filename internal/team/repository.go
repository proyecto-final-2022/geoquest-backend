package team

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Repository interface {
	CreateTeam(c *gin.Context) (int, error)
	AddPlayer(c *gin.Context, teamID, playerID int) error
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

func (r *repository) AddPlayer(c *gin.Context, teamID int, playerID int) error {
	addPlayer := domain.UserXTeam{TeamID: teamID, UserID: playerID}
	if tx := config.MySql.Create(&addPlayer); tx.Error != nil {
		return errors.New("DB Error")
	}
	return nil
}
