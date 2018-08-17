package helpers

import "time"

type DateTime struct {
	*time.Time
	TimeStamp int64
	valid bool
}

func (d *DateTime) Valid() bool {
	return d.TimeStamp > 0
}