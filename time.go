package libtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Time struct {
	time.Time
}

const (
	ISO8601 = "15:04-07:00"
	ISO8601Seconds = "15:04:05-07:00"
)

var (
	ErrTimeParse = errors.New(`ErrTimeParse: should be a string formatted as "15:04-07:00"`)
	ErrTimeScan  = errors.New(`ErrTimeScan: source must be []byte`)
)

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(ISO8601) + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) != len(ISO8601)+2 {
		return ErrTimeParse
	}
	s := string(b)
	ret, err := time.Parse(ISO8601, s[1:len(ISO8601)+1])
	if err != nil {
		return err
	}
	t.Time = ret
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Format(ISO8601), nil
}

func (t *Time) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return ErrTimeScan
	}
	ret, err := time.Parse(ISO8601Seconds, b)
	if err != nil {
		return err
	}
	t.Time = ret
	return nil
}
