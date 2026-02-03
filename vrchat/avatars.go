package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// Avatar はアバター情報です
type Avatar struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Description         string   `json:"description"`
	AuthorID            string   `json:"authorId"`
	AuthorName          string   `json:"authorName"`
	Tags                []string `json:"tags"`
	AssetURL            string   `json:"assetUrl"`
	AssetURLObject      string   `json:"assetUrlObject"`
	ImageURL            string   `json:"imageUrl"`
	ThumbnailImageURL   string   `json:"thumbnailImageUrl"`
	ReleaseStatus       string   `json:"releaseStatus"`
	Version             int      `json:"version"`
	Featured            bool     `json:"featured"`
	UnityPackages       []UnityPackage `json:"unityPackages"`
	UnityPackageURL     string   `json:"unityPackageUrl"`
	UnityPackageURLObject string `json:"unityPackageUrlObject"`
	CreatedAt           string   `json:"created_at"`
	UpdatedAt           string   `json:"updated_at"`
}

// UnityPackage はUnityパッケージ情報です
type UnityPackage struct {
	ID            string `json:"id"`
	AssetURL      string `json:"assetUrl"`
	AssetURLObject string `json:"assetUrlObject"`
	AssetVersion  int    `json:"assetVersion"`
	CreatedAt     string `json:"created_at"`
	Platform      string `json:"platform"`
	PluginURL     string `json:"pluginUrl"`
	PluginURLObject string `json:"pluginUrlObject"`
	UnitySortNumber int64  `json:"unitySortNumber"`
	UnityVersion  string `json:"unityVersion"`
}

// GetAvatar は指定されたアバターIDのアバター情報を取得します
func (c *Client) GetAvatar(ctx context.Context, avatarID string) (*Avatar, error) {
	var avatar Avatar
	err := c.doRequest(ctx, "GET", "/avatars/"+avatarID, nil, &avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatar: %w", err)
	}
	return &avatar, nil
}

// SearchAvatarsOptions はアバター検索のオプションです
type SearchAvatarsOptions struct {
	Featured      bool   // 注目のアバターのみ
	Tag           string // タグでフィルター
	UserID        string // 特定のユーザーのアバター
	N             int    // 取得件数（デフォルト: 60）
	Offset        int    // オフセット
	Order         string // ソート順: "ascending", "descending"
	Sort          string // ソート項目: "created", "updated", "order", "_created_at", "_updated_at"
	ReleaseStatus string // リリースステータス: "public", "private", "hidden", "all"
	MaxUnityVersion string // 最大Unityバージョン
	MinUnityVersion string // 最小Unityバージョン
	Platform      string // プラットフォーム: "android", "standalonewindows"
}

// SearchAvatars はアバターを検索します
func (c *Client) SearchAvatars(ctx context.Context, opts SearchAvatarsOptions) ([]Avatar, error) {
	params := url.Values{}
	if opts.Featured {
		params.Set("featured", "true")
	}
	if opts.Tag != "" {
		params.Set("tag", opts.Tag)
	}
	if opts.UserID != "" {
		params.Set("userId", opts.UserID)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Order != "" {
		params.Set("order", opts.Order)
	}
	if opts.Sort != "" {
		params.Set("sort", opts.Sort)
	}
	if opts.ReleaseStatus != "" {
		params.Set("releaseStatus", opts.ReleaseStatus)
	}
	if opts.MaxUnityVersion != "" {
		params.Set("maxUnityVersion", opts.MaxUnityVersion)
	}
	if opts.MinUnityVersion != "" {
		params.Set("minUnityVersion", opts.MinUnityVersion)
	}
	if opts.Platform != "" {
		params.Set("platform", opts.Platform)
	}

	var avatars []Avatar
	path := "/avatars"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &avatars)
	if err != nil {
		return nil, fmt.Errorf("failed to search avatars: %w", err)
	}
	return avatars, nil
}

// SelectAvatar は指定されたアバターを装着します
func (c *Client) SelectAvatar(ctx context.Context, avatarID string) (*CurrentUser, error) {
	var user CurrentUser
	err := c.doRequest(ctx, "PUT", "/avatars/"+avatarID+"/select", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to select avatar: %w", err)
	}
	return &user, nil
}

// GetOwnAvatar は自分が作成したアバターを取得します
func (c *Client) GetOwnAvatar(ctx context.Context, avatarID string) (*Avatar, error) {
	var avatar Avatar
	err := c.doRequest(ctx, "GET", "/avatars/"+avatarID, nil, &avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to get own avatar: %w", err)
	}
	return &avatar, nil
}

// CreateAvatarRequest はアバター作成リクエストです
type CreateAvatarRequest struct {
	AssetURL        string `json:"assetUrl"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Tags            []string `json:"tags"`
	ImageURL        string `json:"imageUrl"`
	ReleaseStatus   string `json:"releaseStatus"`
	Version         int    `json:"version"`
	UnityPackageURL string `json:"unityPackageUrl"`
}

// CreateAvatar は新しいアバターを作成します
func (c *Client) CreateAvatar(ctx context.Context, req CreateAvatarRequest) (*Avatar, error) {
	var avatar Avatar
	err := c.doRequest(ctx, "POST", "/avatars", req, &avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to create avatar: %w", err)
	}
	return &avatar, nil
}

// UpdateAvatarRequest はアバター更新リクエストです
type UpdateAvatarRequest struct {
	AssetURL        *string  `json:"assetUrl,omitempty"`
	Name            *string  `json:"name,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	ImageURL        *string  `json:"imageUrl,omitempty"`
	ReleaseStatus   *string  `json:"releaseStatus,omitempty"`
	Version         *int     `json:"version,omitempty"`
	UnityPackageURL *string  `json:"unityPackageUrl,omitempty"`
}

// UpdateAvatar はアバター情報を更新します
func (c *Client) UpdateAvatar(ctx context.Context, avatarID string, req UpdateAvatarRequest) (*Avatar, error) {
	var avatar Avatar
	err := c.doRequest(ctx, "PUT", "/avatars/"+avatarID, req, &avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to update avatar: %w", err)
	}
	return &avatar, nil
}

// DeleteAvatar はアバターを削除します
func (c *Client) DeleteAvatar(ctx context.Context, avatarID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/avatars/"+avatarID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete avatar: %w", err)
	}
	return nil
}

// GetFavoriteAvatars はお気に入りのアバターを取得します
func (c *Client) GetFavoriteAvatars(ctx context.Context) ([]Avatar, error) {
	return c.SearchAvatars(ctx, SearchAvatarsOptions{
		Tag: "favorite",
	})
}
