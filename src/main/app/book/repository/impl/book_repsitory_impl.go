package impl

import (
	domainBook "service-api/src/main/app/book/domain"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

// create book
func (repo *bookRepository) CreateBook(book domainBook.Book) (domainBook.Book, error) {
	err := repo.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

// get all book
func (repo *bookRepository) GetAllBook() ([]domainBook.Book, error) {
	var books []domainBook.Book

	err := repo.db.Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

// get book by id
func (repo *bookRepository) GetBookByID(bookID uint) (domainBook.Book, error) {
	var book domainBook.Book

	err := repo.db.Where("id = ?", bookID).First(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

// update book by id
func (repo *bookRepository) UpdateBookByID(bookID uint, book domainBook.Book) (domainBook.Book, error) {
	var bookData domainBook.Book

	err := repo.db.Model(&bookData).Where("id = ?", bookID).Updates(book).Error

	if err != nil {
		return bookData, err
	}

	return bookData, nil
}

// delete book by id
func (repo *bookRepository) DeleteBookByID(bookID uint) (domainBook.Book, error) {
	var book domainBook.Book

	err := repo.db.Where("id = ?", bookID).Delete(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}
