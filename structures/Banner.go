package structures

import "gorm.io/gorm"

type Banner struct {
	gorm.Model

	BannerID string `json:"id" imageUrl:"/document/%s"`
	Title    string `json:"title" display:"Titre"`
	Name     string `json:"name" display:"Nom"`
}
