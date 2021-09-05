package models

type Swap struct {
	ID        string    `json:"id"`
	Timestamp *UnixTime `json:"timestamp"`
	AmountUSD *USD      `json:"amountUSD"`
	Token1    Token     `json:"token1"`
	Token0    Token     `json:"token0"`
}
