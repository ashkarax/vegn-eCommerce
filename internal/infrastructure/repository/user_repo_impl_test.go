package repository

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	responsemodels "github.com/ashkarax/vegn-eCommerce/internal/models/response_models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUserProfile(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    *responsemodels.UserDetails
		wantErr error
	}{
		{
			name: "succesfully got details",
			args: "1",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery("SELECT id,f_name,l_name,email,phone,status FROM users WHERE id = ?").
					WillReturnRows(sqlmock.NewRows([]string{"id", "f_name", "l_name", "email", "phone", "status"}).AddRow(1, "Ashkar", "A.S", "ashkar@example.com", "+918921791915", "active"))
			},
			want: &responsemodels.UserDetails{
				Id:     1,
				FName:  "Ashkar",
				LName:  "A.S",
				Email:  "ashkar@example.com",
				Phone:  "+918921791915",
				Status: "active",
			},
			wantErr: nil,
		},
		{
			name: "no user found",
			args: "1",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectQuery("SELECT id,f_name,l_name,email,phone,status FROM users WHERE id = ?").
					WillReturnRows(sqlmock.NewRows([]string{"id", "f_name", "l_name", "email", "phone", "status"}))
			},
			want: &responsemodels.UserDetails{
				Id:     0,
				FName:  "",
				LName:  "",
				Email:  "",
				Phone:  "",
				Status: "",
			},
			wantErr: errors.New("no results found,Rows affected 0"),
		},
	}

	for _, index := range tests {
		db, mock, _ := sqlmock.New()

		defer db.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})

		index.stub(mock)

		userRepository := NewUserRepository(DB)
		result, err := userRepository.GetUserProfile(&index.args)

		assert.Equal(t, index.want, result)
		assert.Equal(t, index.wantErr, err)

	}
}

func TestIsUserExist(t *testing.T) {
	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user exist",
			args: "+918921791915",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE phone =$1 AND status =$2")).
					WithArgs("+918921791915", "pending").
					WillReturnResult(sqlmock.NewResult(0, 1))

				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2")).
					WithArgs("+918921791915", "deleted").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			want: true,
		},
		{
			name: "user does not exist",
			args: "+918921791915",
			stub: func(mocksql sqlmock.Sqlmock) {
				mocksql.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE phone =$1 AND status =$2")).
					WithArgs("+918921791915", "pending").
					WillReturnResult(sqlmock.NewResult(0, 0))

				mocksql.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM users WHERE phone=$1 AND status!=$2")).
					WithArgs("+918921791915", "deleted").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
			},
			want: false,
		},
	}

	for _, tc := range tests {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("Error creating mock DB: %v", err)
		}
		defer db.Close()

		DB, _ := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{})

		tc.stub(mock)

		userRepository := NewUserRepository(DB)
		result := userRepository.IsUserExist(tc.args)

		assert.Equal(t, tc.want, result, "Test case %q failed: expected %t, got %t", tc.name, tc.want, result)

	}
}
