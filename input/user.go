package input

type InputIDUser struct {
	ID int `uri:"id" binding:"required"`
}

type InputUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role" binding:"required"`
}
