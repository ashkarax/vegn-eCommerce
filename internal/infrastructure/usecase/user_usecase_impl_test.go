package usecase

import (
	"errors"
	"testing"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	"github.com/ashkarax/vegn-eCommerce/internal/infrastructure/repository/interfaces/mocks"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

// func TestUserSignUp(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)

// 	mockUserRepo := mocks.NewMockIuserRepo(mockCtrl)
// 	mockTwilio := mocksTwilio.NewMockITwilio(mockCtrl)
// 	mockHashPass := mocksHashPass.NewMockIHashPass(mockCtrl)
// 	mockJwt := mocksJwt.NewMockIJwt(mockCtrl)

// 	userUseCase := NewUserUseCase(mockUserRepo, &config.Token{})
// 	twilioClient := &twilio.RestClient{}

// 	tests := []struct {
// 		name    string
// 		args    requestmodels.UserSignUpReq
// 		stub    func(requestmodels.UserSignUpReq, mocks.MockIuserRepo, mocksTwilio.MockITwilio, mocksHashPass.MockIHashPass, mocksJwt.MockIJwt)
// 		want    responsemodels.SignupData
// 		wantErr error
// 	}{
// 		{
// 			name: "success",
// 			args: requestmodels.UserSignUpReq{
// 				FirstName:       "Ashkar",
// 				LastName:        "A.S",
// 				Email:           "ashkar@example.com",
// 				Phone:           "+918921791915",
// 				Password:        "Password",
// 				ConfirmPassword: "Password",
// 			},
// 			stub: func(usur requestmodels.UserSignUpReq, userRepo mocks.MockIuserRepo, twilio mocksTwilio.MockITwilio, hashPass mocksHashPass.MockIHashPass, jwt mocksJwt.MockIJwt) {
// 				userRepo.EXPECT().IsUserExist(usur.Phone).Times(1).Return(false)

// 				twilio.EXPECT().TwilioClient().Times(1).Return(twilioClient)
// 				twilio.EXPECT().SendOTP(usur.Phone, twilioClient).Times(1).Return("", nil)

// 				hashPass.EXPECT().HashPassword(usur.ConfirmPassword).Times(1).Return("HashedPassword")

// 				userRepo.EXPECT().CreateUser(usur).Times(1)

// 				jwt.EXPECT().TempTokenForOtpVerification("securitykey", usur.Phone).Times(1).Return("temptoken", nil)
// 			},
// 			want: responsemodels.SignupData{
// 				Token: "temptoken",
// 			},
// 			wantErr: nil,
// 		},
// 		{
// 			name: "failure",
// 			args: requestmodels.UserSignUpReq{
// 				FirstName:       "Ashkar",
// 				LastName:        "A.S",
// 				Email:           "ashkar@example.com",
// 				Phone:           "+918921791915",
// 				Password:        "Password",
// 				ConfirmPassword: "Password",
// 			},
// 			stub: func(usur requestmodels.UserSignUpReq, userRepo mocks.MockIuserRepo, twilio mocksTwilio.MockITwilio, hashPass mocksHashPass.MockIHashPass, jwt mocksJwt.MockIJwt) {
// 				userRepo.EXPECT().IsUserExist(usur.Phone).Return(false)

// 				twilio.EXPECT().TwilioClient().Times(1).Return(twilioClient)
// 				twilio.EXPECT().SendOTP(usur.Phone, twilioClient).Times(1).Return("", nil)

// 				hashPass.EXPECT().HashPassword(usur.ConfirmPassword).Times(1).Return("HashedPassword")

// 				userRepo.EXPECT().CreateUser(usur).Times(1)

// 				jwt.EXPECT().TempTokenForOtpVerification("securitykey", usur.Phone).Times(1).Return("", errors.New("error creating token"))
// 			},
// 			want: responsemodels.SignupData{
// 				Token: "",
// 			},
// 			wantErr: errors.New("error creating token"),
// 		},
// 	}
// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {

// 			tc.stub(tc.args, *mockUserRepo, *mockTwilio, *mockHashPass, *mockJwt)
// 			result, err := userUseCase.UserSignUp(&tc.args)

// 			assert.Equal(t, tc.want, result)
// 			assert.Equal(t, tc.wantErr, err)
// 		})
// 	}
// }

func TestGetLatestUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockIuserRepo(mockCtrl)
	userUseCase := NewUserUseCase(mockUserRepo, &config.Token{})

	tests := []struct {
		name    string
		stub    func(*mocks.MockIuserRepo)
		want    *[]responsemodels.UserDetails
		wantErr error
	}{
		{
			name: "success",
			stub: func(userRepo *mocks.MockIuserRepo) {
				userRepo.EXPECT().GetLatestUsers().Times(1).Return(&[]responsemodels.UserDetails{{
					Id:    1,
					FName: "Ashkar",
					LName: "A.S",
					Email: "ashkar@example.com",
					Phone: "+910000000000",
				}, {
					Id:    2,
					FName: "Vajid",
					LName: "Hussain",
					Email: "vajid@example.com",
					Phone: "+910000000000",
				}}, nil)

			},
			want: &[]responsemodels.UserDetails{{
				Id:    1,
				FName: "Ashkar",
				LName: "A.S",
				Email: "ashkar@example.com",
				Phone: "+910000000000",
			}, {
				Id:    2,
				FName: "Vajid",
				LName: "Hussain",
				Email: "vajid@example.com",
				Phone: "+910000000000",
			}},
			wantErr: nil,
		},
		{
			name: "failure",
			stub: func(userRepo *mocks.MockIuserRepo) {
				userRepo.EXPECT().GetLatestUsers().Return(&[]responsemodels.UserDetails{}, errors.New("error fetching data"))

			},
			want:    &[]responsemodels.UserDetails{},
			wantErr: errors.New("error fetching data"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(mockUserRepo)
			result, err := userUseCase.GetLatestUsers()
			assert.Equal(t, tc.want, result)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestUserProfile(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockIuserRepo(mockCtrl)
	userUseCase := NewUserUseCase(mockUserRepo, &config.Token{})

	tests := []struct {
		name    string
		input   string
		stub    func(*mocks.MockIuserRepo, string)
		want    *responsemodels.UserDetails
		wantErr error
	}{
		{
			name:  "success",
			input: "1",
			stub: func(userRepo *mocks.MockIuserRepo, input string) {
				userRepo.EXPECT().GetUserProfile(&input).Times(1).Return(&responsemodels.UserDetails{
					Id:    1,
					FName: "Ashkar",
					LName: "A.S",
					Email: "ashkar@example.com",
					Phone: "+910000000000",
				}, nil)

			},
			want: &responsemodels.UserDetails{
				Id:    1,
				FName: "Ashkar",
				LName: "A.S",
				Email: "ashkar@example.com",
				Phone: "+910000000000",
			},
			wantErr: nil,
		},
		{
			name:  "failure",
			input: "1",
			stub: func(userRepo *mocks.MockIuserRepo, input string) {
				userRepo.EXPECT().GetUserProfile(&input).Times(1).Return(&responsemodels.UserDetails{}, errors.New("error fetching data"))
			},
			want:    &responsemodels.UserDetails{},
			wantErr: errors.New("error fetching data"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(mockUserRepo, tc.input)
			result, err := userUseCase.UserProfile(&tc.input)
			assert.Equal(t, tc.want, result)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func TestUserByStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockUserRepo := mocks.NewMockIuserRepo(mockCtrl)
	userUseCase := NewUserUseCase(mockUserRepo, &config.Token{})

	tests := []struct {
		name    string
		input   string
		stub    func(*mocks.MockIuserRepo, string)
		want    *[]responsemodels.UserDetails
		wantErr error
	}{
		{
			name:  "success",
			input: "pending",
			stub: func(userRepo *mocks.MockIuserRepo, input string) {
				userRepo.EXPECT().GetUserByStatus(input).Times(1).Return(&[]responsemodels.UserDetails{{
					Id:    1,
					FName: "Ashkar",
					LName: "A.S",
					Email: "ashkar@example.com",
					Phone: "+910000000000",
					Status: "pending",
				}, {
					Id:    2,
					FName: "Vajid",
					LName: "Hussain",
					Email: "vajid@example.com",
					Phone: "+910000000000",
					Status: "pending",

				}}, nil)

			},
			want: &[]responsemodels.UserDetails{{
				Id:    1,
				FName: "Ashkar",
				LName: "A.S",
				Email: "ashkar@example.com",
				Phone: "+910000000000",
				Status: "pending",

			}, {
				Id:    2,
				FName: "Vajid",
				LName: "Hussain",
				Email: "vajid@example.com",
				Phone: "+910000000000",
				Status: "pending",

			}},
			wantErr: nil,
		},
		{
			name:  "failure",
			input: "pending",
			stub: func(userRepo *mocks.MockIuserRepo, input string) {
				userRepo.EXPECT().GetUserByStatus(input).Times(1).Return(&[]responsemodels.UserDetails{}, errors.New("error fetching data"))
			},
			want:    &[]responsemodels.UserDetails{},
			wantErr: errors.New("error fetching data"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub(mockUserRepo, tc.input)
			result, err := userUseCase.UserByStatus(tc.input)
			assert.Equal(t, tc.want, result)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
