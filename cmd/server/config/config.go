package config

import (
	"flag"
	"fmt"
)

// Cfg holds common program configuration
type Cfg struct {
	UniswapAPI string
	ApiPort    string
}

// New returns instance of Cfg
func New() Cfg {
	cfg := Cfg{}

	flag.StringVar(
		&cfg.UniswapAPI,
		"uniswap-api",
		"https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-alt",
		"uniswap API endpoint",
	)

	flag.StringVar(&cfg.ApiPort, "port", "9090", "REST server API PORT")
	flag.Parse()
	return cfg
}

func (c Cfg) GetApiPort() string {
	return fmt.Sprintf(":%s", c.ApiPort)
}
