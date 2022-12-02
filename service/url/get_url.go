package url

import (
	"fmt"
	"net/http"
	"sample-gin-server/clients"

	"github.com/gin-gonic/gin"
)

func GetURL(ctx *gin.Context) {
	shortCode := ctx.Param("shortCode")

	if !clients.URLDB.Contains(shortCode) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"errors": fmt.Sprintf("Short code %v not found!", shortCode),
		})
	}

	url := clients.URLDB.Get(shortCode)

	ctx.Redirect(http.StatusMovedPermanently, url)
}
