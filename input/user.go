package input

type InputIDUser struct {
	ID int `uri:"id" binding:"required"`
}

type InputUser struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role" form:"role" binding:"required"`
}
