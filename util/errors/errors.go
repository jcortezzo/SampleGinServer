package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnInternalErrorResponse(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"errors": "internal error, please try again later",
	})
}
