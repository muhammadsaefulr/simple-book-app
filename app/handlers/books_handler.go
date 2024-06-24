package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/simple-book-app/app/models"
	Service "github.com/muhammadsaefulr/simple-book-app/app/service"
)

type BooksHandler struct {
	BooksService *Service.BooksService
}

func NewBooksHandler(r *gin.RouterGroup, bu *Service.BooksService) {
	handler := &BooksHandler{
		BooksService: bu,
	}

	r.GET("/books", handler.GetBooks)
	r.GET("/books/:category", handler.GetBooksCategory)
	r.POST("/books", handler.CreateBooksList)
	r.PUT("/books/:id", handler.UpdateBooks)
	r.DELETE("/books/:id", handler.DeleteBook)
}

func (bu *BooksHandler) CreateBooksList(c *gin.Context) {
	var newBook models.Detail

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingBook, err := bu.BooksService.GetBookByTitle(newBook.BookTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if existingBook != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Buku dengan judul yang sama, sudah ada"})
		return
	}

	if err := bu.BooksService.CreateBooksList(&newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func (bu *BooksHandler) GetBooks(c *gin.Context) {
	books, err := bu.BooksService.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bu *BooksHandler) GetBooksCategory(c *gin.Context) {
	BookCategory := c.Param("category")

	book, err := bu.BooksService.GetBooksCategory(BookCategory)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Not Found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bu *BooksHandler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, err := bu.BooksService.GetBooksByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bu *BooksHandler) UpdateBooks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Detail
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)

	if err := bu.BooksService.UpdateBooks(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bu *BooksHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	book, err := bu.BooksService.GetBooksByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := bu.BooksService.DeleteBooks(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Berhasil Menghapus Buku Dengan Id %d", id)})
	return
}
