package clienthttp

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Client struct {
	cfg *Config

	rand *rand.Rand

	privateKey *rsa.PrivateKey
	client     *http.Client
}

func New(cfg *Config) (*Client, error) {
	seed, err := randSeed()
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:    cfg,
		rand:   rand.New(rand.NewSource(int64(seed))),
		client: http.DefaultClient,
	}, nil
}

func randSeed() (seed uint64, err error) {
	err = binary.Read(crand.Reader, binary.BigEndian, &seed)
	return
}

func (c *Client) Init(ctx context.Context) error {
	privk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(c.cfg.RSAPrivateKey))
	if err != nil {
		return err
	}

	c.privateKey = privk

	return nil
}

func (c *Client) createAndSignJWTToken(path string, body []byte) (string, error) {
	token := &jwt.MapClaims{
		"uri":      path,
		"nonce":    c.rand.Int63(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Second * 55).Unix(),
		"sub":      c.cfg.APIKey,
		"bodyHash": createHash(body),
	}

	j := jwt.NewWithClaims(jwt.SigningMethodRS256, token)
	signedToken, err := j.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, err
}

func createHash(data []byte) string {
	h := sha256.New()
	_, _ = h.Write(data)
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func (c *Client) createReq(ctx context.Context, method, route string, body io.Reader) (*http.Request, error) {
	var buf bytes.Buffer
	tee := io.TeeReader(body, &buf)

	b, err := io.ReadAll(tee)
	if err != nil {
		return nil, err
	}

	token, err := c.createAndSignJWTToken(route, b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%v%v", c.cfg.APIURL, route), &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", c.cfg.APIKey)

	return req, nil
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	fmt.Printf("Piou doReques %v\n", req.URL.String())
	resp, err := c.client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode >= 300 {
		msg := new(ErrorMsg)
		err = json.NewDecoder(resp.Body).Decode(msg)
		if err != nil {
			raw := new(json.RawMessage)
			_ = json.NewDecoder(resp.Body).Decode(&raw)
			return resp, fmt.Errorf("HTTP Error [status code=%v] [body=%v]", resp.StatusCode, string(*raw))
		}
		return resp, fmt.Errorf("HTTP Error [status code=%v] [message=%v]", resp.StatusCode, msg.Message)
	}

	return resp, err
}
