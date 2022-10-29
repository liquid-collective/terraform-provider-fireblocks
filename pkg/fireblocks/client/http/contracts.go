package clienthttp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) ListContracts(ctx context.Context) ([]*client.UnmanagedContract, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/contracts", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var contracts []*client.UnmanagedContract
	err = json.NewDecoder(resp.Body).Decode(&contracts)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (c *Client) GetContract(ctx context.Context, contractID string) (*client.UnmanagedContract, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/contracts/%v", contractID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contract := new(client.UnmanagedContract)
	err = json.NewDecoder(resp.Body).Decode(contract)
	if err != nil {
		return nil, err
	}

	return contract, nil
}
