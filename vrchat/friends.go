package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// GetFriendsOptions はフレンドリスト取得のオプションです
type GetFriendsOptions struct {
	Offset  int  // オフセット
	N       int  // 取得件数（デフォルト: 60）
	Offline bool // オフラインのフレンドも含める
}

// GetFriends はフレンドリストを取得します
func (c *Client) GetFriends(ctx context.Context, opts GetFriendsOptions) ([]LimitedUser, error) {
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

	var friends []LimitedUser
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

// FriendStatus はフレンドステータスです
type FriendStatus struct {
	IsFriend         bool   `json:"isFriend"`
	IncomingRequest  bool   `json:"incomingRequest"`
	OutgoingRequest  bool   `json:"outgoingRequest"`
}

// GetFriendStatus は指定されたユーザーとのフレンドステータスを取得します
func (c *Client) GetFriendStatus(ctx context.Context, userID string) (*FriendStatus, error) {
	var status FriendStatus
	err := c.doRequest(ctx, "GET", "/user/"+userID+"/friendStatus", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("failed to get friend status: %w", err)
	}
	return &status, nil
}

// Notification は通知情報です
type Notification struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	SenderUserID string `json:"senderUserId"`
	SenderUsername string `json:"senderUsername"`
	ReceiverUserID string `json:"receiverUserId"`
	Message      string `json:"message"`
	Details      map[string]interface{} `json:"details"`
	Seen         bool   `json:"seen"`
	CreatedAt    string `json:"created_at"`
}

// SendFriendRequest はフレンドリクエストを送信します
func (c *Client) SendFriendRequest(ctx context.Context, userID string) (*Notification, error) {
	var notification Notification
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
func (c *Client) AcceptFriendRequest(ctx context.Context, notificationID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/accept", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to accept friend request: %w", err)
	}
	return nil
}

// DeclineFriendRequest はフレンドリクエストを拒否します
func (c *Client) DeclineFriendRequest(ctx context.Context, notificationID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/hide", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to decline friend request: %w", err)
	}
	return nil
}

// GetOnlineFriends はオンラインのフレンドリストを取得します
func (c *Client) GetOnlineFriends(ctx context.Context) ([]LimitedUser, error) {
	return c.GetFriends(ctx, GetFriendsOptions{Offline: false})
}

// GetOfflineFriends はオフラインのフレンドリストを取得します
func (c *Client) GetOfflineFriends(ctx context.Context) ([]LimitedUser, error) {
	params := url.Values{}
	params.Set("offline", "true")

	var friends []LimitedUser
	err := c.doRequest(ctx, "GET", "/auth/user/friends?"+params.Encode(), nil, &friends)
	if err != nil {
		return nil, fmt.Errorf("failed to get offline friends: %w", err)
	}
	return friends, nil
}
