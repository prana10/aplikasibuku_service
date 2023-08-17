package impl

import (
	"errors"
	domainBook "service-api/src/main/app/book/domain"
	repositoryBook "service-api/src/main/app/book/repository"
	bookInput "service-api/src/main/util/input/book"
)

type bookService struct {
	bookRepository repositoryBook.BookRepositoy
}

func NewBookService(bookRepository repositoryBook.BookRepositoy) *bookService {
	return &bookService{bookRepository}
}

// insert book
func (service *bookService) InsertBook(book domainBook.Book) (domainBook.Book, error) {
	book, err := service.bookRepository.CreateBook(book)

	if err != nil {
		return book, err
	}

	return book, nil
}

// get all book
func (service *bookService) GetAllBook() ([]domainBook.Book, error) {
	books, err := service.bookRepository.GetAllBook()

	if err != nil {
		return books, err
	}

	if len(books) == 0 {
		return books, errors.New("book not found")
	}

	return books, nil
}

// get book by id
func (service *bookService) GetBookByID(bookID uint) (domainBook.Book, error) {
	book, err := service.bookRepository.GetBookByID(bookID)

	if err != nil {
		return book, err
	}

	if book.ID == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

// update book by id
func (service *bookService) UpdateBookByID(bookID uint, book bookInput.BookInput) (domainBook.Book, error) {
	var bookUpdate domainBook.Book
	bookUpdate.ID = bookID
	bookUpdate.Title = book.Title
	bookUpdate.Author = book.Author
	bookUpdate.Price = book.Price
	bookUpdate.Isbn = book.Isbn
	bookUpdate.Genre = book.Genre
	bookUpdate.Description = book.Description

	newBook, err := service.bookRepository.UpdateBookByID(bookID, bookUpdate)

	if err != nil {
		return newBook, err
	}

	if newBook.ID == 0 {
		return newBook, errors.New("book not found")
	}

	return newBook, nil
}

// delete book by id
func (service *bookService) DeleteBookByID(bookID uint) (domainBook.Book, error) {
	book, err := service.bookRepository.DeleteBookByID(bookID)

	if err != nil {
		return book, err
	}

	if book.ID == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}
