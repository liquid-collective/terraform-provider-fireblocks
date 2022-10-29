package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this resource to create a proposal
func vaultAccountAssetResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: vaultAccountAssetCreate,
		ReadContext:   vaultAccountAssetRead,
		UpdateContext: vaultAccountAssetUpdate,
		DeleteContext: vaultAccountAssetDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: ``,
		Schema: map[string]*schema.Schema{
			"vault_account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the external wallet the asset is attached to",
			},
			"asset_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Asset ID",
			},
			"total": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total wallet balance. Values are returned according to balance decimal precision",
			},
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Address for this asset",
			},
		},
	}
}

func vaultAccountAssetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	parts := strings.Split(d.Id(), ":")
	asset, err := sdk.GetVaultAccountAssetBalance(ctx, parts[0], parts[1])
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("vault_account_id", parts[0])
	d.Set("asset_id", asset.ID)
	d.Set("total", asset.Total)

	addresses, err := sdk.ListVaultAccountAssetAddresses(ctx, parts[0], parts[1])
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("address", addresses[0].Address)

	return nil
}

func vaultAccountAssetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	_, err := sdk.CreateVaultAccountAsset(
		ctx,
		d.Get("vault_account_id").(string),
		d.Get("asset_id").(string),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v:%v", d.Get("vault_account_id").(string), d.Get("asset_id").(string)))

	return vaultAccountAssetRead(ctx, d, meta)
}

func vaultAccountAssetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("vault account asset can not be updated")
}

func vaultAccountAssetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return []diag.Diagnostic{
		{
			Severity: diag.Warning,
			Summary:  "Vault account asset can not be deleted through this provider. You should archive manually through fireblocks console",
			Detail:   "Vault account asset can not be deleted through this provider. You should archive manually through fireblocks console",
		},
	}
}
