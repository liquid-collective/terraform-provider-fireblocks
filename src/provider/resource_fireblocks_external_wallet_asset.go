package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this resource to create a proposal
func externalWalletAssetResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: externalWalletAssetCreate,
		ReadContext:   externalWalletAssetRead,
		UpdateContext: externalWalletAssetUpdate,
		DeleteContext: externalWalletAssetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: ``,
		Schema: map[string]*schema.Schema{
			"external_wallet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the external wallet the asset is attached to",
			},
			"asset_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Asset Id",
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address of the contract wallet",
			},
			"tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination tag (for XRP, used as memo for EOS/XLM) of the contract wallet.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External wallet asset status",
			},
			"activation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time the external wallet ",
			},
		},
	}
}

func externalWalletAssetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	parts := strings.Split(d.Id(), ":")
	asset, err := sdk.GetExternalWalletAsset(ctx, parts[0], parts[1])
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("external_wallet_id", parts[0])
	d.Set("asset_id", parts[1])
	d.Set("address", asset.Address)
	d.Set("tag", asset.Tag)
	d.Set("status", string(asset.Status))
	d.Set("activation_time", asset.ActivationTime)

	return nil
}

func externalWalletAssetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	addExternalWalletAssetMsg := &fireblocksclient.AddExternalWalletAssetMsg{
		Address: d.Get("address").(string),
		Tag:     d.Get("tag").(string),
	}

	_, err := sdk.AddExternalWalletAsset(
		ctx,
		d.Get("external_wallet_id").(string),
		d.Get("asset_id").(string),
		addExternalWalletAssetMsg,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v:%v", d.Get("external_wallet_id").(string), d.Get("asset_id").(string)))

	return externalWalletAssetRead(ctx, d, meta)
}

func externalWalletAssetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("external wallet asset can not be updated")
}

func externalWalletAssetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	parts := strings.Split(d.Id(), ":")
	err := sdk.DeleteExternalWalletAsset(ctx, parts[0], parts[1])
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
