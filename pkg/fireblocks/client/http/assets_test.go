//go:build integration

package clienthttp

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAssets(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	assets, err := c.ListAssets(context.TODO())
	require.NoError(t, err)
	b, _ := json.MarshalIndent(assets, "", "  ")
	t.Logf("%v", string(b))
	assert.False(t, true)
}
