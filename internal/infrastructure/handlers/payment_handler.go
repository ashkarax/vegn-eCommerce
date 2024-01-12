package handlers

import (
	"fmt"
	"net/http"

	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	PaymentUseCase interfaceUseCase.IPaymentUseCase
}

func NewPaymentHandler(PaymentUseCase interfaceUseCase.IPaymentUseCase) *PaymentHandler {
	return &PaymentHandler{PaymentUseCase: PaymentUseCase}
}

func (u *PaymentHandler) OnlinePayment(c *gin.Context) {
	orderId := c.Query("orderid")
	// userId, _ := c.Get("userId")
	// userIdString, _ := userId.(string)
	userId := c.Query("userid")

	orderDetails, err := u.PaymentUseCase.GetOrderDetails(&userId, &orderId)
	fmt.Println("-----------", userId)
	fmt.Println("--------------", orderId)
	if err != nil {

		c.HTML(http.StatusBadRequest, "razorpay.html", gin.H{"badRequest": "Refine your request"})
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", orderDetails)

}

// @Summary VerifyPayment
// @Description Verifies an online payment for a specific order.
// @Tags Payments
// @Accept json
// @Produce json
// @Security UserAuthTokenAuth
// @Security UserRefTokenAuth
// @Param orderid path string true "The ID of the order to verify payment for."
// @Param razorResp body requestmodels.RazorWebOut true "Razorpay payment response details."
// @Success 200  {object} responsemodels.Response
// @Failure 400  {object} responsemodels.Response
// @Router /payments/verify/{orderid} [post]
func (u *PaymentHandler) VerifyPayment(c *gin.Context) {
	orderId := c.Param("orderid")
	userId, _ := c.Get("userId")
	userIdString, _ := userId.(string)

	var razorResp requestmodels.RazorWebOut

	

	if err := c.BindJSON(&razorResp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	razorResp.OrderId = orderId
	razorResp.UserId = userIdString
	
	order, err := u.PaymentUseCase.OnlinePaymentVerification(&razorResp)
	if err != nil {
		finalReslt := responsemodels.Responses(http.StatusBadRequest, "", nil, err.Error())
		c.JSON(http.StatusBadRequest, finalReslt)
		return
	}

	finalReslt := responsemodels.Responses(http.StatusOK, "", order, nil)
	c.JSON(http.StatusOK, finalReslt)

}
