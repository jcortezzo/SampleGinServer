package main

import (
	"sample-gin-server/service/game"
	"sample-gin-server/service/url"

	"github.com/gin-gonic/gin"
)

func ShortenURL(ctx *gin.Context) {
	url.ShortenURL(ctx)
}

func GetURL(ctx *gin.Context) {
	url.GetURL(ctx)
}

func MakeGame(ctx *gin.Context) {
	game.CreateGame(ctx)
}

func ListGames(ctx *gin.Context) {
	game.ListGames(ctx)
}

func SearchGame(ctx *gin.Context) {
	game.SearchGame(ctx)
}
