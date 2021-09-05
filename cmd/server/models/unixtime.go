package models

import (
	"strconv"
	"time"
)

type UnixTime struct {
	time.Time
}

func (m *UnixTime) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}
	*m = UnixTime{Time: time.Unix(i, 0)}
	return nil
}
