package vrcapi

import (
	"context"
	"fmt"

	"github.com/kqnade/vrcgo/shared"
)

// GetConfig はシステム設定を取得します
func (c *Client) GetConfig(ctx context.Context) (*shared.Config, error) {
	var config shared.Config
	err := c.doRequest(ctx, "GET", "/config", nil, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	return &config, nil
}

// GetTime は現在のAPIサーバー時刻を取得します
func (c *Client) GetTime(ctx context.Context) (string, error) {
	var response struct {
		Time string `json:"time"`
	}
	err := c.doRequest(ctx, "GET", "/time", nil, &response)
	if err != nil {
		return "", fmt.Errorf("failed to get time: %w", err)
	}
	return response.Time, nil
}

// GetHealth はAPIヘルスチェックを行います
func (c *Client) GetHealth(ctx context.Context) (bool, error) {
	var response struct {
		OK bool `json:"ok"`
	}
	err := c.doRequest(ctx, "GET", "/health", nil, &response)
	if err != nil {
		return false, fmt.Errorf("failed to get health: %w", err)
	}
	return response.OK, nil
}

// GetInfoPush は情報プッシュを取得します
func (c *Client) GetInfoPush(ctx context.Context) ([]shared.InfoPush, error) {
	var infoPushes []shared.InfoPush
	err := c.doRequest(ctx, "GET", "/infoPush", nil, &infoPushes)
	if err != nil {
		return nil, fmt.Errorf("failed to get info push: %w", err)
	}
	return infoPushes, nil
}

// GetOnlineUsers はオンラインユーザー数を取得します
func (c *Client) GetOnlineUsers(ctx context.Context) (int, error) {
	var response struct {
		Count int `json:"count"`
	}
	err := c.doRequest(ctx, "GET", "/visits", nil, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to get online users: %w", err)
	}
	return response.Count, nil
}
