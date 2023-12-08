package handlers

import (
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	Usecase interfaceUseCase.ICouponUseCase
}

func NewCouponHandler(usecase interfaceUseCase.ICouponUseCase) *CouponHandler {
	return &CouponHandler{Usecase: usecase}
}

func (u *CouponHandler) AllCoupons(c *gin.Context) {
	couponSlice, err := u.Usecase.GetAllCoupons()
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "failed fetching coupons", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupons fetched succesfully", couponSlice, nil)
	c.JSON(http.StatusOK, response)
}

func (u *CouponHandler) AddNewCoupon(c *gin.Context) {
	var couponData requestmodels.CouponDetailsReq

	if err := c.BindJSON(&couponData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.Usecase.AddNewCoupon(&couponData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't add coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupon added succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}

func (u *CouponHandler) DeleteCoupon(c *gin.Context) {
	couponid := c.Param("couponid")
	newStatus := "deleted"
	err := u.Usecase.ChangeCouponStatus(&couponid, &newStatus)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't delete coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupon deleted succesfully", nil, nil)
	c.JSON(http.StatusOK, response)

}
func (u *CouponHandler) BlockCoupon(c *gin.Context) {
	couponid := c.Param("couponid")
	newStatus := "blocked"
	err := u.Usecase.ChangeCouponStatus(&couponid, &newStatus)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't Block coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupon Blocked succesfully", nil, nil)
	c.JSON(http.StatusOK, response)

}
func (u *CouponHandler) UnBlockCoupon(c *gin.Context) {
	couponid := c.Param("couponid")
	newStatus := "active"
	err := u.Usecase.ChangeCouponStatus(&couponid, &newStatus)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't UnBlock coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupon UnBlocked succesfully", nil, nil)
	c.JSON(http.StatusOK, response)

}

func (u *CouponHandler) EditCoupon(c *gin.Context) {
	couponid := c.Param("couponid")
	var couponData requestmodels.CouponDetailsReq
	couponData.Id = couponid
	
	if err := c.BindJSON(&couponData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.Usecase.EditCoupon(&couponData)
	if err != nil {
		response := responsemodels.Responses(http.StatusBadRequest, "can't edit coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodels.Responses(http.StatusOK, "coupon edited succesfully", nil, nil)
	c.JSON(http.StatusOK, response)
}
