package formatter

import "web-portfolio-backend/schema"

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}

func FormatUser(user schema.User) UserFormatter {
	userFormatter := UserFormatter{}
	userFormatter.ID = user.ID
	userFormatter.Name = user.Name
	userFormatter.Username = user.Username
	userFormatter.Password = user.Password
	userFormatter.Avatar = user.Avatar
	userFormatter.Role = user.Role
	return userFormatter
}

func FormatUsers(users []schema.User) []UserFormatter {
	userFormatters := []UserFormatter{}
	for _, user := range users {
		userFormatter := FormatUser(user)
		userFormatters = append(userFormatters, userFormatter)
	}
	return userFormatters
}
