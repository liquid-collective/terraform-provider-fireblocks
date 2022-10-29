package sdk

import (
	fireblocksclienthttp "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client/http"
)

type Config struct {
	ABIPath    string
	Fireblocks *fireblocksclienthttp.Config
}

func (cfg *Config) SetDefault() *Config {
	if cfg.Fireblocks == nil {
		cfg.Fireblocks = &fireblocksclienthttp.Config{}
	}
	cfg.Fireblocks.SetDefault()

	return cfg
}
