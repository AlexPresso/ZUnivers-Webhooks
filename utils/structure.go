package utils

import "reflect"

func AreDifferent(a interface{}, b interface{}) bool {
	a = reflect.ValueOf(a)
	bType := reflect.TypeOf(b)
	b = reflect.ValueOf(b)

	for i := 0; i < bType.NumField(); i++ {
		if _, hasTag := bType.Field(i).Tag.Lookup("display"); hasTag && a.(reflect.Value).Field(i).Interface() != b.(reflect.Value).Field(i).Interface() {
			return true
		}
	}

	return false
}
