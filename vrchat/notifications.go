package vrchat

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

// GetNotificationsOptions は通知取得のオプションです
type GetNotificationsOptions struct {
	Type   string // 通知タイプでフィルター
	Sent   bool   // 送信した通知
	Hidden bool   // 非表示の通知も含める
	After  string // 指定日時以降の通知
	N      int    // 取得件数（デフォルト: 60）
	Offset int    // オフセット
}

// GetNotifications は通知のリストを取得します
func (c *Client) GetNotifications(ctx context.Context, opts GetNotificationsOptions) ([]Notification, error) {
	params := url.Values{}
	if opts.Type != "" {
		params.Set("type", opts.Type)
	}
	if opts.Sent {
		params.Set("sent", "true")
	}
	if opts.Hidden {
		params.Set("hidden", "true")
	}
	if opts.After != "" {
		params.Set("after", opts.After)
	}
	if opts.N > 0 {
		params.Set("n", strconv.Itoa(opts.N))
	} else {
		params.Set("n", "60")
	}
	if opts.Offset > 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	var notifications []Notification
	path := "/auth/user/notifications"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	err := c.doRequest(ctx, "GET", path, nil, &notifications)
	if err != nil {
		return nil, fmt.Errorf("failed to get notifications: %w", err)
	}
	return notifications, nil
}

// MarkNotificationAsRead は通知を既読にマークします
func (c *Client) MarkNotificationAsRead(ctx context.Context, notificationID string) (*Notification, error) {
	var notification Notification
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/see", nil, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to mark notification as read: %w", err)
	}
	return &notification, nil
}

// DeleteNotification は通知を削除します
func (c *Client) DeleteNotification(ctx context.Context, notificationID string) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/"+notificationID+"/hide", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}

// ClearAllNotifications はすべての通知をクリアします
func (c *Client) ClearAllNotifications(ctx context.Context) error {
	var response struct {
		Success struct {
			Message    string `json:"message"`
			StatusCode int    `json:"status_code"`
		} `json:"success"`
	}
	err := c.doRequest(ctx, "PUT", "/auth/user/notifications/clear", nil, &response)
	if err != nil {
		return fmt.Errorf("failed to clear all notifications: %w", err)
	}
	return nil
}

// SendNotificationRequest は通知送信リクエストです
type SendNotificationRequest struct {
	Type    string                 `json:"type"`
	UserID  string                 `json:"userId"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// SendNotification は通知を送信します
func (c *Client) SendNotification(ctx context.Context, req SendNotificationRequest) (*Notification, error) {
	var notification Notification
	err := c.doRequest(ctx, "POST", "/auth/user/notifications", req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %w", err)
	}
	return &notification, nil
}

// SendInviteRequest は招待通知送信リクエストです
type SendInviteRequest struct {
	InstanceID string `json:"instanceId"`
	MessageSlot int   `json:"messageSlot,omitempty"`
}

// SendInvite はインスタンスへの招待通知を送信します
func (c *Client) SendInvite(ctx context.Context, userID string, req SendInviteRequest) (*Notification, error) {
	var notification Notification
	err := c.doRequest(ctx, "POST", "/invite/"+userID, req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to send invite: %w", err)
	}
	return &notification, nil
}

// RespondToInviteRequest は招待への応答リクエストです
type RespondToInviteRequest struct {
	ResponseSlot int `json:"responseSlot"`
}

// RespondToInvite は招待に応答します
func (c *Client) RespondToInvite(ctx context.Context, notificationID string, req RespondToInviteRequest) (*Notification, error) {
	var notification Notification
	err := c.doRequest(ctx, "POST", "/invite/"+notificationID+"/response", req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to respond to invite: %w", err)
	}
	return &notification, nil
}

// RequestInvite はインスタンスへの招待をリクエストします
func (c *Client) RequestInvite(ctx context.Context, userID string, instanceLocation string) (*Notification, error) {
	req := SendNotificationRequest{
		Type:   "requestInvite",
		UserID: userID,
		Details: map[string]interface{}{
			"platform": "standalonewindows",
		},
	}
	if instanceLocation != "" {
		req.Details["worldId"] = instanceLocation
	}

	var notification Notification
	err := c.doRequest(ctx, "POST", "/requestInvite/"+userID, req, &notification)
	if err != nil {
		return nil, fmt.Errorf("failed to request invite: %w", err)
	}
	return &notification, nil
}

// GetInviteMessages は招待メッセージのリストを取得します
func (c *Client) GetInviteMessages(ctx context.Context) ([]struct {
	MessageSlot int    `json:"messageSlot"`
	Message     string `json:"message"`
	UpdatedAt   string `json:"updatedAt"`
}, error) {
	var messages []struct {
		MessageSlot int    `json:"messageSlot"`
		Message     string `json:"message"`
		UpdatedAt   string `json:"updatedAt"`
	}
	err := c.doRequest(ctx, "GET", "/message", nil, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to get invite messages: %w", err)
	}
	return messages, nil
}

// UpdateInviteMessageRequest は招待メッセージ更新リクエストです
type UpdateInviteMessageRequest struct {
	Message string `json:"message"`
}

// UpdateInviteMessage は招待メッセージを更新します
func (c *Client) UpdateInviteMessage(ctx context.Context, messageSlot int, req UpdateInviteMessageRequest) ([]struct {
	MessageSlot int    `json:"messageSlot"`
	Message     string `json:"message"`
	UpdatedAt   string `json:"updatedAt"`
}, error) {
	var messages []struct {
		MessageSlot int    `json:"messageSlot"`
		Message     string `json:"message"`
		UpdatedAt   string `json:"updatedAt"`
	}
	err := c.doRequest(ctx, "PUT", "/message/"+strconv.Itoa(messageSlot), req, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to update invite message: %w", err)
	}
	return messages, nil
}

// ResetInviteMessage は招待メッセージをリセットします
func (c *Client) ResetInviteMessage(ctx context.Context, messageSlot int) ([]struct {
	MessageSlot int    `json:"messageSlot"`
	Message     string `json:"message"`
	UpdatedAt   string `json:"updatedAt"`
}, error) {
	var messages []struct {
		MessageSlot int    `json:"messageSlot"`
		Message     string `json:"message"`
		UpdatedAt   string `json:"updatedAt"`
	}
	err := c.doRequest(ctx, "DELETE", "/message/"+strconv.Itoa(messageSlot), nil, &messages)
	if err != nil {
		return nil, fmt.Errorf("failed to reset invite message: %w", err)
	}
	return messages, nil
}
