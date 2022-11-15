package coupon

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Service interface {
	GenerateCoupons(c *gin.Context, clientID int, userID int, points float64) (domain.CouponDTO, error)
}

type service struct {
	repo Repository
}

func NewService(rep Repository) Service {
	return &service{repo: rep}
}

func (s *service) GenerateCoupons(c *gin.Context, clientID int, userID int, points float64) (domain.CouponDTO, error) {

	var performance string
	var err error

	if points > 400 {
		performance = "ok"
	}

	if points > 800 {
		performance = "great"
	}

	if points > 10000 {
		performance = "superb"
	}

	couponsClients, err := s.repo.GetCoupons(c, clientID, performance)

	tm := time.Date(2023, 1, 5, 22, 00, 00, 00, time.UTC)

	//se puede hacer un FOR y agarrar lista de cupones

	couponID, err := s.repo.CreateCoupon(c, userID, couponsClients[0].Description, couponsClients[0].ClientID, tm)
	if err != nil {
		return domain.CouponDTO{}, err
	}

	return domain.CouponDTO{ID: couponID, Description: couponsClients[0].Description, UserID: userID, ClientID: couponsClients[0].ClientID}, err
}
