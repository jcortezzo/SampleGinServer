package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sample-gin-server/clients"
	"sample-gin-server/models"
	"sample-gin-server/util/errors"

	"github.com/gin-gonic/gin"
)

type CreateGameRequest struct {
	Game models.Game `json:"game"`
}

type CreateGameResponse CreateGameRequest

func CreateGame(ctx *gin.Context) {
	var gameRequest CreateGameRequest
	var err error
	if err = ctx.BindJSON(&gameRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": fmt.Sprintf("%+v", err),
		})
		return
	}

	gameJSON, err := json.Marshal(gameRequest.Game)
	if err != nil {
		errors.ReturnInternalErrorResponse(ctx)
		return
	}
	gameJSONString := string(gameJSON)

	clients.GameDB.Add(gameJSONString)

	resp := &CreateGameResponse{
		Game: gameRequest.Game,
	}
	ctx.JSON(http.StatusOK, resp)
}
