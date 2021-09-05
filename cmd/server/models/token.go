package models

import "github.com/shopspring/decimal"

type Token struct {
	ID             string          `json:"id"`
	Symbol         string          `json:"symbol"`
	Name           string          `json:"name"`
	Volume         decimal.Decimal `json:"volume"`
	VolumeUSD      *USD            `json:"volumeUSD"`
	TxCount        string          `json:"txCount"`
	PoolCount      string          `json:"poolCount"`
	WhitelistPools []struct {
		ID string `json:"id"`
	} `json:"whitelistPools"`
}
