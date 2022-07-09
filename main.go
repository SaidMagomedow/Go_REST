package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/author")
	router.GET("/author/:id")
	router.GET("/authors")
	router.PUT("/author/:id")
	router.DELETE("/author/:id")

	router.Run()
}
