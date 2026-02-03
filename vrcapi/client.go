package vrcapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"

	"github.com/kqnade/vrcgo/shared"
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
func NewClient(opts ...shared.Option) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	config := &shared.ClientConfig{
		UserAgent: DefaultUserAgent,
		Timeout:   time.Minute,
		BaseURL:   DefaultBaseURL,
	}

	// オプション適用
	for _, opt := range opts {
		opt(config)
	}

	// HTTPクライアントの設定
	var httpClient *http.Client
	if config.HTTPClient != nil {
		httpClient = config.HTTPClient
		if httpClient.Jar == nil {
			httpClient.Jar = jar
		}
	} else {
		httpClient = &http.Client{
			Jar:     jar,
			Timeout: config.Timeout,
		}
		if config.Proxy != nil {
			transport := &http.Transport{
				Proxy: http.ProxyURL(config.Proxy),
			}
			httpClient.Transport = transport
		}
	}

	c := &Client{
		httpClient: httpClient,
		baseURL:    config.BaseURL,
		userAgent:  config.UserAgent,
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
			return &shared.APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error.Message,
			}
		}
		return &shared.APIError{
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
			return &shared.APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error.Message,
			}
		}
		return &shared.APIError{
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

// GetAuthCookie はCookieJarから認証クッキーを取得します
func (c *Client) GetAuthCookie() (string, error) {
	u, err := c.baseURLParsed()
	if err != nil {
		return "", err
	}

	cookies := c.httpClient.Jar.Cookies(u)
	for _, cookie := range cookies {
		if cookie.Name == "auth" || cookie.Name == "authcookie" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("auth cookie not found")
}

// baseURLParsed はベースURLをパースして返します
func (c *Client) baseURLParsed() (*url.URL, error) {
	return url.Parse(c.baseURL)
}

// SaveCookies はCookieをファイルに保存します
func (c *Client) SaveCookies(path string) error {
	return shared.SaveCookies(c.httpClient.Jar, c.baseURL, path)
}

// LoadCookies はCookieをファイルから読み込みます
func (c *Client) LoadCookies(path string) error {
	return shared.LoadCookies(c.httpClient.Jar, c.baseURL, path)
}
