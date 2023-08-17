package repository

import (
	domainBook "service-api/src/main/app/book/domain"
)

type BookRepositoy interface {
	CreateBook(book domainBook.Book) (domainBook.Book, error)
	GetAllBook() ([]domainBook.Book, error)
	GetBookByID(bookID uint) (domainBook.Book, error)
	UpdateBookByID(bookID uint, book domainBook.Book) (domainBook.Book, error)
	DeleteBookByID(bookID uint) (domainBook.Book, error)
}
