package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/coupon"
	"github.com/proyecto-final-2022/geoquest-backend/internal/domain"
)

type Coupon struct {
	service coupon.Service
}

type ClientCouponRequest struct {
	Points    float64 `json:"points"`
	StartTime int64   `json:"start_time"`
}

type ClientCouponResponse struct {
	Coupon        domain.CouponDTO `json:"coupon"`
	QuestDuration string           `json:"quest_duration"`
}

func NewCoupon(s coupon.Service) *Coupon {
	return &Coupon{service: s}
}

// @Summary Coupon
// @Schemes
// @Description
// @Tags Coupons
// @Accept json
// @Produce json
// @Param client_id path string true "Client ID"
// @Param user_id path string true "User ID"
// @Param quest body ClientCouponRequest true "Coupon characteristics"
// @Success 200
// @Failure 422
// @Failure 500
// @Router /coupons/{client_id}/completions/{user_id} [post]
func (co *Coupon) CompletionCoupons() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		var req ClientCouponRequest
		var res ClientCouponResponse

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		paramClientId, _ := strconv.Atoi(c.Param("client_id"))
		paramUserId, _ := strconv.Atoi(c.Param("user_id"))

		coupon, err := co.service.GenerateCoupons(c, paramClientId, paramUserId, req.Points)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		startTime := strconv.FormatInt(req.StartTime, 10)

		i, err := strconv.ParseInt(startTime, 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(i, 0)
		actualTime := time.Now()
		diff := actualTime.Sub(tm)
		//		diff.parse

		res.QuestDuration = diff.String()
		//fmt.Sprintf("%d", int(diff.Hours())) + " H " + fmt.Sprintf("%d", int(diff.Minutes())) + " m " + fmt.Sprintf("%d", int(diff.Seconds())) + " s "
		res.Coupon = coupon

		c.JSON(http.StatusOK, res)
	}
}
