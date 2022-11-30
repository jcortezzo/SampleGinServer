package util

import (
	"encoding/json"
	"sample-gin-server/models"
)

func GetGameFromJSON(jsonStr string) (*models.Game, error) {
	var game models.Game
	if err := json.Unmarshal([]byte(jsonStr), &game); err != nil {
		return nil, err
	}
	return &game, nil
}
