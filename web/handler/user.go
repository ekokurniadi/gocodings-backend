package handler

import (
	"net/http"
	"web-portfolio-backend/input"
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

func (h *userHandler) New(c *gin.Context) {
	session := sessions.Default(c)
	data := session.Get("userName")
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "Create a new user"})
	c.HTML(http.StatusOK, "user_create.html", nil)
	c.HTML(http.StatusOK, "footer", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input input.InputUser

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusOK, "user_create.html", input)
		return
	}

	_, err = h.userService.UserServiceCreate(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}
	session := sessions.Default(c)
	session.Set("message", "Create User Success")
	session.Save()
	c.Redirect(http.StatusFound, "/users")

}
