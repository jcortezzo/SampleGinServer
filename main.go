package main

import (
	"fmt"
	"net/http"
	"sample-gin-server/clients"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		// c.Redirect(http.StatusMovedPermanently, "https://www.singularity6.com")
	})

	setURLRoutes(r)
	setGameRoutes(r)

	clients.Init()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setURLRoutes(r *gin.Engine) {
	r.POST("/", ShortenURL)
	r.GET("/:shortCode", GetURL)
	r.GET("/list", func(ctx *gin.Context) {
		var res []string
		for k, v := range clients.URLDB {
			res = append(res, fmt.Sprintf("%v:%v", k, v))
		}
		ctx.JSON(http.StatusOK, res)
	})
}

func setGameRoutes(r *gin.Engine) {
	gameRouter := r.Group("game")
	gameRouter.POST("/add", MakeGame)
	gameRouter.GET("/list", ListGames)
	gameRouter.POST("/search", SearchGame)
}
