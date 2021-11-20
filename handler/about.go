package handler

import (
	"net/http"
	"web-portfolio-backend/helper"
	"web-portfolio-backend/input"
	"web-portfolio-backend/service"

	"github.com/gin-gonic/gin"
)

type aboutHandler struct {
	service service.AboutService
}

func NewAboutHandler(service service.AboutService) *aboutHandler {
	return &aboutHandler{service}
}

func (h *aboutHandler) GetAbout(c *gin.Context) {
	var input input.InputID
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	aboutDetail, err := h.service.AboutServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Berhasil mendapatkan data", http.StatusOK, "success", aboutDetail)
	c.JSON(http.StatusOK, response)
}
