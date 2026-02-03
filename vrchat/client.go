package vrchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

const (
	// DefaultBaseURL はVRChat APIのデフォルトベースURLです
	DefaultBaseURL = "https://api.vrchat.cloud/api/1"
	// DefaultUserAgent はデフォルトのUser-Agentです
	DefaultUserAgent = "vrc-go/0.1.0"
)

// Client はVRChat APIクライアントです
type Client struct {
	httpClient *http.Client
	baseURL    string
	userAgent  string
}

// NewClient は新しいVRChat APIクライアントを作成します
func NewClient(opts ...Option) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	c := &Client{
		httpClient: &http.Client{
			Jar:     jar,
			Timeout: time.Minute,
		},
		baseURL:   DefaultBaseURL,
		userAgent: DefaultUserAgent,
	}

	// オプション適用
	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

// doRequest はHTTPリクエストを実行し、レスポンスをデコードします
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// エラーレスポンスのチェック
	if resp.StatusCode >= 400 {
		var apiErr struct {
			Error struct {
				Message    string `json:"message"`
				StatusCode int    `json:"status_code"`
			} `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err == nil && apiErr.Error.Message != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error.Message,
			}
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}

	// 成功レスポンスのデコード
	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// doRequestWithBasicAuth はBasic認証付きでHTTPリクエストを実行します
func (c *Client) doRequestWithBasicAuth(ctx context.Context, method, path, username, password string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.SetBasicAuth(username, password)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// エラーレスポンスのチェック
	if resp.StatusCode >= 400 {
		var apiErr struct {
			Error struct {
				Message    string `json:"message"`
				StatusCode int    `json:"status_code"`
			} `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err == nil && apiErr.Error.Message != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error.Message,
			}
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    resp.Status,
		}
	}

	// 成功レスポンスのデコード
	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
