package handler

import (
	"fmt"
	"net/http"
	"web-portfolio-backend/input"
	"web-portfolio-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type portfolioHandler struct {
	portfolioService service.PortfolioService
}

func NewPortfolioHandler(portfolioService service.PortfolioService) *portfolioHandler {
	return &portfolioHandler{portfolioService}
}

func (h *portfolioHandler) Index(c *gin.Context) {
	portfolios, err := h.portfolioService.PortofolioServiceGetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	session := sessions.Default(c)
	data := session.Get("userName")
	flash := session.Get("message")
	session.Set("message", "")
	session.Save()
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "List Of portfolios"})
	c.HTML(http.StatusOK, "list_portfolios", gin.H{"portfolios": portfolios, "data": flash})
	c.HTML(http.StatusOK, "footer", nil)
}

func (h *portfolioHandler) New(c *gin.Context) {
	session := sessions.Default(c)
	data := session.Get("userName")
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "Create a new portofolio"})
	c.HTML(http.StatusOK, "portfolio_create", gin.H{"error": ""})
	c.HTML(http.StatusOK, "footer", nil)
}

func (h *portfolioHandler) Create(c *gin.Context) {
	var input input.InputPortfolio
	file, err := c.FormFile("image_cover")
	session := sessions.Default(c)
	data := session.Get("userName")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "header", gin.H{"nama": data, "title": "Create a new portofolio"})
		c.HTML(http.StatusInternalServerError, "portfolio_create", input)
		c.HTML(http.StatusInternalServerError, "footer", nil)
		return
	}

	input.Title = c.PostForm("title")
	input.Description = c.PostForm("description")
	input.Phil = c.PostForm("phil")
	images := ""
	if file != nil {
		path := fmt.Sprintf("images/%s", file.Filename)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "header", gin.H{"nama": data, "title": "Create a new portofolio"})
			c.HTML(http.StatusInternalServerError, "portfolio_create", input)
			c.HTML(http.StatusInternalServerError, "footer", nil)
			return
		}
		images = file.Filename
	}
	input.ImageCover = images
	_, err = h.portfolioService.PortofolioServiceCreate(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "header", gin.H{"nama": data, "title": "Create a new portofolio"})
		c.HTML(http.StatusInternalServerError, "portfolio_create", input)
		c.HTML(http.StatusInternalServerError, "footer", nil)
		return
	}
	session.Set("message", "Create Portfolio Success")
	session.Save()
	c.Redirect(http.StatusFound, "/portfolios")

}
