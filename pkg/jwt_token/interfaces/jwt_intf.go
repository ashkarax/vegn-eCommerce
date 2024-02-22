package interfaceJwt

type IJwt interface {
	GenerateRefreshToken(secretKey string) (string, error)
	GenerateAcessToken(securityKey string, id string) (string, error)
	TempTokenForOtpVerification(securityKey string, phone string) (string, error)
	UnbindPhoneFromClaim(tokenString string, tempVerificationKey string) (string, error)
	VerifyRefreshToken(accesToken string, securityKey string) error
	VerifyAccessToken(token string, secretkey string) (string, error)
}
