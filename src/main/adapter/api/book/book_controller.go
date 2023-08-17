package book

import (
	"net/http"
	serviceBook "service-api/src/main/app/book/service"
	bookFormatter "service-api/src/main/util/formatter/book"
	bookInput "service-api/src/main/util/input/book"
	"strconv"

	infra "service-api/src/main/infra"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService serviceBook.BookService
}

func NewBookController(bookService serviceBook.BookService) *BookController {
	return &BookController{bookService}
}

// insert book
func (controller *BookController) InsertBook(context *gin.Context) {
	var input bookInput.BookInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", 422, nil)
		context.JSON(
			422,
			response,
		)

		return
	}

	newBook, err := controller.bookService.InsertBook(input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", 400, nil)
		context.JSON(
			400,
			response,
		)

		return
	}

	formatter := bookFormatter.FormatBook(newBook)
	response := infra.NewResponseAPI("success", "success", http.StatusCreated, formatter)
	context.JSON(
		http.StatusCreated,
		response,
	)
}

// get all book
func (controller *BookController) GetAllBook(context *gin.Context) {
	books, err := controller.bookService.GetAllBook()
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", 400, nil)
		context.JSON(
			400,
			response,
		)

		return
	}

	formatter := bookFormatter.FormatBooks(books)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}

// get book by id
func (controller *BookController) GetBookByID(context *gin.Context) {
	bookID := context.Param("book_id")
	bookIDConvert, err := strconv.ParseUint(bookID, 10, 64)
	bookIDUint64 := uint64(bookIDConvert)
	book, err := controller.bookService.GetBookByID(uint(bookIDUint64))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", 400, nil)
		context.JSON(
			400,
			response,
		)

		return
	}

	formatter := bookFormatter.FormatBook(book)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}

// update book
func (controller *BookController) UpdateBook(context *gin.Context) {
	var input bookInput.BookInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusUnprocessableEntity, nil)
		context.JSON(
			422,
			response,
		)

		return
	}

	bookID := context.Param("book_id")
	bookIDConvert, err := strconv.ParseUint(bookID, 10, 64)
	bookIDUint64 := uint64(bookIDConvert)
	book, err := controller.bookService.UpdateBookByID(uint(bookIDUint64), input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", 400, nil)
		context.JSON(
			400,
			response,
		)

		return
	}

	formatter := bookFormatter.FormatBook(book)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}

// delete book by id
func (controller *BookController) DeleteBookByID(context *gin.Context) {
	bookID := context.Param("book_id")
	bookIDConvert, err := strconv.ParseUint(bookID, 10, 64)
	bookIDUint64 := uint64(bookIDConvert)
	book, err := controller.bookService.DeleteBookByID(uint(bookIDUint64))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	formatter := bookFormatter.FormatBook(book)
	response := infra.NewResponseAPI("success", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}
