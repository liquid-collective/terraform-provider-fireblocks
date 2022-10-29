package clienthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liquid-collective/terraform-provider-fireblocks/pkg/fireblocks/client"
)

func (c *Client) CreateTransaction(ctx context.Context, msg *client.CreateTransactionMsg) (*client.CreateTransactionRespMsg, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.createReq(ctx, http.MethodPost, "/v1/transactions", body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	tx := new(client.CreateTransactionRespMsg)
	err = json.NewDecoder(resp.Body).Decode(tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) ListTransactions(ctx context.Context) ([]*client.TransactionMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, "/v1/transactions", http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// TODO: add support for query parameters (see https://docs.fireblocks.com/api/?python#list-vault-accounts-paged)

	defer resp.Body.Close()
	var txs []*client.TransactionMsg
	err = json.NewDecoder(resp.Body).Decode(&txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (c *Client) GetTransaction(ctx context.Context, txID string) (*client.TransactionMsg, error) {
	req, err := c.createReq(ctx, http.MethodGet, fmt.Sprintf("/v1/transactions/%v", txID), http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	tx := new(client.TransactionMsg)
	err = json.NewDecoder(resp.Body).Decode(tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) CancelTransaction(ctx context.Context, txID string) error {
	req, err := c.createReq(ctx, http.MethodPost, fmt.Sprintf("/v1/transactions/%v/cancel", txID), http.NoBody)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
