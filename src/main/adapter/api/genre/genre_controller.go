package genre

import (
	serviceGenre "service-api/src/main/app/genre/service"
	"service-api/src/main/infra"
	genreInput "service-api/src/main/util/input/genre"
	"strconv"

	"github.com/gin-gonic/gin"
)

type genreController struct {
	genreService serviceGenre.GenreService
}

func NewGenreController(genreService serviceGenre.GenreService) *genreController {
	return &genreController{genreService}
}

// genre all
func (controller *genreController) GetAllGenre(context *gin.Context) {
	genres, err := controller.genreService.GetAllGenre()
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	if len(genres) == 0 {
		errMsg := "Genre not found"
		response := infra.NewResponseAPI(errMsg, "Error", 404, nil)
		context.JSON(
			404,
			response,
		)
		return
	}

	response := infra.NewResponseAPI("Success", "Success", 200, genres)
	context.JSON(
		200,
		response,
	)

}

// genre by id
func (controller *genreController) GetGenreByID(context *gin.Context) {
	id := context.Param("id")
	idConvert, err := strconv.Atoi(id)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	genre, err := controller.genreService.GetGenreByID(uint(idConvert))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	response := infra.NewResponseAPI("Success", "Success", 200, genre)
	context.JSON(
		200,
		response,
	)
}

// create genre
func (controller *genreController) CreateGenre(context *gin.Context) {
	var input genreInput.GenreInput
	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 422, nil)
		context.JSON(
			422,
			response,
		)
		return
	}

	genre, err := controller.genreService.InsertGenre(input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	response := infra.NewResponseAPI("Success", "Success", 201, genre)
	context.JSON(
		201,
		response,
	)
}

// update genre
func (controller *genreController) UpdateGenre(context *gin.Context) {
	id := context.Param("id")
	idConvert, err := strconv.Atoi(id)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	var input genreInput.GenreInput
	err = context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 422, nil)
		context.JSON(
			422,
			response,
		)
		return
	}

	genre, err := controller.genreService.UpdateGenre(uint(idConvert), input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	response := infra.NewResponseAPI("Success", "Success", 200, genre)
	context.JSON(
		200,
		response,
	)
}

// delete genre
func (controller *genreController) DeleteGenre(context *gin.Context) {
	id := context.Param("id")
	idConvert, err := strconv.Atoi(id)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	err = controller.genreService.DeleteGenreByID(uint(idConvert))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "Error", 400, nil)
		context.JSON(
			400,
			response,
		)
		return
	}

	response := infra.NewResponseAPI("Success", "Success", 200, nil)
	context.JSON(
		200,
		response,
	)
}
