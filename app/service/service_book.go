package Service

//Service untuk mengambil data dari mysql dan menentukan class yang akan dipakai untuk handler
import (
	"errors"

	"github.com/muhammadsaefulr/simple-book-app/app/models"
	"gorm.io/gorm"
)

type BooksService struct {
	DB *gorm.DB
}

func NewBooksCase(db *gorm.DB) *BooksService {
	return &BooksService{DB: db}
}

func (bu *BooksService) CreateBooksList(books *models.Detail) error {
	if err := bu.DB.Create(books).Error; err != nil {
		return err
	}
	return nil
}

func (bs *BooksService) GetBookByTitle(title string) (*models.Detail, error) {
	var book models.Detail
	if err := bs.DB.Where("book_title = ?", title).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &book, nil
}

func (bu *BooksService) GetBooksByID(id uint) (*models.Detail, error) {
	var books models.Detail

	if err := bu.DB.First(&books, id).Error; err != nil {
		return nil, err
	}
	return &books, nil
}

func (bu *BooksService) GetBooksCategory(BookCategory string) ([]models.Detail, error) {
	var books []models.Detail

	res := bu.DB.Where("book_category = ?", BookCategory).Find(&books)

	if res.Error != nil {
		return nil, res.Error
	}

	return books, nil
}

func (bu *BooksService) GetBooks() ([]models.Detail, error) {
	var books []models.Detail

	if err := bu.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (bu *BooksService) UpdateBooks(books *models.Detail) error {
	return bu.DB.Save(books).Error
}

func (bu *BooksService) DeleteBooks(books *models.Detail) error {
	return bu.DB.Delete(books).Error
}
