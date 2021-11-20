package schema

import "time"

type User struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Avatar    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}