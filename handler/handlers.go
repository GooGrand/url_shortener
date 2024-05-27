package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/googrand/go-url-shortener/shortener"
	"github.com/googrand/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var CreationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&CreationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	shortUrl := shortener.GenerateShortLink(CreationRequest.LongUrl, CreationRequest.UserId)
	store.SaveUrlMapping(shortUrl, CreationRequest.LongUrl, CreationRequest.UserId)
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
