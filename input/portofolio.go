package input

type InputPortfolioID struct {
	ID int `uri:"id" binding:"required"`
}
type InputPortfolio struct {
	ID          int    `json:"id" form:"id"`
	ImageCover  string `json:"image_cover" form:"image_cover"`
	Phil        string `json:"phil" form:"phil" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}
