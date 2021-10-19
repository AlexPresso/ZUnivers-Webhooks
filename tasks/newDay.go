package tasks

import "zunivers-webhooks/services"

func newDay() {
	services.DispatchEvent("new_day", nil, nil)
}
