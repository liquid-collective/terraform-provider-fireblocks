//go:build integration

package clienthttp

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func TestListTransactions(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	txs, err := c.ListTransactions(context.TODO())
	require.NoError(t, err)
	b, _ := json.MarshalIndent(txs, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}

func TestCreateTransaction(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	tx := &client.CreateTransactionMsg{
		AssetID: "ETH_TEST3",
		Amount:  "0",
		Note:    "Test API ERC20 transfer",
		Source: client.TransferPeerPath{
			Type: client.PeerTypeVaultAccount,
			ID:   "2",
		},
		Operation: client.TransactionOperationContractCall,
		Destination: client.DestinationTransferPeerPath{
			Type: client.PeerTypeExternalWallet,
			ID:   "0a06bab7-c069-8c66-b0c9-73a38449e3c3",
			OneTimeAddress: client.OneTimeAddress{
				Address: "0xF43cBB88e6487A819A1F37bB1c558d97e7167dc6",
			},
		},
		ExtraParameters: client.ExtraParameters{
			ContractCallData: "0xd656d80a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007800000000000000000000000000000000000000000000000000000000000000f0000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000d3c21bcecceda10000000000000000000000000000002b7ff5d4c14a9da8d5c9354c7a52ab40ddc1c01e0000000000000000000000000000000000000000000000000000000000000000",
		},
	}
	b, _ := json.MarshalIndent(tx, "", "  ")
	t.Logf("%v", string(b))
	resp, err := c.CreateTransaction(context.TODO(), tx)
	require.NoError(t, err)
	b, _ = json.MarshalIndent(resp, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}

func TestCancelTransaction(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	err = c.CancelTransaction(context.TODO(), "")
	require.NoError(t, err)
	assert.False(t, true)
}

func TestGetTransaction(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	tx, err := c.GetTransaction(context.TODO(), "d10d5334-a46b-49ad-87af-41eb18b4e4d7")
	require.NoError(t, err)
	b, _ := json.MarshalIndent(tx, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}
