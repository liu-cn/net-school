package home

import (
	"github.com/gin-gonic/gin"
)

func Home() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})

	}
}
