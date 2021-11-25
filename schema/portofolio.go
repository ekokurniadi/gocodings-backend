package schema

import "time"

type Portofolio struct {
	ID          int
	ImageCover  string
	Phil        string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
