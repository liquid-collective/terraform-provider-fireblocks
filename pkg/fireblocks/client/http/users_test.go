//go:build integration

package clienthttp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListUsers(t *testing.T) {
	c, err := newTestClient()
	require.NoError(t, err)

	users, err := c.ListUsers(context.TODO())
	require.NoError(t, err)
	assert.Len(t, users, 4)
}
