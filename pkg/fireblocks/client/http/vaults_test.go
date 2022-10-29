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

func TestListVaultAccountsPaged(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	accounts, err := c.ListVaultAccountsPaged(context.TODO())
	require.NoError(t, err)
	b, _ := json.MarshalIndent(accounts.Accounts, "", "  ")
	t.Logf(string(b))
	assert.False(t, true)
}

func TestGetVault(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	accountID := "2"
	account, err := c.GetVaultAccount(context.TODO(), accountID)
	require.NoError(t, err)
	require.NoError(t, err)
	b, _ := json.MarshalIndent(account, "", "  ")
	t.Logf(string(b))
	assert.False(t, true)
}

func TestCreateVaultAccount(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	name := "Test"
	account, err := c.CreateVaultAccount(
		context.TODO(),
		&client.CreateVaultAccountMsg{
			Name: name,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, name, account.Name)
}

func TestListVaultAccountAssetAddresses(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	addresses, err := c.ListVaultAccountAssetAddresses(context.TODO(), "1", "ETH_TEST3")
	require.NoError(t, err)
	b, _ := json.MarshalIndent(addresses, "", "  ")
	t.Logf(string(b))
	assert.False(t, true)
}
