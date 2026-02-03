package shared

import (
	"net/http"
	"net/url"
	"time"
)

// ClientConfig はクライアント設定を保持します
type ClientConfig struct {
	UserAgent  string
	Timeout    time.Duration
	Proxy      *url.URL
	HTTPClient *http.Client
	BaseURL    string
}

// Option はクライアント設定オプションです
type Option func(*ClientConfig)

// WithUserAgent はUser-Agentを設定します
func WithUserAgent(ua string) Option {
	return func(c *ClientConfig) {
		c.UserAgent = ua
	}
}

// WithProxy はプロキシを設定します
func WithProxy(proxyURL string) Option {
	return func(c *ClientConfig) {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return
		}
		c.Proxy = proxy
	}
}

// WithTimeout はタイムアウトを設定します
func WithTimeout(timeout time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = timeout
	}
}

// WithHTTPClient はカスタムHTTPクライアントを設定します
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *ClientConfig) {
		c.HTTPClient = httpClient
	}
}

// WithBaseURL はベースURLを設定します（テスト用）
func WithBaseURL(baseURL string) Option {
	return func(c *ClientConfig) {
		c.BaseURL = baseURL
	}
}
