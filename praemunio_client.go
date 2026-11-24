package praemunio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SandboxURL = "https://sandbox-api.praemunio.org/api/v1"
	LiveURL    = "https://api.praemunio.org/api/v1"
)

// PraemunioClient is a high-level client for the Praemunio API.
type PraemunioClient struct {
	PublicKey    string
	SecretKey    string
	BaseURL      string
	AccessToken  string
	RefreshToken string
	HTTPClient   *http.Client
}

// NewPraemunioClient creates a new client pointing to the sandbox environment.
func NewPraemunioClient(publicKey, secretKey string) *PraemunioClient {
	return &PraemunioClient{
		PublicKey:  publicKey,
		SecretKey:  secretKey,
		BaseURL:    SandboxURL,
		HTTPClient: http.DefaultClient,
	}
}

// NewPraemunioClientSandbox creates a new client pointing to the sandbox environment.
func NewPraemunioClientSandbox(publicKey, secretKey string) *PraemunioClient {
	return &PraemunioClient{
		PublicKey:  publicKey,
		SecretKey:  secretKey,
		BaseURL:    SandboxURL,
		HTTPClient: http.DefaultClient,
	}
}

// NewPraemunioClientLive creates a new client pointing to the live environment.
func NewPraemunioClientLive(publicKey, secretKey string) *PraemunioClient {
	return &PraemunioClient{
		PublicKey:  publicKey,
		SecretKey:  secretKey,
		BaseURL:    LiveURL,
		HTTPClient: http.DefaultClient,
	}
}

// Authenticate exchanges public/secret keys for access and refresh tokens.
func (c *PraemunioClient) Authenticate() error {
	body := map[string]string{
		"publicKey": c.PublicKey,
		"secretKey": c.SecretKey,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal auth request: %w", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, c.BaseURL+"/auth/auth-token", bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create auth request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Result  struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
		} `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode auth response: %w", err)
	}

	if !result.Success {
		return fmt.Errorf("authentication failed: %s", result.Message)
	}

	c.AccessToken = result.Result.AccessToken
	c.RefreshToken = result.Result.RefreshToken

	return nil
}

// GetAccessToken returns the current access token.
func (c *PraemunioClient) GetAccessToken() string {
	return c.AccessToken
}

// GetRefreshToken returns the current refresh token.
func (c *PraemunioClient) GetRefreshToken() string {
	return c.RefreshToken
}

// CreateAPI creates a new API client configured with the current access token.
func (c *PraemunioClient) CreateAPI() *APIClient {
	return &APIClient{
		BaseURL:     c.BaseURL,
		AccessToken: c.AccessToken,
		HTTPClient:  c.HTTPClient,
	}
}

// APIClient is a low-level API client that sends authenticated requests.
type APIClient struct {
	BaseURL     string
	AccessToken string
	HTTPClient  *http.Client
}

// NewAPIClient creates a new APIClient with a bearer token.
func NewAPIClient(baseURL, accessToken string) *APIClient {
	return &APIClient{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		HTTPClient:  http.DefaultClient,
	}