package formatter

import "web-portfolio-backend/schema"

type AboutFormatter struct {
	ID          int    `json:"id"`
	TentangSaya string `json:"tentang_saya"`
	Alamat      string `json:"alamat"`
	Telp        string `json:"telp"`
	Whatsapp    string `json:"whatsapp"`
	Email       string `json:"email"`
}

func FormatAbout(about schema.About) AboutFormatter {
	aboutFormatter := AboutFormatter{}
	aboutFormatter.ID = about.ID
	aboutFormatter.TentangSaya = about.TentangSaya
	aboutFormatter.Alamat = about.Alamat
	aboutFormatter.Telp = about.Telp
	aboutFormatter.Whatsapp = about.Whatsapp
	aboutFormatter.Email = about.Email
	return aboutFormatter
}

func FormatAbouts(abouts []schema.About) []AboutFormatter {
	aboutsFormatter := []AboutFormatter{}
	for _, about := range abouts {
		aboutFormatter := FormatAbout(about)
		aboutsFormatter = append(aboutsFormatter, aboutFormatter)
	}
	return aboutsFormatter
}
