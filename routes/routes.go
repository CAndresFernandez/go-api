package routes

import (
	"github.com/CAndresFernandez/go-api/controllers"
	"github.com/gin-gonic/gin"
)

var BookstoreRoutes = func(router *gin.Engine, handler *controllers.Handler) {
	router.GET("/books", handler.GetBooks)
	router.GET("/books/:id", handler.GetBookById)
	router.POST("/books", handler.CreateBook)
	router.PUT("/books/:id",handler.UpdateBook)
	router.DELETE("/books/:id", handler.DeleteBook)
}