package vrcapi

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kqnade/vrcgo/shared"
)

// GetFriends はフレンドリストを取得します
func (c *Client) GetFriends(ctx context.Context, opts shared.GetFriendsOptions) ([]shared.LimitedUser, error) {
	params := url.Values{}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offline {
		params.Set("offline", "true")
	}

	var friends []shared.LimitedUser
	path := "/auth/user/friends"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &friends)
	if err != nil {
		return nil, fmt.Errorf("failed to get friends: %w", err)
	}
	return friends, nil
}

// GetFriendStatus は指定されたユーザーとのフレンドステータスを取得します
func (c *Client) GetFriendStatus(ctx context.Context, userID string) (*shared.FriendStatus, error) {
	var status shared.FriendStatus
	err := c.doRequest(ctx, "GET", "/user/"+userID+"/friendStatus", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("failed to get friend status: %w", err)
	}
	return &status, nil
}

// SendFriendRequest はフレンドリクエストを送信します
func (c *Client) SendFriendRequest(ctx context.Context, userID string) (*shared.Notification, error) {
	var notification shared.Notification
	err := c.doRequest(ctx, "POST", "/user/"+userID+"/friendRequest", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send friend request: %w", err)
	}
	return &notification, nil
}

// DeleteFriend はフレンドを削除します
func (c *Client) DeleteFriend(ctx context.Context, userID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "DELETE", "/auth/user/friends/"+userID, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete friend: %w", err)
	}
	return nil
}

// AcceptFriendRequest はフレンドリクエストを承認します
func (c *Client) AcceptFriendRequest(ctx context.Context, notificationID string) (*shared.Notification, error) {
	var notification shared.Notification
	err := c.doRequest(ctx, "PUT", "/auth/user/friendRequests/"+notificationID+"/accept", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to accept friend request: %w", err)
	}
	return &notification, nil
}

// RejectFriendRequest はフレンドリクエストを拒否します
func (c *Client) RejectFriendRequest(ctx context.Context, notificationID string) (*shared.Notification, error) {
	var notification shared.Notification
	err := c.doRequest(ctx, "PUT", "/auth/user/friendRequests/"+notificationID+"/reject", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to reject friend request: %w", err)
	}
	return &notification, nil
}
