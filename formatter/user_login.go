package formatter

import "web-portfolio-backend/schema"

type UserLoginFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	ImageUrl string `json:"image_url"`
	Token    string `json:"token"`
}

func FormatUserLogin(user schema.User, token string) UserLoginFormatter {
	formatter := UserLoginFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		ImageUrl: user.Avatar,
		Token:    token,
	}

	return formatter
}
