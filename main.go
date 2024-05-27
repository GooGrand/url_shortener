package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/googrand/go-url-shortener/handler"
	"github.com/googrand/go-url-shortener/store"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hey",
		})
	})
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})
	store.InitializeStore()
	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
