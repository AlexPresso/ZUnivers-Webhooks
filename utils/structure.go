package utils

import (
	"github.com/alexpresso/zunivers-webhooks/structures"
	"reflect"
	"strings"
	"time"
)

func AreDifferent(a, b interface{}) bool {
	a = reflect.ValueOf(a)
	bType := reflect.TypeOf(b)
	b = reflect.ValueOf(b)

	for i := 0; i < bType.NumField(); i++ {
		if val, hasTag := bType.Field(i).Tag.Lookup("zu"); hasTag && strings.Contains(val, "display") {
			aVal := a.(reflect.Value).Field(i).Interface()
			bVal := b.(reflect.Value).Field(i).Interface()

			if IsTime(aVal) {
				if TimeDifference(aVal, bVal) {
					return true
				}
			} else if aVal != bVal {
				return true
			}
		}
	}

	return false
}

func IsTime(v interface{}) bool {
	_, a := v.(*structures.DateTime)
	_, b := v.(*structures.Date)

	return a || b
}

func TimeDifference(a interface{}, b interface{}) bool {
	var aT time.Time
	var bT time.Time

	if _, ok := a.(*structures.DateTime); ok {
		aT = a.(*structures.DateTime).ToTime()
		bT = b.(*structures.DateTime).ToTime()
	}

	if _, ok := a.(*structures.Date); ok {
		aT = a.(*structures.Date).ToTime()
		bT = b.(*structures.Date).ToTime()
	}

	return !aT.Equal(bT)
}
