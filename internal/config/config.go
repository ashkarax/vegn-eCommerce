package config

import (
	"github.com/spf13/viper"
)

type PortManager struct {
	RunnerPort string `mapstructure:"PORTNO"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Token struct {
	AdminSecurityKey      string `mapstructure:"ADMIN_TOKENKEY"`
	RestaurantSecurityKey string `mapstructure:"RESTAURANT_TOKENKEY"`
	UserSecurityKey       string `mapstructure:"USER_TOKENKEY"`
	TempVerificationKey   string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type OTP struct {
	AccountSid string `mapstructure:"TWILIO_ACCOUNT_SID"`
	AuthToken  string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ServiceSid string `mapstructure:"TWILIO_SERVICE_SID"`
}
type AWS struct {
	Region     string `mapstructure:"AWS_REGION"`
	AccessKey  string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecrectKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	Endpoint   string `mapstructure:"AWS_ENDPOINT"`
}

type RazorPay struct {
	KeyId      string `mapstructure:"KEY_ID"`
	SecrectKey string `mapstructure:"KEY_SECRET"`
}

type Config struct {
	DB       DataBase
	Token    Token
	Otp      OTP
	AwsS3    AWS
	RazorP   RazorPay
	PortMngr PortManager
}

func LoadConfig() (*Config, error) {

	var db DataBase
	var token Token
	var otp OTP
	var awsS3 AWS
	var razorP RazorPay
	var portmngr PortManager

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&token)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&otp)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&awsS3)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&razorP)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&portmngr)
	if err != nil {
		return nil, err
	}

	config := Config{DB: db, Token: token, Otp: otp, AwsS3: awsS3, RazorP: razorP, PortMngr: portmngr}
	return &config, nil
}
