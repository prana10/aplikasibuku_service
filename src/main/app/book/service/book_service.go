package service

import (
	domainBook "service-api/src/main/app/book/domain"
	bookInput "service-api/src/main/util/input/book"
)

type BookService interface {
	InsertBook(book bookInput.BookInput) (domainBook.Book, error)
	GetAllBook() ([]domainBook.Book, error)
	GetBookByID(bookID uint) (domainBook.Book, error)
	UpdateBookByID(bookID uint, book bookInput.BookInput) (domainBook.Book, error)
	DeleteBookByID(bookID uint) (domainBook.Book, error)
}
