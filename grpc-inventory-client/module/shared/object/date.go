package object

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const dateLayout = "2006-01-02"

//Date used as type for date data
type Date struct {
	time.Time
}

//Value implements the driver Valuer interface
func (t Date) Value() (driver.Value, error) {
	if t.Time.Equal(time.Time{}) {
		return nil, nil
	}
	return t.Time, nil
}

//Scan implements the Scanner interface
func (t *Date) Scan(value interface{}) error {
	var err error

	if value == nil {
		t.Time = time.Time{}
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		t.Time, err = time.Parse(dateLayout, string(v))
		return err
	case string:
		t.Time, err = time.Parse(dateLayout, v)
		return err
	}

	return nil
}

//MarshalJSON correctly serializes a date to JSON
func (t Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(dateLayout))), nil
}

//UnmarshalJSON correctly deserializes a date from JSON
func (t *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	return t.Scan(s)
}
