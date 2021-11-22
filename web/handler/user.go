package handler

import (
	"net/http"
	"web-portfolio-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.UserServiceGetAll()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	session := sessions.Default(c)
	data := session.Get("userName")
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "List Of Users"})
	c.HTML(http.StatusOK, "index.html", gin.H{"users": users})
	c.HTML(http.StatusOK, "footer", nil)
}
