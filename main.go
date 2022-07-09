package main

import (
	"go_http/handler"
	"go_http/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	storage := storage.NewRunTimeMemoryStorage()
	handler := handler.NewHandler(storage)
	router.POST("/author", handler.CreateAuthors)
	router.GET("/author/:id", handler.GetAuthor)
	// router.GET("/authors")
	router.PUT("/author/:id", handler.UpdateAuthors)
	router.DELETE("/author/:id", handler.DeleteAuthors)

	router.Run()
}
