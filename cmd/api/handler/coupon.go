package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-final-2022/geoquest-backend/internal/coupon"
)

type Coupon struct {
	service coupon.Service
}

type ClientCouponRequest struct {
	Points    float64 `json:"points"`
	StartTime int64   `json:"start_time"`
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
		fmt.Println(tm)

		actualTime := time.Now()

		diff := actualTime.Sub(tm)
		fmt.Println(diff.String())

		c.JSON(http.StatusOK, coupon)
	}
}
