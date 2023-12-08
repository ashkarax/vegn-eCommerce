package razorpay

import (
	"fmt"
	"log"

	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	uuidgenerator "github.com/ashkarax/vegn-eCommerce/pkg/uuid_generator"
	"github.com/razorpay/razorpay-go"
	"github.com/razorpay/razorpay-go/utils"
)

func RazorPayInitialize(totalPrice *float64, idKey *string, securityKey *string) (*string, error) {
	var razorreturnId string

	razorpayClient := razorpay.NewClient(*idKey, *securityKey)

	orderParams := map[string]interface{}{
		"amount":   *totalPrice * 100, // Razorpay amounts are in paise
		"currency": "INR",
		"receipt":  uuidgenerator.ReturnUuid(),
	}
	order, err := razorpayClient.Order.Create(orderParams, nil)
	if err != nil {
		log.Println("Error creating order:", err)
		return &razorreturnId, err
	}

	razorreturnId = order["id"].(string)

	return &razorreturnId, nil
}

func VerifyPayment(razorResp *requestmodels.RazorWebOut) bool {

	params := map[string]interface{}{
		"razorpay_order_id":   razorResp.ROrderId,
		"razorpay_payment_id": razorResp.RPaymentId,
	}

	result := utils.VerifyPaymentSignature(params, razorResp.RSignature, razorResp.RSecrect)
	fmt.Println("*****", result)
	return result
}
