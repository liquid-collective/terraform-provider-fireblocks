package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testDataSourceTransfer = `
data "fireblocks_ethereum_encode" "transfer" {
	method = "transfer"
	args = [jsonencode("0xc0ffee254729296a45a3885639AC7E10F9d54979"), 100000000000000000]
}
`

func TestTxPayloadDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: testProviders(t, nil),
		Steps: []resource.TestStep{
			{
				Config: testDataSourceTransfer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.fireblocks_ethereum_encode.transfer", "result", "0xa9059cbb000000000000000000000000c0ffee254729296a45a3885639ac7e10f9d54979000000000000000000000000000000000000000000000000016345785d8a0000"),
				),
			},
		},
	})
}
