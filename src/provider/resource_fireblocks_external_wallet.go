package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this resource to create a proposal
func externalWalletResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: externalWalletCreate,
		ReadContext:   externalWalletRead,
		UpdateContext: externalWalletUpdate,
		DeleteContext: externalWalletDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: `Resource is used to managed Defender Admin action proposals. 
Any actions created this way will have no approvals initially.`,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the external wallet container",
			},
			"customer_ref_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID for AML providers to associate the owner of funds with transactions",
			},
		},
	}
}

func externalWalletRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	wallet, err := sdk.GetExternalWallet(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", wallet.Name)
	d.Set("customer_ref_id", wallet.CustomerRefID)

	return nil
}

func externalWalletCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	createVaultAccountMsg := &fireblocksclient.CreateExternalWalletMsg{
		Name:          d.Get("name").(string),
		CustomerRefID: d.Get("customer_ref_id").(string),
	}

	account, err := sdk.CreateExternalWallet(ctx, createVaultAccountMsg)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(account.ID)

	return externalWalletRead(ctx, d, meta)
}

func externalWalletUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("external wallet can not be updated")
}

func externalWalletDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	err := sdk.DeleteExternalWallet(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
