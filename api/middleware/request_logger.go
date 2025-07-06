package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		start := time.Now()

		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		utils.Info(
			"RequestID: " + requestID +
				" | " + c.Request.Method +
				" " + c.Request.URL.Path +
				" | Status: " + http.StatusText(status) +
				" | Duração: " + latency.String())
	}
}
