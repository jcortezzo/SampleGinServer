package main

import (
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
	})

	gameRouter := r.Group("game")
	gameRouter.POST("/add", MakeGame)
	gameRouter.GET("/list", ListGames)

	clients.Init()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
