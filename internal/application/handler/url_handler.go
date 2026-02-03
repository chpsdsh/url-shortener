package handler

import (
	"net/http"
	"url-shortener/internal/application/generator"
	"url-shortener/internal/domain"
	"url-shortener/internal/infrastructure/store"

	"github.com/gin-gonic/gin"
)

const host = "http://localhost:9808/"

type UrlHandler struct {
	Storage store.StorageService
}

func (h UrlHandler) CreateShortUrl(c *gin.Context) {
	var creationRequest domain.URLRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := generator.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	if err := h.Storage.SaveUrlMapping(shortUrl, creationRequest.LongUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func (h UrlHandler) ShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl, err := h.Storage.RetrieveOriginalUrl(shortUrl)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Redirect(http.StatusFound, longUrl)
}
