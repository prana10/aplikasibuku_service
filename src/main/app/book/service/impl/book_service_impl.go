package impl

import (
	"errors"
	domainBook "service-api/src/main/app/book/domain"
	repositoryBook "service-api/src/main/app/book/repository"
	domainGenre "service-api/src/main/app/genre/domain"
	repositoryGenre "service-api/src/main/app/genre/repository"
	inputBook "service-api/src/main/util/input/book"

	"gorm.io/gorm"
)

type bookService struct {
	bookRepository  repositoryBook.BookRepositoy
	genreRepository repositoryGenre.GenreRepository
}

func NewBookService(bookRepository repositoryBook.BookRepositoy, genreRepository repositoryGenre.GenreRepository) *bookService {
	return &bookService{bookRepository, genreRepository}
}

// insert book
func (service *bookService) InsertBook(bookInput inputBook.BookInput) (domainBook.Book, error) {
	var genreID []uint
	for _, genre := range bookInput.Genre {
		idGenre, err := service.genreRepository.GetGenreByID(uint(genre))
		if err != nil {
			return domainBook.Book{}, err
		}

		if idGenre.ID == 0 {
			return domainBook.Book{}, errors.New("genre not found")
		}

		genreID = append(genreID, uint(genre))
	}

	newBook := domainBook.Book{}
	newBook.Title = bookInput.Title
	newBook.Author = bookInput.Author
	newBook.Price = bookInput.Price
	newBook.Isbn = bookInput.Isbn
	newBook.Description = bookInput.Description

	// append genre by id to newBook
	for _, genre := range genreID {
		newBook.Genre = append(newBook.Genre, domainGenre.Genre{Model: gorm.Model{
			ID: genre,
		}})
	}

	book, err := service.bookRepository.CreateBook(newBook)

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
func (service *bookService) UpdateBookByID(bookID uint, book inputBook.BookInput) (domainBook.Book, error) {
	var genreID []uint
	for _, genre := range book.Genre {
		idGenre, err := service.genreRepository.GetGenreByID(uint(genre))
		if err != nil {
			return domainBook.Book{}, err
		}

		if idGenre.ID == 0 {
			return domainBook.Book{}, errors.New("genre not found")
		}

		genreID = append(genreID, uint(genre))
	}

	var bookUpdate domainBook.Book
	bookUpdate.ID = bookID
	bookUpdate.Title = book.Title
	bookUpdate.Author = book.Author
	bookUpdate.Price = book.Price
	bookUpdate.Isbn = book.Isbn
	bookUpdate.Description = book.Description

	// append genre by id to newBook
	for _, genre := range genreID {
		bookUpdate.Genre = append(bookUpdate.Genre, domainGenre.Genre{Model: gorm.Model{
			ID: genre,
		}})
	}

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
