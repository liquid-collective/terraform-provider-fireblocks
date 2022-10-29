//go:build integration

package clienthttp

import "golang.org/x/net/context"

var (
	testAPIKey        = ""
	testRSAPrivateKey = ""
)

func newTestClient() (*Client, error) {
	cfg := &Config{
		APIKey:        testAPIKey,
		RSAPrivateKey: testRSAPrivateKey,
	}

	client, err := New(cfg.SetDefault())
	if err != nil {
		return nil, err
	}

	err = client.Init(context.TODO())
	if err != nil {
		return nil, err
	}

	return client, nil
}
