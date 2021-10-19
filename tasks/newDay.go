package tasks

import "github.com/alexpresso/zunivers-webhooks/services"

func newDay() {
	services.DispatchEvent("new_day", nil, nil)
}
