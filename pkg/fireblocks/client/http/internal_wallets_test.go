//go:build integration

package clienthttp

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListInternalWallets(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	wallets, err := c.ListInternalWallets(context.TODO())
	require.NoError(t, err)
	b, _ := json.MarshalIndent(wallets, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}
