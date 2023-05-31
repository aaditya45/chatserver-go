package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SystemAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"Aaditya": "123456",
	})
}
