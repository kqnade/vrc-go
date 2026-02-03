package vrchat

import (
	"context"
	"fmt"
)

// AuthConfig は認証設定です
type AuthConfig struct {
	Username string
	Password string
	TOTPCode string // 2要素認証コード（オプション）
}

// CurrentUser は現在のユーザー情報です
type CurrentUser struct {
	ID                      string   `json:"id"`
	DisplayName             string   `json:"displayName"`
	Username                string   `json:"username"`
	Bio                     string   `json:"bio"`
	Tags                    []string `json:"tags"`
	Status                  string   `json:"status"`
	StatusDescription       string   `json:"statusDescription"`
	CurrentAvatar           string   `json:"currentAvatar"`
	CurrentAvatarThumbnail  string   `json:"currentAvatarImageUrl"`
	RequiresTwoFactorAuth   []string `json:"requiresTwoFactorAuth,omitempty"`
	EmailVerified           bool     `json:"emailVerified"`
	HasBirthday             bool     `json:"hasBirthday"`
	HasEmail                bool     `json:"hasEmail"`
	HasPendingEmail         bool     `json:"hasPendingEmail"`
	ObfuscatedEmail         string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail  string   `json:"obfuscatedPendingEmail"`
	SteamID                 string   `json:"steamId"`
	OculusID                string   `json:"oculusId"`
	AccountDeletionDate     *string  `json:"accountDeletionDate,omitempty"`
	AccountDeletionLog      *string  `json:"accountDeletionLog,omitempty"`
	AcceptedTOSVersion      int      `json:"acceptedTOSVersion"`
	AcceptedPrivacyVersion  int      `json:"acceptedPrivacyVersion"`
	SteamDetails            struct{} `json:"steamDetails"`
	OculusDetails           struct{} `json:"oculusDetails"`
	HasLoggedInFromClient   bool     `json:"hasLoggedInFromClient"`
	FriendKey               string   `json:"friendKey"`
	OnlineFriends           []string `json:"onlineFriends"`
	ActiveFriends           []string `json:"activeFriends"`
	OfflineFriends          []string `json:"offlineFriends"`
	FriendGroupNames        []string `json:"friendGroupNames"`
	CurrentAvatarAssetURL   string   `json:"currentAvatarAssetUrl"`
	FallbackAvatar          string   `json:"fallbackAvatar"`
	IsFriend                bool     `json:"isFriend"`
	LastLogin               string   `json:"last_login"`
	LastPlatform            string   `json:"last_platform"`
	AllowAvatarCopying      bool     `json:"allowAvatarCopying"`
	State                   string   `json:"state"`
	DateJoined              string   `json:"date_joined"`
	PastDisplayNames        []struct {
		DisplayName string `json:"displayName"`
		UpdatedAt   string `json:"updated_at"`
	} `json:"pastDisplayNames"`
	TwoFactorAuthEnabled      bool `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate  *string `json:"twoFactorAuthEnabledDate,omitempty"`
}

// Authenticate はVRChat APIにログインします
func (c *Client) Authenticate(ctx context.Context, config AuthConfig) error {
	var user CurrentUser
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

// TwoFactorAuthRequest は2FA検証リクエストです
type TwoFactorAuthRequest struct {
	Code string `json:"code"`
}

// TwoFactorAuthResponse は2FA検証レスポンスです
type TwoFactorAuthResponse struct {
	Verified bool `json:"verified"`
}

// verifyTwoFactor は2要素認証を実行します
func (c *Client) verifyTwoFactor(ctx context.Context, code string) error {
	req := TwoFactorAuthRequest{Code: code}
	var resp TwoFactorAuthResponse

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
func (c *Client) GetCurrentUser(ctx context.Context) (*CurrentUser, error) {
	var user CurrentUser
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
