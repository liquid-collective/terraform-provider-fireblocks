//go:build integration

package clienthttp

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListExternalWallets(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	wallets, err := c.ListExternalWallets(context.TODO())
	require.NoError(t, err)
	b, _ := json.MarshalIndent(wallets, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}

func TestGetExternalWalletAsset(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	asset, err := c.GetExternalWalletAsset(context.TODO(), "0a06bab7-c069-8c66-b0c9-73a38449e3c3", "ETH_TEST3")
	require.NoError(t, err)
	b, _ := json.MarshalIndent(asset, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}
