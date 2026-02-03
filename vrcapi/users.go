package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetUser は指定されたユーザーIDのユーザー情報を取得します
func (c *Client) GetUser(ctx context.Context, userID string) (*shared.User, error) {
	var user shared.User
	err := c.doRequest(ctx, "GET", "/users/"+userID, nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// GetUserByName は指定されたユーザー名のユーザー情報を取得します
func (c *Client) GetUserByName(ctx context.Context, username string) (*shared.User, error) {
	var user shared.User
	err := c.doRequest(ctx, "GET", "/users/"+username+"/name", nil, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by name: %w", err)
	}
	return &user, nil
}

// SearchUsers はユーザーを検索します
func (c *Client) SearchUsers(ctx context.Context, opts shared.SearchUsersOptions) ([]shared.LimitedUser, error) {
	params := url.Values{}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	if opts.DeveloperType != "" {
		params.Set("developerType", opts.DeveloperType)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Fuzzy {
		params.Set("fuzzy", "true")
	}

	var users []shared.LimitedUser
	path := "/users?" + params.Encode()
	err := c.doRequest(ctx, "GET", path, nil, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	return users, nil
}

// UpdateUser は現在のユーザー情報を更新します
func (c *Client) UpdateUser(ctx context.Context, userID string, req shared.UpdateUserRequest) (*shared.CurrentUser, error) {
	var user shared.CurrentUser
	err := c.doRequest(ctx, "PUT", "/users/"+userID, req, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return &user, nil
}

// GetUserGroups は指定されたユーザーが所属するグループのリストを取得します
func (c *Client) GetUserGroups(ctx context.Context, userID string) ([]shared.UserGroup, error) {
	var groups []shared.UserGroup
	err := c.doRequest(ctx, "GET", "/users/"+userID+"/groups", nil, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to get user groups: %w", err)
	}
	return groups, nil
}
