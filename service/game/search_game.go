package game

import (
	"net/http"
	"sample-gin-server/clients"
	"sample-gin-server/models"
	"sample-gin-server/service/game/util"
	"sample-gin-server/util/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	titleQueryKey = "title"
)

type SearchGameResponse struct {
	Games []models.Game `json:"games"`
}

func SearchGame(ctx *gin.Context) {
	searchTitle, ok := ctx.GetQuery(titleQueryKey)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": "You must specify a title as a query parameter!",
		})
		return
	}
	result := make([]models.Game, 0)
	for k := range clients.GameDB {
		gamePtr, err := util.GetGameFromJSON(k)
		if gamePtr == nil || err != nil {
			errors.ReturnInternalErrorResponse(ctx)
			return
		}
		game := *gamePtr
		title := strings.ToLower(game.Title)
		if strings.Contains(title, searchTitle) {
			result = append(result, game)
		}
	}
	resp := &SearchGameResponse{
		Games: result,
	}
	ctx.JSON(http.StatusOK, resp)
}
