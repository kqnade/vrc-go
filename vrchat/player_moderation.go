package vrchat

import (
	"context"
	"fmt"
)

// PlayerModeration はプレイヤーモデレーション情報です
type PlayerModeration struct {
	ID               string `json:"id"`
	Type             string `json:"type"` // "mute", "unmute", "block", "unblock", "hideAvatar", "showAvatar"
	SourceUserID     string `json:"sourceUserId"`
	SourceDisplayName string `json:"sourceDisplayName"`
	TargetUserID     string `json:"targetUserId"`
	TargetDisplayName string `json:"targetDisplayName"`
	Created          string `json:"created"`
}

// ModerateUserRequest はユーザーモデレーションリクエストです
type ModerateUserRequest struct {
	ModeratedUserId string `json:"moderated"`
	Type            string `json:"type"` // "mute", "unmute", "block", "unblock", "hideAvatar", "showAvatar"
}

// ModerateUser はユーザーをモデレートします
func (c *Client) ModerateUser(ctx context.Context, req ModerateUserRequest) (*PlayerModeration, error) {
	var moderation PlayerModeration
	err := c.doRequest(ctx, "POST", "/auth/user/playermoderations", req, &moderation)
	if err != nil {
		return nil, fmt.Errorf("failed to moderate user: %w", err)
	}
	return &moderation, nil
}

// GetPlayerModerations はプレイヤーモデレーションのリストを取得します
func (c *Client) GetPlayerModerations(ctx context.Context) ([]PlayerModeration, error) {
	var moderations []PlayerModeration
	err := c.doRequest(ctx, "GET", "/auth/user/playermoderations", nil, &moderations)
	if err != nil {
		return nil, fmt.Errorf("failed to get player moderations: %w", err)
	}
	return moderations, nil
}

// ClearPlayerModeration はプレイヤーモデレーションをクリアします
func (c *Client) ClearPlayerModeration(ctx context.Context, moderatedUserId string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/auth/user/playermoderations/"+moderatedUserId, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to clear player moderation: %w", err)
	}
	return nil
}

// GetPlayerModeration は特定のユーザーに対するモデレーション情報を取得します
func (c *Client) GetPlayerModeration(ctx context.Context, moderatedUserId string) (*PlayerModeration, error) {
	var moderation PlayerModeration
	err := c.doRequest(ctx, "GET", "/auth/user/playermoderations/"+moderatedUserId, nil, &moderation)
	if err != nil {
		return nil, fmt.Errorf("failed to get player moderation: %w", err)
	}
	return &moderation, nil
}

// MuteUser はユーザーをミュートします
func (c *Client) MuteUser(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "mute",
	})
}

// UnmuteUser はユーザーのミュートを解除します
func (c *Client) UnmuteUser(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "unmute",
	})
}

// BlockUser はユーザーをブロックします
func (c *Client) BlockUser(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "block",
	})
}

// UnblockUser はユーザーのブロックを解除します
func (c *Client) UnblockUser(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "unblock",
	})
}

// HideUserAvatar はユーザーのアバターを非表示にします
func (c *Client) HideUserAvatar(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "hideAvatar",
	})
}

// ShowUserAvatar はユーザーのアバターを表示します
func (c *Client) ShowUserAvatar(ctx context.Context, userID string) (*PlayerModeration, error) {
	return c.ModerateUser(ctx, ModerateUserRequest{
		ModeratedUserId: userID,
		Type:            "showAvatar",
	})
}
