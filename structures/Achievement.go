package structures

import "gorm.io/gorm"

type AchievementCategory struct {
	gorm.Model

	CategoryID string `json:"id"`
	Name       string `json:"name" zu:"display=Nom"`
}

type AchievementProgress struct {
	Achievement *Achievement `json:"achievement"`
}

type Achievement struct {
	gorm.Model

	Name          string `json:"name" zu:"display=Nom"`
	Description   string `json:"description" zu:"display=Description"`
	AchievementID string `json:"id" zu:"url=/achievement/%s"`
}
