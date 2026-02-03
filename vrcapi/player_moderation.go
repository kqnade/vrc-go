package vrcapi

import (
	"context"
	"fmt"

	"github.com/kqnade/vrcgo/shared"
)

// GetPlayerModerations はプレイヤーモデレーションのリストを取得します
func (c *Client) GetPlayerModerations(ctx context.Context, moderationType string) ([]shared.PlayerModeration, error) {
	var moderations []shared.PlayerModeration
	path := "/auth/user/playermoderations"
	if moderationType != "" {
		path += "?type=" + moderationType
	}
	err := c.doRequest(ctx, "GET", path, nil, &moderations)
	if err != nil {
		return nil, fmt.Errorf("failed to get player moderations: %w", err)
	}
	return moderations, nil
}

// ModeratePlayer はプレイヤーをモデレートします
func (c *Client) ModeratePlayer(ctx context.Context, moderatedUserID, moderationType string) (*shared.PlayerModeration, error) {
	var moderation shared.PlayerModeration
	req := struct {
		ModeratedUserID string `json:"moderated"`
		Type            string `json:"type"`
	}{
		ModeratedUserID: moderatedUserID,
		Type:            moderationType,
	}
	err := c.doRequest(ctx, "POST", "/auth/user/playermoderations", req, &moderation)
	if err != nil {
		return nil, fmt.Errorf("failed to moderate player: %w", err)
	}
	return &moderation, nil
}

// UnmoderatePlayer はプレイヤーのモデレーションを解除します
func (c *Client) UnmoderatePlayer(ctx context.Context, moderatedUserID, moderationType string) error {
	req := struct {
		ModeratedUserID string `json:"moderated"`
		Type            string `json:"type"`
	}{
		ModeratedUserID: moderatedUserID,
		Type:            moderationType,
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/unplayermoderate", req, nil)
	if err != nil {
		return fmt.Errorf("failed to unmoderate player: %w", err)
	}
	return nil
}

// ClearAllPlayerModerations はすべてのプレイヤーモデレーションをクリアします
func (c *Client) ClearAllPlayerModerations(ctx context.Context) error {
	err := c.doRequest(ctx, "DELETE", "/auth/user/playermoderations", nil, nil)
	if err != nil {
		return fmt.Errorf("failed to clear all player moderations: %w", err)
	}
	return nil
}
