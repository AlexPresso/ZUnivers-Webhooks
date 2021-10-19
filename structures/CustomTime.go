package structures

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	*ct = CustomTime(t)

	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct)
}

func (ct CustomTime) Format(s string) string {
	t := time.Time(ct)
	return t.Format(s)
}

func (ct *CustomTime) Scan(v interface{}) error {
	*ct = CustomTime(v.(time.Time))
	return nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	return time.Time(ct), nil
}
