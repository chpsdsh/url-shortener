package rest

import (
	"fmt"
	"net/http"
	"url-shortener/internal/application/handler"

	"github.com/gin-gonic/gin"
)

type Client struct {
	engine  *gin.Engine
	handler handler.UrlHandler
}

func NewRestClient(engine *gin.Engine, handler handler.UrlHandler) Client {
	return Client{engine: engine, handler: handler}
}

func (c Client) StartEndpoints() error {
	c.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello from url-shortener"})
	})

	c.engine.POST("/create-short-url", c.handler.CreateShortUrl)

	c.engine.GET("/:shortUrl", c.handler.ShortUrlRedirect)

	if err := c.engine.Run(":9808"); err != nil {
		return fmt.Errorf("failed to start the web server - Error: %v", err)
	}
	return nil
}
