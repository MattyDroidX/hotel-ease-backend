package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"status":   "success",
		"mensagem": "Requisição bem-sucedida",
		"dados":    data,
	})
}

func Fail(c *gin.Context, statusCode int, msg string, detail string) {
	c.JSON(statusCode, gin.H{
		"status":   "error",
		"mensagem": msg,
		"detalhe":  detail,
	})
}
