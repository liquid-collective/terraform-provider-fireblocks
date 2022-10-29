package clienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) CreateVaultAccount(ctx context.Context, msg *client.CreateVaultAccountMsg) (*client.VaultAccount, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, "/v1/vault/accounts", body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	account := new(client.VaultAccount)
	err = json.NewDecoder(resp.Body).Decode(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *Client) ListVaultAccountsPaged(ctx context.Context) (*client.VaultAccountsWithPageInfoMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/vault/accounts_paged", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// TODO: add support for query parameters (see https://docs.fireblocks.com/api/?python#list-vault-accounts-paged)

	defer resp.Body.Close()
	vaults := new(client.VaultAccountsWithPageInfoMsg)
	err = json.NewDecoder(resp.Body).Decode(vaults)
	if err != nil {
		return nil, err
	}

	return vaults, nil
}

func (c *Client) GetVaultAccount(ctx context.Context, vaultAccountID string) (*client.VaultAccount, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/vault/accounts/%v", vaultAccountID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	account := new(client.VaultAccount)
	err = json.NewDecoder(resp.Body).Decode(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *Client) CreateVaultAccountAsset(ctx context.Context, vaultID, assetID string) (*client.CreateVaultAssetResponse, error) {
	req, err := c.createReq(ctx, http.MethodPost, fmt.Sprintf("/v1/vault/accounts/%v/%v", vaultID, assetID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	asset := new(client.CreateVaultAssetResponse)
	err = json.NewDecoder(resp.Body).Decode(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (c *Client) UpdateVaultAccount(ctx context.Context, vaultAccountID string, msg *client.UpdateVaultAccountMsg) error {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return err
	}

	req, err := c.createReq(ctx, http.MethodPut, fmt.Sprintf("/v1/vault/accounts/%v", vaultAccountID), body)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) HideVaultAccount(ctx context.Context, vaultAccountID string) error {
	req, err := c.createReq(ctx, http.MethodPost, fmt.Sprintf("/v1/vault/accounts/%v/hide", vaultAccountID), http.NoBody)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UnhideVaultAccount(ctx context.Context, vaultAccountID string) error {
	req, err := c.createReq(ctx, http.MethodPost, fmt.Sprintf("/v1/vault/accounts/%v/unhide", vaultAccountID), http.NoBody)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SetAutoFuelVaultAccount(ctx context.Context, vaultAccountID string, status bool) error {
	// TODO: to be implemented
	return nil
}

func (c *Client) GetVaultAccountAssetBalance(ctx context.Context, vaultID, assetID string) (*client.VaultAsset, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/vault/accounts/%v/%v", vaultID, assetID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	asset := new(client.VaultAsset)
	err = json.NewDecoder(resp.Body).Decode(asset)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (c *Client) ListVaultAccountAssetAddresses(ctx context.Context, vaultID, assetID string) ([]*client.VaultAccountAssetAddress, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/vault/accounts/%v/%v/addresses", vaultID, assetID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var addresses []*client.VaultAccountAssetAddress
	err = json.NewDecoder(resp.Body).Decode(&addresses)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}
