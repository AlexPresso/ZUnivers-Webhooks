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
		if val, hasTag := bType.Field(i).Tag.Lookup("zu"); hasTag && strings.Contains(val, "display") && a.(reflect.Value).Field(i).Interface() != b.(reflect.Value).Field(i).Interface() {
			return true
		}
	}

	return false
}
