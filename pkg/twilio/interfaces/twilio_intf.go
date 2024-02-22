package interfaceTwilio

import "github.com/twilio/twilio-go"

type ITwilio interface {
	TwilioClient() *twilio.RestClient
	SendOTP(phone string, client *twilio.RestClient) (string, error)
	VerifyOtp(phone string, otp string, client *twilio.RestClient) error
}
