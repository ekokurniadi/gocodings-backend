package input

type InputID struct {
	ID int `uri:"id" binding:"required"`
}

type InputAbout struct {
	TentangSaya string `json:"tentang_saya" binding:"required"`
	Alamat      string `json:"alamat" binding:"required"`
	Telp        string `json:"telp" binding:"required"`
	Whatsapp    string `json:"whatsapp" binding:"required"`
	Email       string `json:"email" binding:"required"`
}
