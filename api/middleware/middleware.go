package middleware

import (
	"Trial/BANK-NOVANNA/pkg/jwttoken"
	"Trial/BANK-NOVANNA/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwttoken.TokenValid(c.Request)
		if err != nil {
			response.ResponseError(c, err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
