package twilio

import (
	"errors"
	"strings"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type twilioService struct {
	twilioCredentials config.OTP
}

var twilioservice twilioService

func OTPServiceSetup(data config.OTP) {
	twilioservice.twilioCredentials = data
}


func TwilioClient() *twilio.RestClient {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioservice.twilioCredentials.AccountSid,
		Password: twilioservice.twilioCredentials.AuthToken,
	})

	return client

}

func SendOTP(phone string,client *twilio.RestClient) (string, error) {

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(twilioservice.twilioCredentials.ServiceSid, params)
	if err != nil {
		if strings.Contains(err.Error(), "ApiError 60200") {
			err = errors.New("please enter country code in the phone number")
		}
		return "", err
	}

	return *resp.Status, nil

}

func VerifyOtp (phone string,otp string,client *twilio.RestClient) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(twilioservice.twilioCredentials.ServiceSid, params)
	if err != nil {
		if strings.Contains(err.Error(), "ApiError 60200") {
			err = errors.New("please add country code in the phone number")
		}
		return err
	}

	if *resp.Status != "approved" {
		err := errors.New(*resp.Status)
		return err
	}

	return nil
}
