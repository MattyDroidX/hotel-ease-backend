package middleware

import (
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		start     := time.Now()

		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()

		line := fmt.Sprintf("request=%s | %s %s | status=%d(%s) | %v",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			http.StatusText(c.Writer.Status()),
			time.Since(start),
		)

		switch status := c.Writer.Status(); {
		case status >= 500:
			utils.Error(line)
		case status >= 400:
			utils.Warn(line)
		default:
			utils.Info(line)
		}
	}
}
