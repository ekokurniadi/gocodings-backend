package handler

import (
	"net/http"
	"web-portfolio-backend/formatter"
	"web-portfolio-backend/helper"
	"web-portfolio-backend/service"

	"github.com/gin-gonic/gin"
)

type portofolioHandler struct {
	service service.PortfolioService
}

func NewPortofolioHandler(service service.PortfolioService) *portofolioHandler {
	return &portofolioHandler{service}
}

func (h *portofolioHandler) GetPortofolios(c *gin.Context) {

	portofolios, err := h.service.PortofolioServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Error to get portfolios", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	response := helper.ApiResponse("List of abouts", http.StatusOK, "success", formatter.FormatPortfolios(portofolios))
	c.JSON(http.StatusOK, response)
}
