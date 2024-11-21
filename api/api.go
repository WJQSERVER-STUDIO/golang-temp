package api

import (
	"go/config"

	"github.com/gin-gonic/gin"
)

func InitHandleRouter(cfg *config.Config, router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
		apiRouter.Any("/v1", func(c *gin.Context) {

		})

	}
}
