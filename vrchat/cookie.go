package vrchat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// cookieStore はCookieの保存形式です
type cookieStore struct {
	Cookies []*http.Cookie `json:"cookies"`
}

// SaveCookies はCookieをファイルに保存します
func (c *Client) SaveCookies(path string) error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}

	cookies := c.httpClient.Jar.Cookies(u)
	store := cookieStore{Cookies: cookies}

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cookies: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("failed to write cookie file: %w", err)
	}

	return nil
}

// LoadCookies はCookieをファイルから読み込みます
func (c *Client) LoadCookies(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read cookie file: %w", err)
	}

	var store cookieStore
	if err := json.Unmarshal(data, &store); err != nil {
		return fmt.Errorf("failed to unmarshal cookies: %w", err)
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}

	c.httpClient.Jar.SetCookies(u, store.Cookies)

	return nil
}
