package clienthttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) ListUsers(ctx context.Context) ([]*client.UserMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/users", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var users []*client.UserMsg
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
