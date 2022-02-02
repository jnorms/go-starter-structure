package routes

import (
	"github.com/gin-gonic/gin"
)

func (Routes *Engine) Api() {
	v1 := Routes.Engine.Group("/api")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
