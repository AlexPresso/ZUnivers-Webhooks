package structures

import "gorm.io/gorm"

type Event struct {
	gorm.Model

	EventID     string    `json:"id"`
	BeginDate   *DateTime `json:"beginDate" zu:"display=Début"`
	EndDate     *DateTime `json:"endDate" zu:"display=Fin"`
	ImageURL    string    `json:"imageUrl" zu:"imageUrl=%s"`
	Name        string    `json:"name" zu:"display=Nom"`
	BalanceCost int       `json:"balanceCost" zu:"display=Coût d'invocation"`
}
