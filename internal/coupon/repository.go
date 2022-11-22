package coupon

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/config"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Repository interface {
	GetCoupons(c *gin.Context, clientID int, performance string) ([]domain.CouponClient, error)
	GetTeam(c *gin.Context, teamID int) ([]domain.UserXTeam, error)
	CreateCoupon(c *gin.Context, userID int, description string, clientID int, date time.Time) (int, error)
	UpdateQuestProgressionFinish(c *gin.Context, questId int, teamId int) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetCoupons(c *gin.Context, clientID int, performance string) ([]domain.CouponClient, error) {
	var coupons []domain.CouponClient
	if tx := config.MySql.Where("client_id = ? AND quest_performance = ?", clientID, performance).Find(&coupons); tx.Error != nil {
		return nil, errors.New("DB Error")
	}
	return coupons, nil
}

func (r *repository) GetTeam(c *gin.Context, teamID int) ([]domain.UserXTeam, error) {
	var team []domain.UserXTeam
	if tx := config.MySql.Where("team_id = ?", teamID).Find(&team); tx.Error != nil {
		return nil, errors.New("DB Error")
	}

	return team, nil
}

func (r *repository) CreateCoupon(c *gin.Context, userID int, description string, clientID int, date time.Time) (int, error) {

	coupon := domain.Coupon{UserID: userID, ClientID: clientID, Description: description, ExpirationDate: date}

	if tx := config.MySql.Create(&coupon); tx.Error != nil {
		return 0, errors.New("DB Error")
	}
	return coupon.ID, nil
}

func (r *repository) UpdateQuestProgressionFinish(c *gin.Context, questId int, teamId int) error {

	var questProgress domain.QuestProgress

	if tx := config.MySql.Model(&questProgress).Where("quest_id = ? AND team_id = ?", questId, teamId).Update("finished", true); tx.Error != nil {
		return tx.Error
	}

	return nil
}
