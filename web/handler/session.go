package handler

import (
	"net/http"
	"web-portfolio-backend/input"
	"web-portfolio-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService service.UserService
}

func NewSessionHandler(userService service.UserService) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	session := sessions.Default(c)
	data := session.Get("message")
	session.Set("message", "")
	session.Save()
	c.HTML(http.StatusOK, "auth.html", gin.H{"data": data})
}

func (h *sessionHandler) LoginAction(c *gin.Context) {
	var input input.LoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		session := sessions.Default(c)
		session.Set("message", "User not Found")
		session.Save()
		c.Redirect(http.StatusFound, "/")
		return
	}

	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		session := sessions.Default(c)
		session.Set("message", "User not Found")
		session.Save()
		c.Redirect(http.StatusFound, "/")
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *sessionHandler) Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	data := session.Get("userName")
	c.HTML(http.StatusOK, "header", gin.H{"nama": data, "title": "Dashboard"})
	c.HTML(http.StatusOK, "dashboard.html", nil)
	c.HTML(http.StatusOK, "footer", nil)
}
func (h *sessionHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
