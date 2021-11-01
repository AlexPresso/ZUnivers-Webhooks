package utils

import (
	"reflect"
	"strings"
)

func AreDifferent(a, b interface{}) bool {
	a = reflect.ValueOf(a)
	bType := reflect.TypeOf(b)
	b = reflect.ValueOf(b)

	for i := 0; i < bType.NumField(); i++ {
		if val, hasTag := bType.Field(i).Tag.Lookup("zu"); hasTag && strings.Contains(val, "display") {
			aVal := a.(reflect.Value).Field(i).Interface()
			bVal := b.(reflect.Value).Field(i).Interface()

			//TODO: check for time castable and time.Time.Equal

			return aVal != bVal
		}
	}

	return false
}
