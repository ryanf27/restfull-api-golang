package routes

import (
	"restfull-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine)  {

	router.GET("/books", handlers.GetBooks)
	router.GET("/book/:id", handlers.GetBookById)
	router.POST("/books", handlers.AddBook)
}