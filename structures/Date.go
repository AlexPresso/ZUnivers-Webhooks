package structures

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type Date time.Time

func (ct *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*ct = Date(t)

	return nil
}

func (ct Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct)
}

func (ct Date) Format(s string) string {
	t := time.Time(ct)
	return t.Format(s)
}

func (ct *Date) Scan(v interface{}) error {
	*ct = Date(v.(time.Time))
	return nil
}

func (ct Date) Value() (driver.Value, error) {
	return time.Time(ct), nil
}

func (ct Date) String() string {
	return ct.Format("02.01.2006")
}
