package structures

import (
	"gorm.io/gorm"
)

type Season struct {
	gorm.Model

	SeasonID  string `json:"id"`
	StartDate *Date  `json:"beginDate" zu:"display=Début"`
	EndDate   *Date  `json:"endDate" zu:"display=Fin"`
	Tower     `json:"tower" zu:"display=Nom"`
}

type Tower struct {
	Name string `json:"name"`
}

func (t Tower) String() string {
	return t.Name
}
