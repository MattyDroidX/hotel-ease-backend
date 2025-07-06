package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
)

func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log := "[" + param.TimeStamp.Format(time.RFC1123) + "] " +
			param.Method + " " + param.Path + " - " +
			param.ClientIP + " - " +
			param.Latency.String()
		utils.Info(log)
		return ""
	})
}
