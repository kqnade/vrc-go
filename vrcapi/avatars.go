package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetAvatar は指定されたアバターIDのアバター情報を取得します
func (c *Client) GetAvatar(ctx context.Context, avatarID string) (*shared.Avatar, error) {
	var avatar shared.Avatar
	err := c.doRequest(ctx, "GET", "/avatars/"+avatarID, nil, &avatar)
	if err != nil {
		return nil, fmt.Errorf("failed to get avatar: %w", err)
	}
	return &avatar, nil
}

// SearchAvatars はアバターを検索します
func (c *Client) SearchAvatars(ctx context.Context, opts shared.SearchAvatarsOptions) ([]shared.Avatar, error) {
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

	var avatars []shared.Avatar
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

// WearAvatar は指定されたアバターを装着します
func (c *Client) WearAvatar(ctx context.Context, avatarID string) (*shared.CurrentUser, error) {
	var user shared.CurrentUser
	req := struct {
		AvatarID string `json:"avatarId"`
	}{AvatarID: avatarID}
	err := c.doRequest(ctx, "PUT", "/auth/user/avatar", req, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to wear avatar: %w", err)
	}
	return &user, nil
}
