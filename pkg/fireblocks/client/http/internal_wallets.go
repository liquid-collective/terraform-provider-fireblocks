package clienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) CreateInternalWallet(ctx context.Context, msg *client.CreateInternalWalletMsg) (*client.UnmanagedWallet, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, "/v1/internal_wallets", body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	wallet := new(client.UnmanagedWallet)
	err = json.NewDecoder(resp.Body).Decode(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (c *Client) ListInternalWallets(ctx context.Context) ([]*client.UnmanagedWallet, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/internal_wallets", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var wallets []*client.UnmanagedWallet
	err = json.NewDecoder(resp.Body).Decode(&wallets)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (c *Client) GetInternalWallet(ctx context.Context, walletID string) (*client.UnmanagedWallet, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/internal_wallets/%v", walletID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	wallet := new(client.UnmanagedWallet)
	err = json.NewDecoder(resp.Body).Decode(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
