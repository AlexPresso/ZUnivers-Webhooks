package utils

import (
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/sergi/go-diff/diffmatchpatch"
	"reflect"
	"regexp"
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

func GenerateDiff(a, b string) (diff string, hasDiff bool) {
	dmp := diffmatchpatch.New()

	ca, cb, lines := dmp.DiffLinesToChars(a, b)
	diffs := dmp.DiffMain(ca, cb, false)
	diffs = dmp.DiffCharsToLines(diffs, lines)
	diffs = dmp.DiffCleanupSemantic(diffs)

	spacesRgx := regexp.MustCompile(`^\s+`)

	var finalDiff strings.Builder
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffEqual:
			finalDiff.WriteString(diff.Text)
		case diffmatchpatch.DiffDelete:
			hasDiff = true
			lineDiffFormatAppender("- %s\n", diff.Text, spacesRgx, &finalDiff)
		case diffmatchpatch.DiffInsert:
			hasDiff = true
			lineDiffFormatAppender("+ %s\n", diff.Text, spacesRgx, &finalDiff)
		}
	}

	return finalDiff.String(), hasDiff
}

func lineDiffFormatAppender(format string, fullDiffText string, offsetRegex *regexp.Regexp, sb *strings.Builder) {
	for _, line := range strings.Split(fullDiffText, "\n") {
		offset := len(offsetRegex.FindAllStringIndex(line, -1))
		sb.WriteString(fmt.Sprintf(format, line[offset:]))
	}
}
