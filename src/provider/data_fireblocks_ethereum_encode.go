package provider

import (
	"context"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this data source to obtain various information parsed from an existing node key in hex format.
//
// Node key encodes a private key that defines an identity of a Quorum node in the network. It is primarily used in P2P networking.
func ethereumEncodeDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: ethereumEncodeDataSourceRead,
		Schema: map[string]*schema.Schema{
			"method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the method to call",
			},
			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Arguments to pass to the method in JSON format",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"result": {
				Type:        schema.TypeString,
				Description: "Hex encoded result",
				Computed:    true,
			},
		},
	}
}

func ethereumEncodeDataSourceRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	var args [][]byte
	for _, rawArg := range d.Get("args").([]interface{}) {
		args = append(args, []byte(rawArg.(string)))
	}

	res, err := sdk.ABI.PackJSON(
		d.Get("method").(string),
		args...,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	u, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(u)
	d.Set("result", res)

	return nil
}
