package vrcapi

import (
	"net/http"
	"time"

	"github.com/kqnade/vrcgo/shared"
)

// Option はクライアント設定オプションです
type Option = shared.Option

// WithUserAgent はUser-Agentを設定します
func WithUserAgent(ua string) Option {
	return shared.WithUserAgent(ua)
}

// WithProxy はプロキシを設定します
func WithProxy(proxyURL string) Option {
	return shared.WithProxy(proxyURL)
}

// WithTimeout はタイムアウトを設定します
func WithTimeout(timeout time.Duration) Option {
	return shared.WithTimeout(timeout)
}

// WithHTTPClient はカスタムHTTPクライアントを設定します
func WithHTTPClient(httpClient *http.Client) Option {
	return shared.WithHTTPClient(httpClient)
}

// WithBaseURL はベースURLを設定します（テスト用）
func WithBaseURL(baseURL string) Option {
	return shared.WithBaseURL(baseURL)
}
