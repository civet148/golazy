package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CorsMiddleware struct {
}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func (m *CorsMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		//set header for cross-domain
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Auth-Token, *")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, content-Disposition")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Set("content-type", "application/json, text/plain, multipart/form-data, */*")

		// abort with options method (code=204)
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		//TODO: add your middleware logic here

		// Pass through to next handler
		c.Next()
	}
}
