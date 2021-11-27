package handler

import (
	"net/http"
	"web-portfolio-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type aboutHandler struct {
	aboutService service.AboutService
}

func NewAboutHandler(aboutService service.AboutService) *aboutHandler {
	return &aboutHandler{aboutService}
}

func (h *aboutHandler) Index(c *gin.Context) {
	abouts, err := h.aboutService.AboutServiceGetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	session := sessions.Default(c)
	data := session.Get("userName")
	flash := session.Get("message")
	session.Set("message", "")
	session.Save()
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "List Of abouts"})
	c.HTML(http.StatusOK, "list_about", gin.H{"abouts": abouts, "data": flash})
	c.HTML(http.StatusOK, "footer", nil)
}
