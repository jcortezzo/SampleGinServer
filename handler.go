package main

import (
	"sample-gin-server/service/game"

	"github.com/gin-gonic/gin"
)

func MakeGame(ctx *gin.Context) {
	game.CreateGame(ctx)
}

func ListGames(ctx *gin.Context) {
	game.ListGames(ctx)
}

func SearchGame(ctx *gin.Context) {
	game.SearchGame(ctx)
}
