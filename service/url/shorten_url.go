package url

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sample-gin-server/clients"
	"sample-gin-server/util/errors"

	"github.com/gin-gonic/gin"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
}

type ShortenURLResponse struct {
	URL          string `json:"url"`
	ShortUrlCode string `json:"short_url_code"`
}

func ShortenURL(ctx *gin.Context) {
	var req ShortenURLRequest
	err := ctx.Bind(&req)
	if err != nil {
		errors.ReturnInternalErrorResponse(ctx)
	}
	_, err = url.ParseRequestURI(req.URL)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": fmt.Sprintf("Bad URL=%v", req.URL),
		})
		return
	}

	resp := &ShortenURLResponse{
		URL:          req.URL,
		ShortUrlCode: getShortenedUrlCode(req.URL),
	}

	clients.URLDB.Add(resp.ShortUrlCode, resp.URL)

	ctx.JSON(http.StatusOK, resp)
}

func getShortenedUrlCode(s string) string {
	var maxLength int
	if len(s) < 3 {
		maxLength = len(s)
	} else {
		maxLength = 3
	}
	shortString := s[:maxLength]
	randNums := rand.Intn(999)
	return fmt.Sprintf("%v%v", shortString, randNums)
}
