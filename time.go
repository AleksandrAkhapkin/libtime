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
	HHMM        = "15:04"
	HHMMPostgre = "15:04:05-07:00"
) ////

var (
	ErrTimeParse = errors.New(`ErrTimeParse: should be a string formatted as "15:04-07:00"`)
	ErrTimeScan  = errors.New(`ErrTimeScan: source must be []byte`)
)

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format(HHMMPostgre) + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	if len(b) != len(HHMM)+2 {
		return ErrTimeParse
	}
	s := string(b)
	ret, err := time.Parse(HHMM, s[1:len(HHMM)+1])
	if err != nil {
		return err
	}
	t.Time = ret
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return t.Format(HHMMPostgre), nil
}

func (t *Time) Scan(src interface{}) error {
	b, ok := src.(string)
	if !ok {
		return ErrTimeScan
	}
	ret, err := time.Parse(HHMMPostgre, b)
	if err != nil {
		return err
	}
	t.Time = ret
	return nil
}
