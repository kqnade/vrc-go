package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// Favorite はお気に入り情報です
type Favorite struct {
	ID         string   `json:"id"`
	Type       string   `json:"type"` // "world", "avatar", "friend"
	FavoriteID string   `json:"favoriteId"`
	Tags       []string `json:"tags"`
}

// FavoriteGroup はお気に入りグループ情報です
type FavoriteGroup struct {
	ID             string `json:"id"`
	Type           string `json:"type"` // "world", "avatar", "friend"
	OwnerID        string `json:"ownerId"`
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	Visibility     string `json:"visibility"`
	Tags           []string `json:"tags"`
	OwnerDisplayName string `json:"ownerDisplayName"`
}

// AddFavoriteRequest はお気に入り追加リクエストです
type AddFavoriteRequest struct {
	Type       string   `json:"type"` // "world", "avatar", "friend"
	FavoriteID string   `json:"favoriteId"`
	Tags       []string `json:"tags"`
}

// AddFavorite はお気に入りを追加します
func (c *Client) AddFavorite(ctx context.Context, req AddFavoriteRequest) (*Favorite, error) {
	var favorite Favorite
	err := c.doRequest(ctx, "POST", "/favorites", req, &favorite)
	if err != nil {
		return nil, fmt.Errorf("failed to add favorite: %w", err)
	}
	return &favorite, nil
}

// RemoveFavorite はお気に入りを削除します
func (c *Client) RemoveFavorite(ctx context.Context, favoriteID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/favorites/"+favoriteID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}
	return nil
}

// GetFavoritesOptions はお気に入り取得のオプションです
type GetFavoritesOptions struct {
	Type   string // "world", "avatar", "friend"
	Tag    string // タグでフィルター
	N      int    // 取得件数（デフォルト: 60）
	Offset int    // オフセット
}

// GetFavorites はお気に入りのリストを取得します
func (c *Client) GetFavorites(ctx context.Context, opts GetFavoritesOptions) ([]Favorite, error) {
	params := url.Values{}
	if opts.Type != "" {
		params.Set("type", opts.Type)
	}
	if opts.Tag != "" {
		params.Set("tag", opts.Tag)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	var favorites []Favorite
	path := "/favorites"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &favorites)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorites: %w", err)
	}
	return favorites, nil
}

// GetFavoriteGroups はお気に入りグループのリストを取得します
func (c *Client) GetFavoriteGroups(ctx context.Context, favoriteType string) ([]FavoriteGroup, error) {
	params := url.Values{}
	if favoriteType != "" {
		params.Set("type", favoriteType)
	}

	var groups []FavoriteGroup
	path := "/favorite/groups"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite groups: %w", err)
	}
	return groups, nil
}

// GetFavoriteGroup は指定されたお気に入りグループ情報を取得します
func (c *Client) GetFavoriteGroup(ctx context.Context, favoriteGroupType, favoriteGroupName, userID string) (*FavoriteGroup, error) {
	var group FavoriteGroup
	err := c.doRequest(ctx, "GET", "/favorite/group/"+favoriteGroupType+"/"+favoriteGroupName+"/"+userID, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite group: %w", err)
	}
	return &group, nil
}

// UpdateFavoriteGroupRequest はお気に入りグループ更新リクエストです
type UpdateFavoriteGroupRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
	Visibility  *string `json:"visibility,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// UpdateFavoriteGroup はお気に入りグループを更新します
func (c *Client) UpdateFavoriteGroup(ctx context.Context, favoriteGroupType, favoriteGroupName, userID string, req UpdateFavoriteGroupRequest) (*FavoriteGroup, error) {
	var group FavoriteGroup
	err := c.doRequest(ctx, "PUT", "/favorite/group/"+favoriteGroupType+"/"+favoriteGroupName+"/"+userID, req, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to update favorite group: %w", err)
	}
	return &group, nil
}

// ClearFavoriteGroup はお気に入りグループをクリアします
func (c *Client) ClearFavoriteGroup(ctx context.Context, favoriteGroupType, favoriteGroupName, userID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/favorite/group/"+favoriteGroupType+"/"+favoriteGroupName+"/"+userID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to clear favorite group: %w", err)
	}
	return nil
}

// GetFavorite は指定されたお気に入り情報を取得します
func (c *Client) GetFavorite(ctx context.Context, favoriteID string) (*Favorite, error) {
	var favorite Favorite
	err := c.doRequest(ctx, "GET", "/favorites/"+favoriteID, nil, &favorite)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite: %w", err)
	}
	return &favorite, nil
}
