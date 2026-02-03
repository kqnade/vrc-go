package vrcapi

import (
	"context"
	"fmt"

	"github.com/kqnade/vrcgo/shared"
)

// Authenticate はVRChat APIにログインします
func (c *Client) Authenticate(ctx context.Context, config shared.AuthConfig) error {
	var user shared.CurrentUser
	err := c.doRequestWithBasicAuth(
		ctx,
		"GET",
		"/auth/user",
		config.Username,
		config.Password,
		nil,
		&user,
	)
	if err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}

	// 2FA必要チェック
	if len(user.RequiresTwoFactorAuth) > 0 {
		if config.TOTPCode == "" {
			return fmt.Errorf("two-factor authentication required (methods: %v) but TOTP code not provided", user.RequiresTwoFactorAuth)
		}
		return c.verifyTwoFactor(ctx, config.TOTPCode)
	}

	return nil
}

// verifyTwoFactor は2要素認証を実行します
func (c *Client) verifyTwoFactor(ctx context.Context, code string) error {
	req := shared.TwoFactorAuthRequest{Code: code}
	var resp shared.TwoFactorAuthResponse

	err := c.doRequest(ctx, "POST", "/auth/twofactorauth/totp/verify", req, &resp)
	if err != nil {
		return fmt.Errorf("two-factor authentication failed: %w", err)
	}

	if !resp.Verified {
		return fmt.Errorf("two-factor authentication code invalid")
	}

	return nil
}

// GetCurrentUser は現在のユーザー情報を取得します
func (c *Client) GetCurrentUser(ctx context.Context) (*shared.CurrentUser, error) {
	var user shared.CurrentUser
	err := c.doRequest(ctx, "GET", "/auth/user", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}
	return &user, nil
}

// Logout はログアウトします
func (c *Client) Logout(ctx context.Context) error {
	err := c.doRequest(ctx, "PUT", "/logout", nil, nil)
	if err != nil {
		return fmt.Errorf("logout failed: %w", err)
	}
	return nil
}
