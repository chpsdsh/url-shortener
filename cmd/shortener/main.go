package main

import (
	"fmt"
	"os"
	"url-shortener/internal/application/handler"
	"url-shortener/internal/infrastructure/rest"
	"url-shortener/internal/infrastructure/store"

	"github.com/gin-gonic/gin"
)

func main() {
	storage, err := store.NewStorageService()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer storage.Close()

	urlHandler := handler.UrlHandler{Storage: storage}

	engine := gin.Default()
	restClient := rest.NewRestClient(engine, urlHandler)
	if err := restClient.StartEndpoints(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
