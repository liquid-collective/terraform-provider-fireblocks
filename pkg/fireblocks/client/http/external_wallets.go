package clienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) CreateExternalWallet(ctx context.Context, msg *client.CreateExternalWalletMsg) (*client.ExternalWallet, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, "/v1/external_wallets", body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	wallet := new(client.ExternalWallet)
	err = json.NewDecoder(resp.Body).Decode(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (c *Client) ListExternalWallets(ctx context.Context) ([]*client.ExternalWallet, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/external_wallets", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var wallets []*client.ExternalWallet
	err = json.NewDecoder(resp.Body).Decode(&wallets)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (c *Client) GetExternalWallet(ctx context.Context, walletID string) (*client.ExternalWallet, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/external_wallets/%v", walletID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	wallet := new(client.ExternalWallet)
	err = json.NewDecoder(resp.Body).Decode(wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (c *Client) DeleteExternalWallet(ctx context.Context, walletID string) error {
	req, err := c.createReq(ctx, http.MethodDelete, fmt.Sprintf("/v1/external_wallets/%v", walletID), http.NoBody)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AddExternalWalletAsset(ctx context.Context, walletID, assetID string, msg *client.AddExternalWalletAssetMsg) (*client.ExternalWalletAsset, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, fmt.Sprintf("/v1/external_wallets/%v/%v", walletID, assetID), body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	asset := new(client.ExternalWalletAsset)
	err = json.NewDecoder(resp.Body).Decode(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (c *Client) GetExternalWalletAsset(ctx context.Context, walletID, assetID string) (*client.ExternalWalletAsset, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/external_wallets/%v/%v", walletID, assetID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	asset := new(client.ExternalWalletAsset)
	err = json.NewDecoder(resp.Body).Decode(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (c *Client) DeleteExternalWalletAsset(ctx context.Context, walletID, assetID string) error {
	req, err := c.createReq(ctx, http.MethodDelete, fmt.Sprintf("/v1/external_wallets/%v/%v", walletID, assetID), http.NoBody)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
