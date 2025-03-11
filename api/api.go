package api

import (
	"golang-temp/config"

	"github.com/gin-gonic/gin"
)

func InitHandleRouter(cfg *config.Config, router *gin.Engine, version string) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
		apiRouter.GET("/version", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": version,
			})
		})

		apiRouter.Any("/v1", func(c *gin.Context) {

		})

	}
}
