package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/simple-book-app/app/models"
	"github.com/muhammadsaefulr/simple-book-app/app/usecase"
)

type BooksHandler struct {
	BooksUseCase *usecase.BooksUseCase
}

// bu *usecase.BooksUseCase Untuk mengambil data dari mysql dan menentukan berdasarkan handler nya
func NewBooksHandler(r *gin.RouterGroup, bu *usecase.BooksUseCase) {
	handler := &BooksHandler{
		BooksUseCase: bu,
	}

	r.GET("/books", handler.GetBooks)
	r.GET("/books/:category", handler.GetBooksCategory)
	r.POST("/books", handler.CreateBooksList)
	r.PUT("/books/:id", handler.UpdateBooks)
	r.DELETE("/books/:id", handler.DeleteBook)
}

func (bu *BooksHandler) CreateBooksList(c *gin.Context) {
	var books models.Detail

	if err := c.ShouldBindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bu.BooksUseCase.CreateBooksList(&books); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func (bu *BooksHandler) GetBooks(c *gin.Context) {
	books, err := bu.BooksUseCase.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bu *BooksHandler) GetBooksCategory(c *gin.Context) {
	BookCategory := c.Param("category")

	book, err := bu.BooksUseCase.GetBooksCategory(BookCategory)
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

	book, err := bu.BooksUseCase.GetBooksByID(uint(id))
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

	if err := bu.BooksUseCase.UpdateBooks(&book); err != nil {
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

	book, err := bu.BooksUseCase.GetBooksByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := bu.BooksUseCase.DeleteBooks(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
