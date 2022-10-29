package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fireblocksclient "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
	fireblockssdk "github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/sdk"
)

// Use this resource to create a proposal
func transactionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: transactionCreate,
		ReadContext:   transactionRead,
		UpdateContext: transactionUpdate,
		DeleteContext: transactionDelete,
		Description:   `Resource is used to manage Fireblocks transactions.`,
		Schema: map[string]*schema.Schema{
			"asset_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vault account name",
			},
			"source_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source of the transaction asset ID",
			},
			"source_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source of the transaction type",
			},
			"destination_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination of the transaction asset ID",
			},
			"destination_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination of the transaction asset type",
			},
			"destination_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_one_time_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Destination of the transaction asset type",
			},
			"amount": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If the transfer is a withdrawal from an exchange, the actual amount that was requested to be transferred. Otherwise, the requested amount",
			},
			"gas_price": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "For ETH-based assets only this will be used instead of the fee property, value is in Gwei",
			},
			"gas_limit": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "For ETH-based assets only",
			},
			"priority_fee": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "For ETH-based assets only",
			},
			"operation": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "TRANSFER",
				Description: "Type of transaction operation",
			},
			"customer_ref_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID for AML providers to associate the owner of funds with transactions",
			},
			"extra_parameters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Source of the transaction",
				Elem:        schema.TypeString,
			},
			"note": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customer note of the transaction",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the transaction",
			},
			"sub_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "More detailed status of the transaction",
			},
			"tx_hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Blockchain hash of the transaction",
			},
			"source_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "For account based assets only, the source address of the transaction",
			},
			"destination_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Address where the asset were transferred",
			},
			"block_height": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Block number the transaction was included in",
			},
			"block_hash": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Block hash the transaction was included in",
			},
			"signed_by": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Signers of the transaction",
				Elem:        schema.TypeString,
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Initiator of the transaction",
			},
			"rejected_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User that rejected the transaction (in case it was rejected)",
			},
			"created_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Date when transaction was created",
			},
			"last_updated": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Date when transaction was last updates",
			},
			// TODO: to complete
			// "authorization_info": {
			// 	Type:        schema.TypeMap,
			// 	Computed:    true,
			// 	Description: "The information about Transaction Authorization Policy (TAP).",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"allow_operator_as_authorizer": {
			// 				Type:        schema.TypeBool,
			// 				Description: "Set to true if the intiator of the transaction can be one of the approvers",
			// 			},
			// 			"logic": {
			// 				Type:        schema.TypeString,
			// 				Description: "This is the logic that is applied between the different authorization groups listed below",
			// 			},
			// 			"groups": {
			// 				Type:        schema.TypeMap,
			// 				Optional:    true,
			// 				Description: "The list of authorization groups and users that are required to approve this transaction. The logic applied between the different groups is the “logic” field above. Each element in the response is the user ID (the can found see via the users endpoint) and theirApprovalStatus",
			// 				Elem: &schema.Resource{
			// 					Schema: map[string]*schema.Schema{
			// 						"address": {
			// 							Type:        schema.TypeString,
			// 							Optional:    true,
			// 						},
			// 						"tag": {
			// 							Type:     schema.TypeString,
			// 							Optional: true,
			// 						}
			// 					},
			// 				},
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func transactionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	tx, err := sdk.GetTransaction(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("asset_id", tx.AssetID)
	d.Set("source_id", tx.Source.ID)
	d.Set("source_type", tx.Source.TransferType)
	d.Set("source_name", tx.Source.Name)
	d.Set("amount", tx.Amount)
	d.Set("operation", string(tx.Operation))
	d.Set("customer_ref_id", tx.CustomerRefID)
	d.Set("extra_parameters", tx.ExtraParameters)
	d.Set("note", tx.Note)
	d.Set("status", string(tx.Status))
	d.Set("sub_status", string(tx.SubStatus))
	d.Set("tx_hash", tx.TxHash)
	d.Set("block_height", tx.BlockInfo.BlockHeight)
	d.Set("block_hash", tx.BlockInfo.BlockHash)
	d.Set("source_address", tx.SourceAddress)
	d.Set("destination_address", tx.DestinationAddress)
	d.Set("signed_by", tx.SignedBy)
	d.Set("created_by", tx.CreatedBy)
	d.Set("rejected_by", tx.RejectedBy)
	d.Set("created_at", tx.CreatedAt)
	d.Set("last_updated", tx.LastUpdated)

	return nil
}

func transactionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)

	createTxMsg := &fireblocksclient.CreateTransactionMsg{
		AssetID: d.Get("asset_id").(string),
		Source: fireblocksclient.TransferPeerPath{
			ID:   d.Get("source_id").(string),
			Type: fireblocksclient.PeerType(d.Get("source_type").(string)),
		},
		Destination: fireblocksclient.DestinationTransferPeerPath{
			ID:   d.Get("destination_id").(string),
			Type: fireblocksclient.PeerType(d.Get("destination_type").(string)),
			OneTimeAddress: fireblocksclient.OneTimeAddress{
				Address: d.Get("destination_one_time_address").(string),
			},
		},
		Amount:        d.Get("amount").(string),
		GasPrice:      d.Get("gas_price").(string),
		GasLimit:      d.Get("gas_limit").(string),
		PriorityFee:   d.Get("priority_fee").(string),
		Note:          d.Get("note").(string),
		Operation:     fireblocksclient.TransactionOperation(d.Get("operation").(string)),
		CustomerRefID: d.Get("customer_ref_id").(string),
		ExtraParameters: fireblocksclient.ExtraParameters{
			ContractCallData: d.Get("extra_parameters.contractCallData").(string),
		},
	}

	resp, err := sdk.CreateTransaction(ctx, createTxMsg)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.ID)

	return transactionRead(ctx, d, meta)
}

func transactionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("transaction can not be updated")
}

func transactionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdk := (meta).(*fireblockssdk.SDK)
	err := sdk.CancelTransaction(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
