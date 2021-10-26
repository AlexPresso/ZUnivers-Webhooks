package structures

import "gorm.io/gorm"

type Config struct {
	gorm.Model

	ConfigID     string `json:"id"`
	CraftValue   int    `json:"craftValue" zu:"display=Valeur de craft"`
	IsGolden     bool   `json:"isGolden" zu:"display=Gold"`
	Rarity       int    `json:"rarity" zu:"display=Rareté"`
	RecycleValue int    `json:"recycleValue" zu:"display=Valeur de recyclage"`
}
