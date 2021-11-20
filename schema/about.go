package schema

import "time"

type About struct {
	ID          int
	TentangSaya string
	Alamat      string
	Telp        string
	Whatsapp    string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
