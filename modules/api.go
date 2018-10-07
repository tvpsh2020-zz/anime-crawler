package modules

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tvpsh2020/anime-crawler/config"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("api")

	api.GET("/stat", func(c *gin.Context) {
		c.JSON(http.StatusOK, AnimeQueueList)
	})

	api.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, config.Anime)
	})

	return router
}
