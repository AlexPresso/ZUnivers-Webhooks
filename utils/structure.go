package utils

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"reflect"
	"sort"
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

func GenerateDiff(a, b map[string]interface{}) (diff string, changed bool) {
	var sb strings.Builder
	changed = false

	keys := make(map[string]struct{})
	for key := range a {
		keys[key] = struct{}{}
	}
	for key := range b {
		keys[key] = struct{}{}
	}

	sortedKeys := make([]string, 0, len(keys))
	for key := range keys {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		valueA, existsA := a[key]
		valueB, existsB := b[key]

		if existsA && !existsB {
			sb.WriteString(fmt.Sprintf("- %s: %v\n", key, valueA))
			changed = true
		} else if !existsA && existsB {
			sb.WriteString(fmt.Sprintf("+ %s: %v\n", key, valueB))
			changed = true
		} else if existsA && existsB && valueA != valueB {
			sb.WriteString(fmt.Sprintf("- %s: %v\n", key, valueA))
			sb.WriteString(fmt.Sprintf("+ %s: %v\n", key, valueB))
			changed = true
		}
	}

	diff = sb.String()
	return
}
