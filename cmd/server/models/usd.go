package models

import (
	"github.com/shopspring/decimal"
)

type USD string

func (u *USD) UnmarshalJSON(b []byte) error {
	dec, err := decimal.NewFromString(string(b[1 : len(b)-1]))
	if err != nil {
		return err
	}

	*u = USD(dec.Round(2).String())
	return nil
}
