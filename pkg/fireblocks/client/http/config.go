package clienthttp

type Config struct {
	APIKey, RSAPrivateKey, APIURL string
}

func (cfg *Config) SetDefault() *Config {
	if cfg.APIURL == "" {
		cfg.APIURL = APIURLDefault
	}

	return cfg
}
