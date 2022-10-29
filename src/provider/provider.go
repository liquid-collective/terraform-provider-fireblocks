package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client/http"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(fireblocksclient.APIKeyEnv, nil),
			},
			"rsa_private_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(fireblocksclient.RSAPrivateKeyEnv, nil),
			},
			"api_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(fireblocksclient.APIURLEnv, fireblocksclient.APIURLDefault),
			},
			"abi_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fireblocks_vault_account":         vaultAccountResource(),
			"fireblocks_vault_account_asset":   vaultAccountAssetResource(),
			"fireblocks_external_wallet":       externalWalletResource(),
			"fireblocks_external_wallet_asset": externalWalletAssetResource(),
			"fireblocks_transaction":           transactionResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fireblocks_ethereum_encode": ethereumEncodeDataSource(),
		},
	}

	provider.ConfigureContextFunc = ConfigureProvider()

	return provider
}

func ConfigureProvider() schema.ConfigureContextFunc {
	return func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		cfg := &fireblockssdk.Config{
			Fireblocks: &fireblocksclient.Config{
				APIURL:        data.Get("api_url").(string),
				APIKey:        data.Get("api_key").(string),
				RSAPrivateKey: data.Get("rsa_private_key").(string),
			},
			ABIPath: data.Get("abi_path").(string),
		}

		sdk, err := fireblockssdk.New(cfg.SetDefault())
		if err != nil {
			return nil, diag.FromErr(err)
		}

		err = sdk.Init(ctx)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return sdk, nil
	}
}
