package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CAndresFernandez/go-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

func(h *Handler)GetBooks(c *gin.Context) {
	var books []models.Book
	if result := h.db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &books)
}

func (h *Handler) GetBookById(c *gin.Context) {
	var book models.Book
	bookId := c.Param("id")
	fmt.Println("unparsed:", bookId)
	// ID, err := strconv.ParseInt(bookId, 10, 64)
	// if err != nil {
	// 	fmt.Println("error while parsing")
	// }

	if result := h.db.Where("ID=?", bookId).Find(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *Handler) CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &book)
}

func (h *Handler) DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	if result := h.db.Delete(&models.Book{}, bookId); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) UpdateBook(c *gin.Context) {
	var book = &models.Book{}
	bookId := c.Param("id")

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	if result := h.db.First(&book, ID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Save(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}