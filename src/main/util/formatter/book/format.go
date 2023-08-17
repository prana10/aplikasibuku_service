package book

import (
	domainBook "service-api/src/main/app/book/domain"
	domainGenre "service-api/src/main/app/genre/domain"
)

type BookFormatter struct {
	ID          uint                `json:"id"`
	Title       string              `json:"title"`
	Author      string              `json:"author"`
	Price       float64             `json:"price"`
	Isbn        int                 `json:"isbn"`
	Description string              `json:"description"`
	Genre       []domainGenre.Genre `json:"genre"`
}

func FormatBook(book domainBook.Book) BookFormatter {
	var bookFormatter BookFormatter

	bookFormatter.ID = book.ID
	bookFormatter.Title = book.Title
	bookFormatter.Author = book.Author
	bookFormatter.Price = book.Price
	bookFormatter.Isbn = book.Isbn
	bookFormatter.Description = book.Description
	bookFormatter.Genre = book.Genre

	return bookFormatter
}

func FormatBooks(books []domainBook.Book) []BookFormatter {
	var booksFormatter []BookFormatter

	for _, user := range books {
		bookFormatter := FormatBook(user)
		booksFormatter = append(booksFormatter, bookFormatter)
	}

	return booksFormatter
}
