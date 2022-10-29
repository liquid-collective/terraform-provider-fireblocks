package clienthttp

import (
	kilncmdutils "github.com/kilnfi/go-utils/cmd/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Flags register viper compatible pflags for auth
func Flags(v *viper.Viper, f *pflag.FlagSet) {
	APIKeyFlag(v, f)
	RSAPrivateKeyFlag(v, f)
	APIURLFlag(v, f)
}

const (
	apiKeyFlag     = "fireblocks-api-key"
	apiKeyViperKey = "fireblocks.apiKey"
	APIKeyEnv      = "FIREBLOCKS_API_KEY"
)

// APIKeyFlag register flag for Authentication apiKey
func APIKeyFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Fireblocks API Key",
		APIKeyEnv,
	)

	f.String(apiKeyFlag, "", desc)
	_ = v.BindPFlag(apiKeyViperKey, f.Lookup(apiKeyFlag))
	_ = v.BindEnv(apiKeyViperKey, APIKeyEnv)
}

func GetAPIKey(v *viper.Viper) string {
	return v.GetString(apiKeyViperKey)
}

const (
	rsaPrivateKeyFlag     = "fireblocks-api-secret"
	rsaPrivateKeyViperKey = "fireblocks.rsaPrivateKey"
	RSAPrivateKeyEnv      = "FIREBLOCKS_RSA_PRIVATE_KEY"
)

// RSAPrivateKeyFlag register flag for Authentication rsaPrivateKey
func RSAPrivateKeyFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Fireblocks API RSA Private Key",
		RSAPrivateKeyEnv,
	)

	f.String(rsaPrivateKeyFlag, "", desc)
	_ = v.BindPFlag(rsaPrivateKeyViperKey, f.Lookup(rsaPrivateKeyFlag))
	_ = v.BindEnv(rsaPrivateKeyViperKey, RSAPrivateKeyEnv)
}

func GetRSAPrivateKey(v *viper.Viper) string {
	return v.GetString(rsaPrivateKeyViperKey)
}

const (
	apiURLFlag     = "fireblocks-api-url"
	apiURLViperKey = "fireblocks.apiURL"
	APIURLEnv      = "FIREBLOCKS_API_URL"
	APIURLDefault  = "https://api.fireblocks.io"
)

// APIURLFlag register flag for Authentication apiURL
func APIURLFlag(v *viper.Viper, f *pflag.FlagSet) {
	desc := kilncmdutils.FlagDesc(
		"Fireblocks API URL",
		APIURLEnv,
	)

	f.String(apiURLFlag, "", desc)
	_ = v.BindPFlag(apiURLViperKey, f.Lookup(apiURLFlag))
	_ = v.BindEnv(apiURLViperKey, APIURLEnv)
}

func GetAPIURL(v *viper.Viper) string {
	return v.GetString(apiURLViperKey)
}

func NewConfigFromViper(v *viper.Viper) *Config {
	return &Config{
		APIKey:        GetAPIKey(v),
		RSAPrivateKey: GetRSAPrivateKey(v),
		APIURL:        GetAPIURL(v),
	}
}
