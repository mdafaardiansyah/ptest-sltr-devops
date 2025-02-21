package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/ptest-sltr-devops/internal/handler"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/welcome/:nama", handler.WelcomeHandler)
	r.GET("/welcome", handler.WelcomeHandler)

	return r
}
