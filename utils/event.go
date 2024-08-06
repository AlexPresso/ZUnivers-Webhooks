package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func EventsEnabled(eventNames []string) (res bool) {
	res = true
	for _, name := range eventNames {
		if !viper.GetBool(fmt.Sprintf("events.%s.enabled", name)) {
			res = false
		}
	}

	return
}
