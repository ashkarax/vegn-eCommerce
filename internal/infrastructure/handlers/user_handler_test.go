package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces/mocksUseCase"
	requestmodels "github.com/ashkarax/vegn-eCommerce/internal/models/request_models"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUserSignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUseCase := mocksUseCase.NewMockIuserUseCase(mockCtrl)
	userHandler := NewUserhandler(mockUseCase)
	defer mockCtrl.Finish()

	tests := []struct {
		name          string
		input         requestmodels.UserSignUpReq
		stub          func(mocksUseCase.MockIuserUseCase, requestmodels.UserSignUpReq)
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "signup success",
			input: requestmodels.UserSignUpReq{
				FirstName:       "Ashkar",
				LastName:        "A.S",
				Email:           "ashkar@example.com",
				Phone:           "+91000000000",
				Password:        "password",
				ConfirmPassword: "password",
			},
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.UserSignUpReq) {
				miuc.EXPECT().UserSignUp(&usur).Times(1).Return(responsemodels.SignupData{
					Token: "temporarytokenforverification",
				}, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "signup failed",
			input: requestmodels.UserSignUpReq{
				FirstName:       "Ashkar",
				LastName:        "A.S",
				Email:           "ashkar@example.com",
				Phone:           "+91000000000",
				Password:        "password",
				ConfirmPassword: "newpassword",
			},
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.UserSignUpReq) {
				miuc.EXPECT().UserSignUp(&usur).Times(1).Return(responsemodels.SignupData{}, errors.New("invalid credentials"))
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(*mockUseCase, tc.input)

			server := gin.Default()
			server.POST("/signup", userHandler.UserSignUp)

			jsonData, _ := json.Marshal(tc.input)
			body := bytes.NewBuffer(jsonData)

			mockRequest, _ := http.NewRequest(http.MethodPost, "/signup", body)
			responseRecorder := httptest.NewRecorder()
			server.ServeHTTP(responseRecorder, mockRequest)

			tc.checkResponse(t, responseRecorder)

		})
	}
}

func TestUserOTPVerication(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUseCase := mocksUseCase.NewMockIuserUseCase(mockCtrl)
	mockHandler := NewUserhandler(mockUseCase)
	defer mockCtrl.Finish()

	tests := []struct {
		name          string
		input1        requestmodels.OtpVerification
		input2        string
		stub          func(mocksUseCase.MockIuserUseCase, requestmodels.OtpVerification, string)
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{

		{
			name: "success verification",
			input1: requestmodels.OtpVerification{
				Otp: "0000",
			},
			input2: "TempTokenForOtpVerification",
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.OtpVerification, token string) {
				miuc.EXPECT().VerifyOtp(&usur, token).Times(1).Return(responsemodels.OtpVerifResult{
					RefreshToken: "RefreshToken",
					AccessToken:  "AccessToken",
				}, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "verification failed",
			input1: requestmodels.OtpVerification{
				Otp: "0000",
			},
			input2: "TempTokenForOtpVerification",
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.OtpVerification, token string) {
				miuc.EXPECT().VerifyOtp(&usur, token).Times(1).Return(responsemodels.OtpVerifResult{}, errors.New("invalid otp"))
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(*mockUseCase, tc.input1, tc.input2)

			server := gin.Default()
			server.POST("/verify", mockHandler.UserOTPVerication)

			jsonData, _ := json.Marshal(tc.input1)
			body := bytes.NewBuffer(jsonData)

			mockRequest, _ := http.NewRequest(http.MethodPost, "/verify", body)
			mockRequest.Header.Set("Authorizations", tc.input2)

			responseRecorder := httptest.NewRecorder()
			server.ServeHTTP(responseRecorder, mockRequest)

			tc.checkResponse(t, responseRecorder)

		})

	}

}

func TestUserLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockUseCase := mocksUseCase.NewMockIuserUseCase(mockCtrl)
	mockHandler := NewUserhandler(mockUseCase)
	defer mockCtrl.Finish()

	tests := []struct {
		name          string
		input         requestmodels.UserLoginReq
		stub          func(mocksUseCase.MockIuserUseCase, requestmodels.UserLoginReq)
		checkResponse func(*testing.T, *httptest.ResponseRecorder)
	}{

		{
			name: "successful login attempt",
			input: requestmodels.UserLoginReq{
				Phone:    "+9100000000",
				Password: "password",
			},
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.UserLoginReq) {
				miuc.EXPECT().UserLogin(&usur).Times(1).Return(responsemodels.UserLoginRes{
					RefreshToken: "RefreshToken",
					AccessToken:  "AccessToken",
				}, nil)
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, rr.Code)
			},
		},
		{
			name: "failed login attempt",
			input: requestmodels.UserLoginReq{
				Phone:    "+9100000000",
				Password: "InvalidPassword",
			},
			stub: func(miuc mocksUseCase.MockIuserUseCase, usur requestmodels.UserLoginReq) {
				miuc.EXPECT().UserLogin(&usur).Times(1).Return(responsemodels.UserLoginRes{}, errors.New("wrong credentials"))
			},
			checkResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, rr.Code)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(*mockUseCase, tc.input)

			server := gin.Default()
			server.POST("/login", mockHandler.UserLogin)

			jsonData, _ := json.Marshal(tc.input)
			body := bytes.NewBuffer(jsonData)

			mockRequest, _ := http.NewRequest(http.MethodPost, "/login", body)

			responseRecorder := httptest.NewRecorder()
			server.ServeHTTP(responseRecorder, mockRequest)

			tc.checkResponse(t, responseRecorder)

		})

	}

}
