package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ashkarax/vegn-eCommerce/internal/config"
	interfaceUseCase "github.com/ashkarax/vegn-eCommerce/internal/infrastructure/usecase/interfaces"
	jwttoken "github.com/ashkarax/vegn-eCommerce/pkg/jwt_token"
	"github.com/gin-gonic/gin"
)

type TokenRequirements struct {
	keys       *config.Token
	JWTUseCase interfaceUseCase.IJWTUseCase
}

func NewJWTTokenMiddleware(JWTUseCase interfaceUseCase.IJWTUseCase, Keys *config.Token) *TokenRequirements {
	return &TokenRequirements{JWTUseCase: JWTUseCase, keys: Keys}
}

func (r *TokenRequirements) AdminAuthorization(c *gin.Context) {
	adminRefToken := c.GetHeader("refreshtoken")

	err := jwttoken.VerifyRefreshToken(adminRefToken, r.keys.AdminSecurityKey)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		c.Abort()
	}
	c.Next()
}

func (r *TokenRequirements) RestaurantAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("accesstoken")
	refreshToken := c.Request.Header.Get("refreshtoken")

	if refreshToken == "" || accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "In your request,The Required tokens to get into this page are not available."})
		c.Abort()
		return
	}

	restaurantId, err := jwttoken.VerifyAccessToken(accessToken, r.keys.RestaurantSecurityKey)
	restaurantIdInt, _ := strconv.Atoi(restaurantId)
	if err != nil {
		if restaurantId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Token Tampared ,Id not accessible"})
			c.Abort()
			return
		}

		errn := jwttoken.VerifyRefreshToken(refreshToken, r.keys.RestaurantSecurityKey)
		if errn != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": errn.Error()})
			c.Abort()
			return
		}

		_, err1 := r.JWTUseCase.GetRestStatForAccessToken(&restaurantIdInt)
		if err1 != nil {
			c.JSON(http.StatusUnauthorized, err1.Error())
			c.Abort()
			return
		}
		newAcessToken, err2 := jwttoken.GenerateAcessToken(r.keys.RestaurantSecurityKey, restaurantId)
		if err2 != nil {
			c.JSON(http.StatusUnauthorized, err2.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"NewAccessToken": newAcessToken})
		c.Set("RestaurantId", restaurantId)
		c.Next()
		return
	}

	c.Set("RestaurantId", restaurantId)
	fmt.Println("access token is upto date")
	c.Next()

}

func (r *TokenRequirements) UserAuthorization(c *gin.Context) {
	accessToken := c.Request.Header.Get("accesstoken")
	refreshToken := c.Request.Header.Get("refreshtoken")

	if refreshToken == "" || accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"err": "In your request,The Required tokens to get into this page are not available."})
		c.Abort()
		return
	}

	userId, err := jwttoken.VerifyAccessToken(accessToken, r.keys.UserSecurityKey)
	if err != nil {
		if userId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "Token Tampared ,Id not accessible"})
			c.Abort()
			return
		}

		errn := jwttoken.VerifyRefreshToken(refreshToken, r.keys.UserSecurityKey)
		if errn != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
			c.Abort()
			return
		}

		_, err1 := r.JWTUseCase.GetUserStatForGeneratingAccessToken(&userId)
		if err1 != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		newAcessToken, err2 := jwttoken.GenerateAcessToken(r.keys.UserSecurityKey, userId)
		if err2 != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"NewAccessToken": newAcessToken})
		c.Set("userId", userId)
		c.Next()
		return
	}

	c.Set("userId", userId)
	fmt.Println("access token is upto date")
	c.Next()

}
