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
				mockUseCase.EXPECT().UserSignUp(&usur).Times(1).Return(responsemodels.SignupData{
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
				mockUseCase.EXPECT().UserSignUp(&usur).Times(1).Return(responsemodels.SignupData{}, errors.New("invalid credentials"))
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
