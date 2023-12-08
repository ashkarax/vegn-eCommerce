package requestmodels

type RazorWebOut struct {
	RPaymentId string `json:"razorpay_payment_id"`
	ROrderId   string `json:"razorpay_order_id"`
	RSignature string `json:"razorpay_signature"`

	RSecrect string
	OrderId  string
	UserId   string
}
