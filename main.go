package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/fcc-file-metadata-ms/handlers"
	"github.com/madjiebimaa/fcc-file-metadata-ms/middlewares"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.CORS())

	fileHandler := handlers.NewFileHandler()

	r.POST("/api/files", fileHandler.Analysis)

	if err := r.Run(":3000"); err != nil {
		log.Fatal("can't connect to gin server port 3000")
	}
}
