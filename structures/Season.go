package structures

import (
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model

	SeasonID  string `json:"id"`
	StartDate *Date  `json:"beginDate" display:"Début"`
	EndDate   *Date  `json:"endDate" display:"Fin"`
	Tower     `json:"tower" display:"Nom"`
}

type Tower struct {
	Name string `json:"name"`
}

func (t Tower) String() string {
	return t.Name
}
