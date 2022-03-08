package structures

import "gorm.io/gorm"

type ChallengeProgress struct {
	Challenge *Challenge `json:"challenge"`
}

type Challenge struct {
	gorm.Model

	ChallengeID    string `json:"id"`
	Description    string `json:"description" zu:"display=Description"`
	RewardLoreDust string `json:"rewardLoreDust" zu:"display=Poussière de lore"`
	Score          string `json:"score" zu:"display=Score"`
	Type           string `json:"type" zu:"display=Type"`
}
