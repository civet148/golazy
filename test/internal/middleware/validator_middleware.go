package middleware

import (
	"github.com/gin-gonic/gin"
)

type ValidatorMiddleware struct {
}

func NewValidatorMiddleware() *ValidatorMiddleware {
	return &ValidatorMiddleware{}
}

func (m *ValidatorMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {

		//TODO: add your middleware logic here

		// Pass through to next handler
		c.Next()
	}
}
