package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadsaefulr/simple-book-app/app/handlers"
	"github.com/muhammadsaefulr/simple-book-app/app/service"
	"github.com/muhammadsaefulr/simple-book-app/db"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := db.NewDatabase()
	if err != nil {
		panic("Failed to connect to database")
	}

	todoService := Service.NewBooksCase(db)

	// Menghubungkan handler dengan router
	v1 := r.Group("/api/v1")
	{
		handlers.NewBooksHandler(v1, todoService)
	}

	return r
}
