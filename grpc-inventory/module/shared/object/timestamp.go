package object

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const timestampLayout = "2006-01-02 15:04:05"

//Timestamp used as type for timestamp data
type Timestamp struct {
	time.Time
}

//Value implements the driver Valuer interface
func (t Timestamp) Value() (driver.Value, error) {
	if t.Time.Equal(time.Time{}) {
		return nil, nil
	}
	return t.Time, nil
}

//Scan implements the Scanner interface
func (t *Timestamp) Scan(value interface{}) error {
	var err error

	if value == nil {
		t.Time = time.Time{}
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		t.Time, err = time.Parse(timestampLayout, string(v))
		return err
	case string:
		t.Time, err = time.Parse(timestampLayout, v)
		return err
	}

	return nil
}

//MarshalJSON correctly serializes a timestamp to JSON
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(timestampLayout))), nil
}

//UnmarshalJSON correctly deserializes a timestamp from JSON
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	return t.Scan(s)
}
