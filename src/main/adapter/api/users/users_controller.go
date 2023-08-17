package users

import (
	"net/http"
	serviceAuthJWT "service-api/src/main/adapter/auth"
	serviceUser "service-api/src/main/app/users/service"
	infra "service-api/src/main/infra"
	formatUser "service-api/src/main/util/formatter/users"
	inputUser "service-api/src/main/util/input/users"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService serviceUser.UserService
	jwtService  serviceAuthJWT.Service
}

func NewUserController(service serviceUser.UserService, jwtService serviceAuthJWT.Service) *userController {
	return &userController{
		userService: service,
		jwtService:  jwtService,
	}
}

func (controller *userController) RegisterUser(context *gin.Context) {
	var input inputUser.RegisterUserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusUnprocessableEntity, nil)
		context.JSON(
			http.StatusUnprocessableEntity,
			response,
		)

		return
	}

	newUser, err := controller.userService.RegisterUser(input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	token, err := controller.jwtService.GenerateToken(int(newUser.ID))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	formatter := formatUser.FormatUser(newUser, token)
	response := infra.NewResponseAPI("success login user", "success", http.StatusCreated, formatter)
	context.JSON(
		http.StatusCreated,
		response,
	)
}

func (controller *userController) LoginUser(context *gin.Context) {
	var input inputUser.LoginUserInput

	err := context.ShouldBindJSON(&input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusUnprocessableEntity, nil)
		context.JSON(
			http.StatusUnprocessableEntity,
			response,
		)

		return
	}

	// proses login
	userLogin, err := controller.userService.LoginUser(input)
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	// generate token
	token, err := controller.jwtService.GenerateToken(int(userLogin.ID))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	formatter := formatUser.FormatUser(userLogin, token)
	response := infra.NewResponseAPI("success login user", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}

func (controller *userController) GetAllUser(context *gin.Context) {
	users, err := controller.userService.FindAllUser()
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	formatter := formatUser.FormatUsers(users)
	response := infra.NewResponseAPI("success get all user", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}

func (controller *userController) GetUserByID(context *gin.Context) {
	id := context.Param("id")
	token := context.GetHeader("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	idConvert, err := strconv.ParseUint(id, 10, 64)
	idUint32 := uint32(idConvert)

	user, err := controller.userService.GetUserByID(uint(idUint32))
	if err != nil {
		errMsg := err.Error()
		response := infra.NewResponseAPI(errMsg, "failed", http.StatusBadRequest, nil)
		context.JSON(
			http.StatusBadRequest,
			response,
		)

		return
	}

	formatter := formatUser.FormatUser(user, token)
	response := infra.NewResponseAPI("success get user by id", "success", http.StatusOK, formatter)
	context.JSON(
		http.StatusOK,
		response,
	)
}
