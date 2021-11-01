package structures

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type DateTime time.Time

func (ct *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	*ct = DateTime(t)

	return nil
}

func (ct DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct)
}

func (ct DateTime) Format(s string) string {
	t := time.Time(ct)
	return t.Format(s)
}

func (ct *DateTime) Scan(v interface{}) error {
	*ct = DateTime(v.(time.Time))
	return nil
}

func (ct DateTime) Value() (driver.Value, error) {
	return time.Time(ct), nil
}

func (ct DateTime) String() string {
	return ct.Format("02.01.2006 à 15:04:05")
}

func (ct DateTime) ToTime() time.Time {
	return time.Time(ct)
}
