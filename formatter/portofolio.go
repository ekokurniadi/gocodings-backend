package formatter

import "web-portfolio-backend/schema"

type PortfolioFormatter struct {
	ID          int    `json:"id" `
	ImageCover  string `json:"image_cover"`
	Phil        string `json:"phil"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func FormatPortfolio(portofolio schema.Portofolio) PortfolioFormatter {
	portofolioFormatter := PortfolioFormatter{}
	portofolioFormatter.ID = portofolio.ID
	portofolioFormatter.Description = portofolio.Description
	portofolioFormatter.Title = portofolio.Title
	portofolioFormatter.ImageCover = portofolio.ImageCover
	portofolioFormatter.Phil = portofolio.Phil
	return portofolioFormatter
}

func FormatPortfolios(portofolios []schema.Portofolio) []PortfolioFormatter {
	portofoliosFormatter := []PortfolioFormatter{}
	for _, portofolio := range portofolios {
		portofolioFormatter := FormatPortfolio(portofolio)
		portofoliosFormatter = append(portofoliosFormatter, portofolioFormatter)
	}
	return portofoliosFormatter
}
