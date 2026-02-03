package vrchat

import (
	"net/http"
	"net/url"
	"time"
)

// Option はクライアント設定オプションです
type Option func(*Client)

// WithUserAgent はUser-Agentを設定します
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// WithProxy はプロキシを設定します
func WithProxy(proxyURL string) Option {
	return func(c *Client) {
		proxy, err := url.Parse(proxyURL)
		if err != nil {
			return
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
		c.httpClient.Transport = transport
	}
}

// WithTimeout はタイムアウトを設定します
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient はカスタムHTTPクライアントを設定します
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL はベースURLを設定します（テスト用）
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
