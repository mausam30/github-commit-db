package github

import (
	"time"

	"github.com/gin-gonic/gin"
)

var rateLimit = time.NewTicker(time.Second)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		<-rateLimit.C
		c.Next()
	}
}
