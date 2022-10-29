package clienthttp

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) ListAssets(ctx context.Context) ([]*client.AssetTypeResponse, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/supported_assets", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var assets []*client.AssetTypeResponse
	err = json.NewDecoder(resp.Body).Decode(&assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
