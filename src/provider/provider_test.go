package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/ethereum"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

func testProviders(t *testing.T, client fireblocksclient.Client) map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"fireblocks": func() (*schema.Provider, error) {
			t.Setenv("FIREBLOCKS_API_KEY", "test-api-key")
			t.Setenv("FIREBLOCKS_RSA_PRIVATE_KEY", "test-api-secret")

			provider := Provider()
			provider.ConfigureContextFunc = func(_ context.Context, _ *schema.ResourceData) (interface{}, diag.Diagnostics) {
				abi, err := ethereum.LoadABI("testdata/ERC20.json")
				if err != nil {
					return nil, diag.FromErr(err)
				}
				return &fireblockssdk.SDK{
					Client: client,
					ABI:    abi,
				}, nil
			}
			return provider, nil
		},
	}
}
