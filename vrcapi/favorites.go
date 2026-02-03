package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// AddFavorite はお気に入りに追加します
func (c *Client) AddFavorite(ctx context.Context, favoriteType, favoriteID string, tags []string) (*shared.Favorite, error) {
	var favorite shared.Favorite
	req := struct {
		Type       string   `json:"type"`
		FavoriteID string   `json:"favoriteId"`
		Tags       []string `json:"tags"`
	}{
		Type:       favoriteType,
		FavoriteID: favoriteID,
		Tags:       tags,
	}
	err := c.doRequest(ctx, "POST", "/favorites", req, &favorite)
	if err != nil {
		return nil, fmt.Errorf("failed to add favorite: %w", err)
	}
	return &favorite, nil
}

// RemoveFavorite はお気に入りから削除します
func (c *Client) RemoveFavorite(ctx context.Context, favoriteID string) error {
	err := c.doRequest(ctx, "DELETE", "/favorites/"+favoriteID, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to remove favorite: %w", err)
	}
	return nil
}

// GetFavorites はお気に入りのリストを取得します
func (c *Client) GetFavorites(ctx context.Context, n, offset int, favoriteType, tag string) ([]shared.Favorite, error) {
	var favorites []shared.Favorite
	params := url.Values{}
	params.Set("n", strconv.Itoa(n))
	params.Set("offset", strconv.Itoa(offset))
	if favoriteType != "" {
		params.Set("type", favoriteType)
	}
	if tag != "" {
		params.Set("tag", tag)
	}
	path := "/favorites?" + params.Encode()
	err := c.doRequest(ctx, "GET", path, nil, &favorites)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorites: %w", err)
	}
	return favorites, nil
}

// GetFavoriteGroups はお気に入りグループのリストを取得します
func (c *Client) GetFavoriteGroups(ctx context.Context, n, offset int, ownerID string) ([]shared.FavoriteGroup, error) {
	var groups []shared.FavoriteGroup
	params := url.Values{}
	params.Set("n", strconv.Itoa(n))
	params.Set("offset", strconv.Itoa(offset))
	if ownerID != "" {
		params.Set("ownerId", ownerID)
	}
	path := "/favorite/groups?" + params.Encode()
	err := c.doRequest(ctx, "GET", path, nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to get favorite groups: %w", err)
	}
	return groups, nil
}
