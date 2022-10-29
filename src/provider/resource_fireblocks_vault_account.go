package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this resource to create a proposal
func vaultAccountResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: vaultAccountCreate,
		ReadContext:   vaultAccountRead,
		UpdateContext: vaultAccountUpdate,
		DeleteContext: vaultAccountDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: ``,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vault account name",
			},
			"hidden_on_ui": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Should be set to true if you wish this account will not appear in the web console",
			},
			"customer_ref_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID for AML providers to associate the owner of funds with transactions",
			},
			"auto_fuel": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "In case the Gas Station service is enabled on your workspace, this flag needs to be set to true if you wish to add this account's Ethereum address to be monitored and fueled upon detected deposits of ERC20 tokens.",
			},
		},
	}
}

func vaultAccountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	account, err := sdk.GetVaultAccount(
		ctx,
		d.Id(),
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", account.Name)
	d.Set("hidden_on_ui", account.HiddenOnUI)
	d.Set("customer_ref_id", account.CustomerRefID)
	d.Set("auto_fuel", account.AutoFuel)

	return nil
}

func vaultAccountCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	createVaultAccountMsg := &fireblocksclient.CreateVaultAccountMsg{
		Name:          d.Get("name").(string),
		HiddenOnUI:    d.Get("hidden_on_ui").(bool),
		CustomerRefID: d.Get("customer_ref_id").(string),
		AutoFuel:      d.Get("auto_fuel").(bool),
	}

	account, err := sdk.CreateVaultAccount(ctx, createVaultAccountMsg)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(account.ID)

	return vaultAccountRead(ctx, d, meta)
}

func vaultAccountUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)
	if d.HasChange("name") {
		err := sdk.UpdateVaultAccount(
			ctx,
			d.Id(),
			&fireblocksclient.UpdateVaultAccountMsg{
				Name: d.Get("name").(string),
			},
		)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("hidden_on_ui") {
		var err error
		_, newHidden := d.GetChange("hidden_on_ui")
		if (newHidden).(bool) {
			err = sdk.HideVaultAccount(ctx, d.Id())
		} else {
			err = sdk.UnhideVaultAccount(ctx, d.Id())
		}

		if err != nil {
			return diag.FromErr(err)
		}
	}

	updateFields := []string{"name", "hidden_on_ui"}
	if d.HasChangesExcept(updateFields...) {
		return diag.Errorf("only fields %q can be updated", updateFields)
	}

	return nil
}

func vaultAccountDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return []diag.Diagnostic{
		{
			Severity: diag.Warning,
			Summary:  "Vault account can not be deleted through this provider. You should archive manually through fireblocks console",
			Detail:   "Vault account can not be deleted through this provider. You should archive manually through fireblocks console",
		},
	}
}
