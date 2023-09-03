package usecase

//Usecase untuk mengambil data dari mysql dan menentukan class yang akan dipakai untuk handler
import (
	"github.com/muhammadsaefulr/simple-book-app/app/models"
	"gorm.io/gorm"
)

type BooksUseCase struct {
	DB *gorm.DB
}

func NewBooksCase(db *gorm.DB) *BooksUseCase {
	return &BooksUseCase{DB: db}
}

func (bu *BooksUseCase) CreateBooksList(books *models.Detail) error {
	if err := bu.DB.Create(books).Error; err != nil {
		return err
	}
	return nil
}

func (bu *BooksUseCase) GetBooksByID(id uint) (*models.Detail, error) {
	var books models.Detail

	if err := bu.DB.First(&books, id).Error; err != nil {
		return nil, err
	}
	return &books, nil
}

func (bu *BooksUseCase) GetBooksCategory(BookCategory string) ([]models.Detail, error) {
	var books []models.Detail

	res := bu.DB.Where("book_category = ?", BookCategory).Find(&books)

	if res.Error != nil {
		return nil, res.Error
	}

	return books, nil
}

func (bu *BooksUseCase) GetBooks() ([]models.Detail, error) {
	var books []models.Detail

	if err := bu.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (bu *BooksUseCase) UpdateBooks(books *models.Detail) error {
	return bu.DB.Save(books).Error
}

func (bu *BooksUseCase) DeleteBooks(books *models.Detail) error {
	return bu.DB.Delete(books).Error
}
