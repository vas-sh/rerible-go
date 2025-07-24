package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vas-sh/rerible-go/internal/client"
	"github.com/vas-sh/rerible-go/internal/config"
	"github.com/vas-sh/rerible-go/internal/handlers"
)

func main() {
	client := client.New(config.Config.ApiKey, &http.Client{}, config.Config.RaribleRootURL)
	handler := handlers.New(client)
	r := gin.Default()
	handler.Register(r)

	log.Println("Server started")
	if err := r.Run(":" + config.Config.Port); err != nil {
		log.Fatal(err)
	}
}
