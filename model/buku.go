package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Judul       string `json:"judul" form:"judul"`
	Pengarang   string `json:"pengarang" form:"pengarang"`
	Penerbit    string `json:"penerbit" form:"penerbit"`
	TahunTerbit int    `json:"tahun_terbit" form:"tahun_terbit"`
}
