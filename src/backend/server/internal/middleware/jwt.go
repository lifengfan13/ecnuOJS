package middleware

import (
	"server/internal/util"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Telephone string `json:"telephone"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var mySecret = []byte("欧豪你是我的神")

func GenToken(telephone string) (string, error) {
	c := MyClaims{
		telephone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "paper",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(mySecret)
}

func ParseToken(tokenString string) (*MyClaims, int) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})

	if err != nil {
		// log
		return nil, util.ServerInternalError
	}

	if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
		return claims, util.RequestSuccess
	}
	return nil, util.AuthTokenInvalid
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.RespErr(c, util.AuthTokenError)
			c.Abort()
			return
		}
		myClaim, code := ParseToken(authHeader)
		if code != util.RequestSuccess {
			util.RespErr(c, code)
			c.Abort()
			return
		}

		c.Set("telephone", myClaim.Telephone)
		c.Next()

	}
}
