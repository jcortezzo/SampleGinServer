package game

import (
	"encoding/json"
	"net/http"
	"sample-gin-server/clients"
	"sample-gin-server/models"
	"sample-gin-server/util/errors"

	"github.com/gin-gonic/gin"
)

type ListGamesResponse struct {
	Games []models.Game `json:"games"`
}

func ListGames(ctx *gin.Context) {
	result := make([]models.Game, len(clients.GameDB))
	i := 0
	for k := range clients.GameDB {
		var game models.Game
		if err := json.Unmarshal([]byte(k), &game); err != nil {
			errors.ReturnInternalErrorResponse(ctx)
			return
		}
		result[i] = game
		i++
	}
	resp := &ListGamesResponse{
		Games: result,
	}
	ctx.JSON(http.StatusOK, resp)
}
