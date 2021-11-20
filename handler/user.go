package handler

import (
	"net/http"
	"web-portfolio-backend/formatter"
	"web-portfolio-backend/helper"
	"web-portfolio-backend/input"
	"web-portfolio-backend/middleware"
	"web-portfolio-backend/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     service.UserService
	authService middleware.Service
}

func NewUserHandler(service service.UserService, authService middleware.Service) *userHandler {
	return &userHandler{service, authService}
}
func (h *userHandler) Create(c *gin.Context) {
	var input input.InputUser
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response :=
			helper.ApiResponse("Create user is failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.service.UserServiceCreate(input)

	if err != nil {
		response :=
			helper.ApiResponse("Create user is failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response :=
			helper.ApiResponse("Register Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatUserLogin(newUser, token)
	response :=
		helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	var input input.LoginInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response :=
			helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.service.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response :=
			helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID)

	if err != nil {
		response :=
			helper.ApiResponse("Login Account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUserLogin(loggedInUser, token)
	response := helper.ApiResponse("Login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
