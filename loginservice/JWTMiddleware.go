package loginservice

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthrorizeJWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		const BAREAR_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BAREAR_SCHEMA):]
		token, err := JWTAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		}else{
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}