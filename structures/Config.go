package structures

import "gorm.io/gorm"

type Config struct {
	gorm.Model

	ConfigID     string `json:"id"`
	CraftValue   int    `json:"craftValue" display:"Valeur de craft"`
	IsGolden     bool   `json:"isGolden" display:"Gold"`
	Rarity       int    `json:"rarity" display:"Rareté"`
	RecycleValue int    `json:"recycleValue" display:"Valeur de recyclage"`
}
