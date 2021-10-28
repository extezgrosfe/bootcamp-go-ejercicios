package middleware

import (
	"errors"
	"goweb_clase4_tt/pkg/web"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	ErrUnauthorized = errors.New("acceso denegado")
)

func TokenMiddleware() gin.HandlerFunc {
	expectedToken := os.Getenv("API_TOKEN")

	if expectedToken == "" {
		panic("API_TOKEN is not set")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		isValid := false

		if token == expectedToken {
			isValid = true
		}

		if !isValid {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, ErrUnauthorized))
			c.Abort()
			return
		}

		c.Next()
	}
}
