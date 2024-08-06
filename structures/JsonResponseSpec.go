package structures

import "gorm.io/gorm"

type JsonResponseSpec struct {
	gorm.Model

	EventName     string
	PreviousValue string
}
